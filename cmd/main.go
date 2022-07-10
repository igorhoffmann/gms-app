package main

import (
	"github.com/igorgofman/GMS-app"
	"github.com/igorgofman/GMS-app/pkg/handler"
	"github.com/igorgofman/GMS-app/pkg/repository"
	"github.com/igorgofman/GMS-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository() //db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(gym.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
