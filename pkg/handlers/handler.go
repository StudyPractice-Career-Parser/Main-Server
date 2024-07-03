package handler

import (
	"main-server/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{service: serv}
}

func (e *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	search := router.Group("/search")
	{
		search.GET("/vacancies", e.getVacancies)
		search.GET("/resume", e.getResume)
	}

	return router
}
