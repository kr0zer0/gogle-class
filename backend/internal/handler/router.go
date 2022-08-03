package handler

import (
	"github.com/gin-gonic/gin"
	"gogle-class/backend/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRouter(port string) error {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, "NICE")
	})

	return router.Run(port)
}
