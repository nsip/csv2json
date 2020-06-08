package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"time"

	eg "github.com/cdutwhu/n3-util/n3errs"
	"github.com/opentracing/opentracing-go"
	tags "github.com/opentracing/opentracing-go/ext"
)

// DOwithTrace :
func DOwithTrace(ctx context.Context, configfile, fn string, args Args) (string, error) {
	failOnErrWhen(!initEnvVarFromTOML(envVarName, configfile), "%v", eg.CFG_INIT_ERR)
	ICfg, err := env2Struct(envVarName, &Config{})
	failOnErr("%v", err)
	Cfg := ICfg.(*Config)
	serviceName := Cfg.ServiceName

	if span := opentracing.SpanFromContext(ctx); span != nil {
		tracer := initTracer(serviceName)
		span := tracer.StartSpan(fn, opentracing.ChildOf(span.Context()))
		tags.SpanKindRPCClient.Set(span)
		tags.PeerService.Set(span, serviceName)
		span.SetTag(fn, args)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
	}
	return DO(configfile, fn, args)
}

// DO : fn ["HELP", "CSV2JSON", "JSON2CSV"]
func DO(configfile, fn string, args Args) (string, error) {
	failOnErrWhen(!initEnvVarFromTOML(envVarName, configfile), "%v", eg.CFG_INIT_ERR)
	ICfg, err := env2Struct(envVarName, &Config{})
	failOnErr("%v", err)
	Cfg := ICfg.(*Config)

	server := Cfg.Server
	protocol, ip, port := server.Protocol, server.IP, server.Port
	timeout := Cfg.Access.Timeout
	setLog(Cfg.LogFile)

	mFnURL, fields := initMapFnURL(protocol, ip, port, Cfg.Route)
	url, ok := mFnURL[fn]
	if err := warnOnErrWhen(!ok, "%v: Need %v", eg.PARAM_NOT_SUPPORTED, fields); err != nil {
		return "", err
	}

	chStr, chErr := make(chan string), make(chan error)
	go func() {
		rest(fn, url, args, chStr, chErr)
	}()

	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		return "", warnOnErr("%v: Didn't get response in %d(s)", eg.NET_TIMEOUT, timeout)
	case str := <-chStr:
		err := <-chErr
		if err == eg.NO_ERROR {
			return str, nil
		}
		return str, err
	}
}

// rest :
func rest(fn, url string, args Args, chStr chan string, chErr chan error) {

	paramN := ""
	if args.ToNATS {
		paramN = fSf("nats=true")
	}
	url = fSf("%s?%s", url, paramN)
	url = sReplace(url, "?&", "?", 1)
	url = sTrimRight(url, "?&")

	logWhen(true, "accessing ... %s", url)

	var (
		Resp    *http.Response
		Err     error
		RetData []byte
	)

	switch fn {
	case "HELP":
		if Resp, Err = http.Get(url); Err != nil {
			goto ERR_RET
		}

	case "CSV2JSON", "JSON2CSV":
		data := args.Data
		if data == nil {
			Err = eg.HTTP_REQBODY_EMPTY
			goto ERR_RET
		}
		if fn == "JSON2CSV" && !isJSON(string(data)) {
			Err = eg.PARAM_INVALID_JSON
			goto ERR_RET
		}
		if Resp, Err = http.Post(url, "application/json", bytes.NewBuffer(data)); Err != nil {
			goto ERR_RET
		}
	}

	if Resp == nil {
		Err = eg.NET_NO_RESPONSE
		goto ERR_RET
	}
	defer Resp.Body.Close()

	if RetData, Err = ioutil.ReadAll(Resp.Body); Err != nil {
		goto ERR_RET
	}

ERR_RET:
	if Err != nil {
		chStr <- ""
		chErr <- Err
		return
	}

	chStr <- string(RetData)
	chErr <- eg.NO_ERROR
	return
}
