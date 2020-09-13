package main

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/nsip/n3-csv2json/Config/cfg"
)

func TestMain(t *testing.T) {
	main()
}

func TestLoad(t *testing.T) {
	c := cfg.NewCfg(
		"Config",
		map[string]string{
			"[s]": "Service",
			"[v]": "Version",
		},
		"../Config/config.toml",
	).(*cfg.Config)
	spew.Dump(*c)
}

func TestInit(t *testing.T) {
	c := cfg.NewCfg(
		"Config",
		map[string]string{
			"[s]":    "Service",
			"[v]":    "Version",
			"[port]": "WebService.Port",
		},
		"../Config/config.toml",
	).(*cfg.Config)
	spew.Dump(*c)

	c = env2Struct("Config", &cfg.Config{}).(*cfg.Config)
	spew.Dump(*c)
}
