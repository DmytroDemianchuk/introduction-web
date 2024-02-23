package main

import (
	"log"

	"github.com/dmytrodemianchuk/go-auth/internal/config"
	"github.com/dmytrodemianchuk/go-auth/pkg/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := server.NewApp()

	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
