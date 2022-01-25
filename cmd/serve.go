package cmd

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/handler"
)

var (
	conf *Config
)

func init() {
	loadConfig()
	conf = initConfig()
}

func Execute() {
	//register all server needs, db,repo, etc

	api := handler.NewAPI()
	api.Start(conf.Host, conf.Port)
}
