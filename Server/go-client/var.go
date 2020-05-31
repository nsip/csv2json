package client

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
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

// Args is arguments for "Route"
type Args struct {
	Data   []byte
	ToNATS bool
}

func initMapFnURL(protocol, ip string, port int, route interface{}) (map[string]string, []string) {
	mFnURL := make(map[string]string)
	for k, v := range struct2Map(route) {
		mFnURL[k] = fSf("%s://%s:%d%s", protocol, ip, port, v)
	}
	return mFnURL, mapKeys(mFnURL).([]string)
}
