package webapi

import (
	"net/http"
	"os"

	"github.com/cdutwhu/n3-util/n3csv"
	eg "github.com/cdutwhu/n3-util/n3errs"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/middleware"
	glb "github.com/nsip/n3-csv2json/Server/global"
)

// HostHTTPAsync : Host a HTTP Server for CSV or JSON
func HostHTTPAsync() {
	e := echo.New()
	defer e.Close()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2G"))

	// Add Jaeger Tracer into Middleware
	c := jaegertracing.New(e, nil)
	defer c.Close()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST},
		AllowCredentials: true,
	}))

	cfg := glb.Cfg
	port := cfg.WebService.Port
	fullIP := localIP() + fSf(":%d", port)
	route := cfg.Route
	file := cfg.File

	mMtx := initMutex()

	defer e.Start(fSf(":%d", port))

	// *************************************** List all API, FILE *************************************** //

	path := route.HELP
	e.GET(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		return c.String(http.StatusOK,
			fSf("wget %-55s-> %s\n", fullIP+"/client-linux64", "Get Client(Linux64)")+
				fSf("wget %-55s-> %s\n", fullIP+"/client-mac", "Get Client(Mac)")+
				fSf("wget %-55s-> %s\n", fullIP+"/client-win64", "Get Client(Windows64)")+
				fSf("wget -O config.toml %-40s-> %s\n", fullIP+"/client-config", "Get Client Config")+
				fSf("\n")+
				fSf("POST %-55s-> %s\n"+
					"POST %-55s-> %s\n",
					fullIP+route.CSV2JSON, "Upload CSV, return JSON.",
					fullIP+route.JSON2CSV, "Upload JSON, return CSV."))
	})

	// ------------------------------------------------------------------------------------ //

	mRouteRes := map[string]string{
		"/client-linux64": file.ClientLinux64,
		"/client-mac":     file.ClientMac,
		"/client-win64":   file.ClientWin64,
		"/client-config":  file.ClientConfig,
	}

	routeFun := func(rt, res string) func(c echo.Context) error {
		return func(c echo.Context) (err error) {
			if _, err = os.Stat(res); err == nil {
				fPln(rt, res)
				return c.File(res)
			}
			fPf("%v\n", warnOnErr("%v: [%s]  get [%s]", eg.FILE_NOT_FOUND, rt, res))
			return err
		}
	}

	for rt, res := range mRouteRes {
		e.GET(rt, routeFun(rt, res))
	}

	// ------------------------------------------------------------------------------------ //

	path = route.CSV2JSON
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		// bytes, err := ioutil.ReadAll(c.Request().Body)
		// csvstr := string(bytes)
		// log("\n%s\n", csvstr)

		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, result{
		// 		Data:  nil,
		// 		Info:  "",
		// 		Error: err.Error(),
		// 	})
		// }

		// jsonstr, headers := n3csv.Reader2JSON(c.Request().Body, "")

		// Trace [n3csv.Reader2JSON]
		results := jaegertracing.TraceFunction(c, n3csv.Reader2JSON, c.Request().Body, "")
		jsonstr := results[0].Interface().(string)
		headers := results[1].Interface().([]string)

		return c.JSON(http.StatusOK, result{
			Data:  &jsonstr,
			Info:  sJoin(headers, ","),
			Error: "",
		})
	})

	path = route.JSON2CSV
	e.POST(path, func(c echo.Context) error {
		defer func() { mMtx[path].Unlock() }()
		mMtx[path].Lock()

		return c.JSON(http.StatusInternalServerError, result{
			Data:  nil,
			Info:  "Not implemented",
			Error: "Not implemented",
		})
	})
}
