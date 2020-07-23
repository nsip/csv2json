package main

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/n3log"
)

var (
	fPln           = fmt.Println
	fSf            = fmt.Sprintf
	failOnErr      = fn.FailOnErr
	failOnErrWhen  = fn.FailOnErrWhen
	enableLog2F    = fn.EnableLog2F
	logWhen        = fn.LoggerWhen
	logger         = fn.Logger
	env2Struct     = rflx.Env2Struct
	localIP        = net.LocalIP
	lrInit         = n3log.LrInit
	lrOut          = n3log.LrOut
	enableLoggly   = n3log.EnableLoggly
	setLogglyToken = n3log.SetLogglyToken
)
