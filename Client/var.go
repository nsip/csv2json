package main

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/judge"
	"github.com/cdutwhu/gotil/rflx"
)

var (
	fPt  = fmt.Print
	fPf  = fmt.Printf
	fPln = fmt.Println
	fSf  = fmt.Sprintf

	sJoin      = strings.Join
	sHasSuffix = strings.HasSuffix
	sTrimRight = strings.TrimRight
	sReplace   = strings.Replace

	failOnErrWhen = fn.FailOnErrWhen
	failOnErr     = fn.FailOnErr
	warnOnErrWhen = fn.WarnOnErrWhen
	warnOnErr     = fn.WarnOnErr
	setLog        = fn.SetLog
	resetLog      = fn.ResetLog
	isXML         = judge.IsXML
	isJSON        = judge.IsJSON
	env2Struct    = rflx.Env2Struct
	struct2Env    = rflx.Struct2Env
	structFields  = rflx.StructFields
)
