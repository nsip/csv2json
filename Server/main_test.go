package main

import (
	"testing"

	"github.com/cdutwhu/n3-util/n3cfg"
	"github.com/cdutwhu/n3-util/n3err"
	"github.com/davecgh/go-spew/spew"
)

func TestMain(t *testing.T) {
	main()
}

func TestLoad(t *testing.T) {
	cfg := n3cfg.ToEnvN3csv2jsonServer(nil, "TestKey", "./config.toml")
	failOnErrWhen(cfg == nil, "%v", n3err.CFG_INIT_ERR)
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	cfg := n3cfg.ToEnvN3csv2jsonServer(nil, "TestKey", "./config.toml")
	failOnErrWhen(cfg == nil, "%v", n3err.CFG_INIT_ERR)
	cfg1 := n3cfg.FromEnvN3csv2jsonServer("TestKey")
	spew.Dump(cfg1)
}
