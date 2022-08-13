package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

type statusResponse struct {
	Status string `json:"status"`
}

type TokenResponse struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}
