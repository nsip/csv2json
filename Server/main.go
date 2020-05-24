package main

import (
	eg "github.com/cdutwhu/n3-util/n3errs"
	g "github.com/nsip/n3-csv2json/Server/global"
	api "github.com/nsip/n3-csv2json/Server/webapi"
)

func main() {
	failOnErrWhen(!g.Init(), "%v: Global Config Init Error", eg.CFG_INIT_ERR)

	cfg := g.Cfg
	ws, logfile, servicename := cfg.WebService, cfg.LogFile, cfg.ServiceName

	setLog(logfile)
	fPln(logWhen(true, "[%s] Hosting on: [%v:%d], version [%v]", servicename, localIP(), ws.Port, ws.Version))

	done := make(chan string)
	go api.HostHTTPAsync()
	<-done
}
