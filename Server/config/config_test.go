package config

import "testing"

func TestLoad(t *testing.T) {
	cfg := newCfg("./config.toml")
	fPln(cfg.Path)
	fPln(cfg.LogFile)
	fPln(cfg.ServiceName)
	fPln(cfg.WebService)
	fPln(cfg.Route)
	fPln(cfg.NATS)
	fPln(cfg.File)
}

func TestInit(t *testing.T) {
	InitEnvVarFromTOML("Cfg", "./config.toml")
	cfg := env2Struct("Cfg", &Config{}).(*Config)
	fPln(cfg.Path)
	fPln(cfg.LogFile)
	fPln(cfg.ServiceName)
	fPln(cfg.WebService)
	fPln(cfg.Route)
	fPln(cfg.NATS)
	fPln(cfg.File)
}
