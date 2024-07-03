package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
