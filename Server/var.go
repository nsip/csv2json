package main

import (
	"fmt"

	cmn "github.com/cdutwhu/n3-util/common"
)

var (
	fPln          = fmt.Println
	fSf           = fmt.Sprintf
	failOnErr     = cmn.FailOnErr
	failOnErrWhen = cmn.FailOnErrWhen
	localIP       = cmn.LocalIP
	setLog        = cmn.SetLog
	logWhen       = cmn.LogWhen
	logger        = cmn.Log
	env2Struct    = cmn.Env2Struct
)
