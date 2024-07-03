package main

import (
	"log"
	parser "main-server"
	handler "main-server/pkg/handlers"
)

func main() {
	handlers := handler.Handler{}
	srv := new(parser.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured")
	}
}
