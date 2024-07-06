package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Printf("Status code: %d\nmessage: %s\n", statusCode, message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
