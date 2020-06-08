package webapi

import (
	"fmt"
	"strings"
	"sync"

	cmn "github.com/cdutwhu/n3-util/common"
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
	struct2Map    = cmn.Struct2Map
	url1Value     = cmn.URL1Value
	env2Struct    = cmn.Env2Struct
)

func initMutex(route interface{}) map[string]*sync.Mutex {
	mMtx := make(map[string]*sync.Mutex)
	m, err := struct2Map(route)
	failOnErr("%v", err)
	for _, v := range m {
		mMtx[v.(string)] = &sync.Mutex{}
	}
	return mMtx
}

type result struct {
	Data  string `json:"data"`
	Info  string `json:"info"`
	Error string `json:"error"`
}
