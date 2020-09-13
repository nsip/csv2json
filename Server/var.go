package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/io"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3log"
	"github.com/cdutwhu/n3-util/rest"
)

var (
	fPt              = fmt.Print
	fPf              = fmt.Printf
	fEf              = fmt.Errorf
	fPln             = fmt.Println
	fSf              = fmt.Sprintf
	sReplaceAll      = strings.ReplaceAll
	sJoin            = strings.Join
	sReplace         = strings.Replace
	sTrimRight       = strings.TrimRight
	failOnErr        = fn.FailOnErr
	failOnErrWhen    = fn.FailOnErrWhen
	enableLog2F      = fn.EnableLog2F
	enableWarnDetail = fn.EnableWarnDetail
	logWhen          = fn.LoggerWhen
	logger           = fn.Logger
	warner           = fn.Warner
	warnOnErr        = fn.WarnOnErr
	warnOnErrWhen    = fn.WarnOnErrWhen
	env2Struct       = rflx.Env2Struct
	struct2Env       = rflx.Struct2Env
	struct2Map       = rflx.Struct2Map
	mapKeys          = rflx.MapKeys
	localIP          = net.LocalIP
	isXML            = judge.IsXML
	isJSON           = judge.IsJSON
	mustWriteFile    = io.MustWriteFile
	url1Value        = rest.URL1Value
	setLoggly        = n3log.SetLoggly
	syncBindLog      = n3log.SyncBindLog
	logBind          = n3log.Bind
	loggly           = n3log.Loggly
	cfgRepl          = n3cfg.Modify
)

const (
	envKey = "C2JSvr"
)

var (
	logGrp  = logBind(logger) // logBind(logger, loggly("info"))
	warnGrp = logBind(warner) // logBind(warner, loggly("warn"))
)

func initMutex(route interface{}) map[string]*sync.Mutex {
	mMtx := make(map[string]*sync.Mutex)
	for _, v := range struct2Map(route) {
		mMtx[v.(string)] = &sync.Mutex{}
	}
	return mMtx
}
