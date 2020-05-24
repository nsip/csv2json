package config

import (
	"fmt"
	"strings"

	cmn "github.com/cdutwhu/n3-util/common"
)

var (
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	sReplaceAll = strings.ReplaceAll

	failOnErr = cmn.FailOnErr
	localIP   = cmn.LocalIP
	cfgRepl   = cmn.CfgRepl
)
