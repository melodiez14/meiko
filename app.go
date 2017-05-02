package main

import (
	"runtime"

	"log"

	"github.com/melodiez14/meiko/src/util/conn"
	"github.com/melodiez14/meiko/src/util/env"
	"github.com/melodiez14/meiko/src/util/jsonconfig"
	"github.com/melodiez14/meiko/src/util/webserver"
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
	isLoaded := jsonconfig.Load(&config, "/etc/meiko", cfgenv) || jsonconfig.Load(&config, "./etc/meiko", cfgenv)
	if !isLoaded {
		log.Fatal("Cannot load configuration")
	}

	conn.InitDB(config.Database)
	webserver.Start(config.Webserver)
}
