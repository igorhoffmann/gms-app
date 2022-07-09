package main

import (
	"log"

	"github.com/igorgofman/GMS-app"
	"github.com/igorgofman/GMS-app/pkg/handler"
	"github.com/spf13/viper"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(gymSys.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}
