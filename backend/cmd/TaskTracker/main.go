package main

import (
	"log"

	"github.com/gunSlaveUnit/TaskTracker/pkg/db"
	"github.com/gunSlaveUnit/TaskTracker/pkg/handler"
	"github.com/gunSlaveUnit/TaskTracker/pkg/repository"
	"github.com/gunSlaveUnit/TaskTracker/pkg/server"
	"github.com/gunSlaveUnit/TaskTracker/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("ERROR: failed to initializing config: %s", err.Error())
	}

	database, err := db.NewDB(&db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.Get("username").(string),
		Password: viper.Get("password").(string),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})

	if err != nil {
		log.Fatalf("ERROR: failed to connect to the database: %s", err.Error())
	}

	repos := repository.NewRepository(database)
	serv := service.NewService(repos)
	hand := handler.NewHandler(serv)

	s := new(server.Server)
	address := viper.GetString("server.host") + ":" + viper.GetString("server.port")
	if err := s.Run(address, hand.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")

	viper.SetConfigName("server")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetConfigName("db")
	if err := viper.MergeInConfig(); err != nil {
		return err
	}

	viper.AddConfigPath("envs")

	viper.SetConfigName(".db")
	viper.SetConfigType("env")

	return viper.MergeInConfig()
}
