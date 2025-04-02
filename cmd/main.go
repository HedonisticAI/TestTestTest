package main

import (
	"testtesttest/config"
	"testtesttest/pkg/logger"
	"testtesttest/pkg/postgres"
)

func main() {
	Logger := logger.NewLogger()
	Config := config.NewConfig()
	if Config == nil {
		Logger.Info("Bad config")
	}
	DB, err := postgres.NewDB(*Config)
	if err != nil {
		Logger.Error(err.Error())
	}
	defer DB.Close()
}
