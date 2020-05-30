package main

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
	sHasSuffix = strings.HasSuffix
	sTrimRight = strings.TrimRight
	sReplace   = strings.Replace

	isFLog        = cmn.IsFLog
	failOnErrWhen = cmn.FailOnErrWhen
	failOnErr     = cmn.FailOnErr
	warnOnErrWhen = cmn.WarnOnErrWhen
	warnOnErr     = cmn.WarnOnErr
	setLog        = cmn.SetLog
	resetLog      = cmn.ResetLog
	isXML         = cmn.IsXML
	isJSON        = cmn.IsJSON
	env2Struct    = cmn.Env2Struct
	struct2Env    = cmn.Struct2Env
	structFields  = cmn.StructFields
)
