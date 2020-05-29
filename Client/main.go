package main

import (
	"encoding/json"
	"flag"
	"os"

	eg "github.com/cdutwhu/n3-util/n3errs"
	client "github.com/nsip/n3-csv2json/Server/go-client"
)

func main() {
	if e := warnOnErrWhen(len(os.Args) < 3, "%v: need [config.toml] [HELP CSV2JSON JSON2CSV]", eg.CLI_SUBCMD_ERR); e != nil {
		if isFLog() {
			fPf("%v\n", e)
		}
		return
	}

	cltcfg, fn := os.Args[1], os.Args[2]

	cmd := flag.NewFlagSet(fn, flag.ExitOnError)
	iPtr := cmd.String("i", "", "Path of original CSV/JSON file to be uploaded")
	wPtr := cmd.Bool("w", false, "whole dump flag: Print INFO & ERROR")    // true: print INFO & ERROR from Server
	nPtr := cmd.Bool("n", false, "indicate server to send a copy to NATS") // true: indicate server
	cmd.Parse(os.Args[3:])

	if fn == "CSV2JSON" || fn == "JSON2CSV" {
		failOnErrWhen(*iPtr == "", "%v: [-i] is required", eg.CLI_FLAG_ERR)
	}

	str, err := client.DO(
		cltcfg,
		fn,
		client.Args{
			File:      *iPtr,
			WholeDump: *wPtr,
			ToNATS:    *nPtr,
		})
	failOnErr("Access CSV2JSON Failed: %v", err)

	if fn == "HELP" {
		fPt(str)
	} else {
		m := make(map[string]interface{})
		failOnErr("json.Unmarshal ... %v", json.Unmarshal([]byte(str), &m))
		if *wPtr {
			if m["info"] != nil && m["info"] != "" {
				fPf("INFO: %v\n", m["info"])
			}
			fPln("-----------------------------")
			if m["error"] != nil && m["error"] != "" {
				fPf("ERROR: %v\n", m["error"])
			}
			fPln("-----------------------------")
		}
		if m["data"] != nil && m["data"] != "" {
			fPf("%s\n", m["data"])
		}
	}
}
