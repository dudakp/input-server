package main

import (
	"github.com/dudakp/input-server/cmd/session/app"
	"github.com/dudakp/input-server/internal/config"
)

type Config struct {
	*config.AppConfig

	GrpcPort int `mapstructure:"GRPC_PORT" default:"50051" flag:"gprcPort"`
}

func main() {
	conf := &Config{}
	defaultValues, err := config.ExtractConfigValues(conf)
	appConfig := config.NewAppConfig("SESSION", defaultValues)
	conf.AppConfig = appConfig

	err = conf.InitializeAppConfig()
	if err != nil {
		panic(err)
	}
	err = conf.GetConfigValues(conf)
	if err != nil {
		panic(err)
	}
	app.StartTypingServer(conf.GrpcPort)
}
