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
	cfg := env2Struct("Cfg-Clt-C2J", &Config{}).(*Config)
	spew.Dump(cfg)
}
