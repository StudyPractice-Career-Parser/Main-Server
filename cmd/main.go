package main

import (
	"log"
	parser "main-server"
	handler "main-server/pkg/handlers"
	"main-server/pkg/repository"
	"main-server/pkg/service"

	"github.com/spf13/viper"
)

// @title           GetVacancies
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8000
// @BasePath  /
func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error occured while reading config")
	}
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(parser.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured when server runs")
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
