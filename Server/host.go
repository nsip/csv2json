package main

import (
	"context"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3csv"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
)

func shutdownAsync(e *echo.Echo, sig <-chan os.Signal, done chan<- string) {
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	failOnErr("%v", e.Shutdown(ctx))
	time.Sleep(20 * time.Millisecond)
	done <- "Shutdown Successfully"
}

// HostHTTPAsync : Host a HTTP Server for CSV or JSON
func HostHTTPAsync(sig <-chan os.Signal, done chan<- string) {
	defer func() { logGrp.Do("HostHTTPAsync Exit") }()

	e := echo.New()
	defer e.Close()

	// waiting for shutdown
	go shutdownAsync(e, sig, done)

	// Add Jaeger Tracer into Middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2G"))
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST},
		AllowCredentials: true,
	}))

	var (
		Cfg    = n3cfg.FromEnvN3csv2jsonServer(envKey)
		port   = Cfg.WebService.Port
		fullIP = localIP() + fSf(":%d", port)
		route  = Cfg.Route
		mMtx   = initMutex(&route)
	)

	logGrp.Do("Echo Service is Starting")
	defer e.Start(fSf(":%d", port))

	// *************************************** List all API, FILE *************************************** //

	path := route.HELP
	e.GET(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		return c.String(http.StatusOK,
			// fSf("wget %-55s-> %s\n", fullIP+"/client-linux64", "Get Client(Linux64)")+
			// 	fSf("wget %-55s-> %s\n", fullIP+"/client-mac", "Get Client(Mac)")+
			// 	fSf("wget %-55s-> %s\n", fullIP+"/client-win64", "Get Client(Windows64)")+
			// 	fSf("wget -O config.toml %-40s-> %s\n", fullIP+"/client-config", "Get Client Config")+
			// 	fSf("\n")+
			fSf("POST %-55s-> %s\n"+
				"POST %-55s-> %s\n",
				fullIP+route.CSV2JSON, "Upload CSV, return JSON.",
				fullIP+route.JSON2CSV, "Upload JSON, return CSV."))
	})

	// ------------------------------------------------------------------------------------ //

	// mRouteRes := map[string]string{
	// 	"/client-linux64": Cfg.File.ClientLinux64,
	// 	"/client-mac":     Cfg.File.ClientMac,
	// 	"/client-win64":   Cfg.File.ClientWin64,
	// 	"/client-config":  Cfg.File.ClientConfig,
	// }

	// routeFun := func(rt, res string) func(c echo.Context) error {
	// 	return func(c echo.Context) (err error) {
	// 		if _, err = os.Stat(res); err == nil {
	// 			fPln(rt, res)
	// 			return c.File(res)
	// 		}
	// 		return warnOnErr("%v: [%s]  get [%s]", n3err.FILE_NOT_FOUND, rt, res)
	// 	}
	// }

	// for rt, res := range mRouteRes {
	// 	e.GET(rt, routeFun(rt, res))
	// }

	// ------------------------------------------------------------------------------------------------------------- //
	// ------------------------------------------------------------------------------------------------------------- //

	path = route.CSV2JSON
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		var (
			status  = http.StatusOK
			ret     string
			results []reflect.Value
			msg     bool
		)

		logGrp.Do("Parsing Params")
		if ok, n := url1Value(c.QueryParams(), 0, "nats"); ok && n != "" && n != "false" {
			msg = true
		}

		logGrp.Do("n3csv.Reader2JSON")
		// jsonstr, headers, err := n3csv.Reader2JSON(bytes.NewReader(body), "")
		// Trace [n3csv.Reader2JSON]
		results = jaegertracing.TraceFunction(c, n3csv.Reader2JSON, c.Request().Body, "")
		ret = results[0].Interface().(string)
		if !results[2].IsNil() {
			status = http.StatusInternalServerError
			ret = results[2].Interface().(error).Error()
			goto RET
		}
		logGrp.Do("CSV Headers: " + sJoin(results[1].Interface().([]string), " "))

		// Send a copy to NATS
		if msg {
			url, subj, timeout := Cfg.NATS.URL, Cfg.NATS.Subject, time.Duration(Cfg.NATS.Timeout)
			nc, err := nats.Connect(url)
			if err != nil {
				status = http.StatusInternalServerError
				ret = err.Error() + fSf(" @NATS Connect @Subject: [%s@%s]", url, subj)
				goto RET
			}
			msg, err := nc.Request(subj, []byte(ret), timeout*time.Millisecond)
			if err != nil {
				status = http.StatusInternalServerError
				ret = err.Error() + fSf(" @NATS Request @Subject: [%s@%s]", url, subj)
				goto RET
			}
			logGrp.Do(string(msg.Data))
		}

	RET:
		if status != http.StatusOK {
			warnGrp.Do(ret + " --> Failed")
		} else {
			logGrp.Do("--> Finish CSV2JSON")
		}
		return c.String(status, ret) // ret is already JSON String, so return String
	})

	// ------------------------------------------------------------------------------------------------------------- //
	// ------------------------------------------------------------------------------------------------------------- //

	// path = route.JSON2CSV
	// e.POST(path, func(c echo.Context) error {
	// 	defer func() { mMtx[path].Unlock() }()
	// 	mMtx[path].Lock()

	// 	RetErr := n3err.NOT_IMPLEMENTED

	// 	RetErrStr := ""
	// 	if RetErr != nil {
	// 		RetErrStr = RetErr.Error()
	// 	}
	// 	return c.JSON(http.StatusInternalServerError, result{
	// 		Data:  "",
	// 		Info:  "Not implemented",
	// 		Error: RetErrStr,
	// 	})
	// })
}
