package main

import (
	"fmt"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
)

var (
	fPln          = fmt.Println
	fSf           = fmt.Sprintf
	failOnErr     = fn.FailOnErr
	failOnErrWhen = fn.FailOnErrWhen
	setLog        = fn.SetLog
	logWhen       = fn.LoggerWhen
	logger        = fn.Logger
	env2Struct    = rflx.Env2Struct
	localIP       = net.LocalIP
)
