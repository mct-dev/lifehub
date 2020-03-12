package main

import (
	"github.com/mct-dev/lifehub/config"
	"github.com/mct-dev/lifehub/service"
)

func main() {
	config := config.InitConfig()
	server := service.Init(config)
	server.Start()
}
