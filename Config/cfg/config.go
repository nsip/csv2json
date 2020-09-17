package cfg

import "github.com/cdutwhu/n3-util/n3cfg"

// Config : AUTO Created From /home/qmiao/Desktop/temp/n3-csv2json/Config/config.toml
type Config struct {
	Log string
	Service interface{}
	Version string
	Loggly struct {
		Token string
	}
	WebService struct {
		Port int
	}
	NATS struct {
		URL string
		Subject string
		Timeout int
	}
	Route struct {
		Help string
		ToCSV string
		ToJSON string
	}
	Server struct {
		IP interface{}
		Port int
		Protocol string
	}
	Access struct {
		Timeout int
	}
}

// NewCfg :
func NewCfg(cfgStruName string, mReplExpr map[string]string, cfgPaths ...string) interface{} {
	var cfg interface{}
	switch cfgStruName {
	case "Config":
		cfg = &Config{}
	default:
		return nil
	}
	return n3cfg.InitEnvVar(cfg, mReplExpr, cfgStruName, cfgPaths...)
}
