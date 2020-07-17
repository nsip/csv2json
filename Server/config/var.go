package config

import (
	"fmt"
	"strings"

	"github.com/cdutwhu/debog/fn"
	"github.com/cdutwhu/gotil/net"
	"github.com/cdutwhu/gotil/rflx"
	"github.com/cdutwhu/n3-util/cfg"
)

var (
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	sReplaceAll = strings.ReplaceAll
	failOnErr   = fn.FailOnErr
	localIP     = net.LocalIP
	cfgRepl     = cfg.Modify
	gitver      = cfg.GitVer
	struct2Env  = rflx.Struct2Env
	env2Struct  = rflx.Env2Struct
)
