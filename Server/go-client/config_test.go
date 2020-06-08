package client

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestLoad(t *testing.T) {
	cfg := newCfg("./config.toml")
	spew.Dump(cfg)
}

func TestInit(t *testing.T) {
	initEnvVarFromTOML("Cfg-Clt-C2J", "./config.toml")
	Icfg, err := env2Struct("Cfg-Clt-C2J", &Config{})
	failOnErr("%v", err)
	cfg := Icfg.(*Config)
	spew.Dump(cfg)
}
