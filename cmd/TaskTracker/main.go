package main

import (
	"github.com/gunSlaveUnit/TaskTracker/pkg/handler"
	"github.com/gunSlaveUnit/TaskTracker/pkg/repository"
	"github.com/gunSlaveUnit/TaskTracker/pkg/server"
	"github.com/gunSlaveUnit/TaskTracker/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("ERROR: failed to initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	serv := service.NewService(repos)
	hand := handler.NewHandler(serv)

	s := new(server.Server)
	address := viper.GetString("host") + ":" + viper.GetString("port")
	if err := s.Run(address, hand.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("server")

	return viper.ReadInConfig()
}
