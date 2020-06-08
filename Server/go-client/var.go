package client

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

var (
	fPt  = fmt.Print
	fPf  = fmt.Printf
	fPln = fmt.Println
	fSf  = fmt.Sprintf

	sJoin      = strings.Join
	sReplace   = strings.Replace
	sTrimRight = strings.TrimRight

	struct2Map    = cmn.Struct2Map
	mapKeys       = cmn.MapKeys
	failOnErrWhen = cmn.FailOnErrWhen
	failOnErr     = cmn.FailOnErr
	logWhen       = cmn.LogWhen
	warnOnErr     = cmn.WarnOnErr
	warnOnErrWhen = cmn.WarnOnErrWhen
	env2Struct    = cmn.Env2Struct
	struct2Env    = cmn.Struct2Env
	setLog        = cmn.SetLog
	isXML         = cmn.IsXML
	isJSON        = cmn.IsJSON
	cfgRepl       = cmn.CfgRepl
)

const (
	envVarName = "CfgClt-C2J"
)

// Args is arguments for "Route"
type Args struct {
	Data   []byte
	ToNATS bool
}

func initMapFnURL(protocol, ip string, port int, route interface{}) (map[string]string, []string) {
	mFnURL := make(map[string]string)
	m, err := struct2Map(route)
	failOnErr("%v", err)
	for k, v := range m {
		mFnURL[k] = fSf("%s://%s:%d%s", protocol, ip, port, v)
	}
	IKeys, err := mapKeys(mFnURL)
	failOnErr("%v", err)
	return mFnURL, IKeys.([]string)
}

func initTracer(serviceName string) opentracing.Tracer {
	cfg, err := config.FromEnv()
	failOnErr("%v: ", err)
	cfg.ServiceName = serviceName
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1

	tracer, _, err := cfg.NewTracer()
	failOnErr("%v: ", err)
	return tracer
}
