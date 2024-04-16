package main

import (
	"context"
	"fmt"
	"lms_back/api"
	"lms_back/config"
	"lms_back/pkg/logger"
	"lms_back/service"
	"lms_back/storage/postgres"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.ServiceName)
	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	services := service.New(store, log)
	server := api.New(services, log)


	fmt.Println("programm is running on localhost:8080...")
	server.Run(":8080")
}
