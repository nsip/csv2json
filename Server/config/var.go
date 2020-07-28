package config

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/n3cfg"
)

var (
	fPln          = fmt.Println
	fSf           = fmt.Sprintf
	sReplaceAll   = strings.ReplaceAll
	sHasSuffix    = strings.HasSuffix
	failOnErr     = fn.FailOnErr
	failOnErrWhen = fn.FailOnErrWhen
	localIP       = net.LocalIP
	cfgRepl       = n3cfg.Modify
	gitver        = n3cfg.GitVer
	struct2Env    = rflx.Struct2Env
	env2Struct    = rflx.Env2Struct
)
