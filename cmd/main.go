package main

import (
	"log"
	parser "main-server"
	handler "main-server/pkg/handlers"
	"main-server/pkg/repository"
	"main-server/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(parser.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured")
	}
}
