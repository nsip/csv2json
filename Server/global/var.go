package global

import (
	cfg "github.com/nsip/n3-csv2json/Server/config"
)

var (
	// Cfg : global variable
	Cfg *cfg.Config
)

// Init : initialize the global variables
func Init(configs ...string) bool {
	configs = append(configs, "./config.toml", "../config.toml", "../../config.toml", "./config/config.toml")
	Cfg = cfg.NewCfg(configs...)
	return Cfg != nil
}
