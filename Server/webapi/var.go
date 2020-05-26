package webapi

import (
	"fmt"
	"strings"
	"sync"

	cmn "github.com/cdutwhu/n3-util/common"
	glb "github.com/nsip/n3-csv2json/Server/global"
)

var (
	fPf         = fmt.Printf
	fPln        = fmt.Println
	fSf         = fmt.Sprintf
	sReplaceAll = strings.ReplaceAll
	sJoin       = strings.Join

	localIP       = cmn.LocalIP
	isXML         = cmn.IsXML
	isJSON        = cmn.IsJSON
	setLog        = cmn.SetLog
	log           = cmn.Log
	warnOnErr     = cmn.WarnOnErr
	failOnErr     = cmn.FailOnErr
	mustWriteFile = cmn.MustWriteFile
	mapFromStruct = cmn.MapFromStruct
)

func initMutex() map[string]*sync.Mutex {
	mMtx := make(map[string]*sync.Mutex)
	for _, v := range mapFromStruct(glb.Cfg.Route) {
		mMtx[v.(string)] = &sync.Mutex{}
	}
	return mMtx
}

type result struct {
	Data  *string `json:"data"`
	Info  string  `json:"info"`
	Error string  `json:"error"`
}
