package main

import (
	"runtime"

	"log"

	"github.com/melodiez14/lastcake/src/util/conn"
	"github.com/melodiez14/lastcake/src/util/env"
	"github.com/melodiez14/lastcake/src/util/jsonconfig"
	"github.com/melodiez14/lastcake/src/util/webserver"
)

type configuration struct {
	Database  conn.DatabaseConfig `json:"database"`
	Webserver webserver.Config    `json:"webserver"`
}

func init() {
	// use all CPU core
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	cfgenv := env.Get()
	config := &configuration{}
	isLoaded := jsonconfig.Load(&config, "/etc/lastcake", cfgenv) || jsonconfig.Load(&config, "./etc/lastcake", cfgenv)
	if !isLoaded {
		log.Fatal("Cannot load configuration")
	}

	conn.InitDB(config.Database)
	webserver.Start(config.Webserver)
}
