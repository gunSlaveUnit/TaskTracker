package main

import (
	"github.com/gunSlaveUnit/TaskTracker/pkg/handler"
	"github.com/gunSlaveUnit/TaskTracker/pkg/repository"
	"github.com/gunSlaveUnit/TaskTracker/pkg/server"
	"github.com/gunSlaveUnit/TaskTracker/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	serv := service.NewService(repos)
	hand := handler.NewHandler(serv)

	s := new(server.Server)
	if err := s.Run("localhost:9876", hand.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}
