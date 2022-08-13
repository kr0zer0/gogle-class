package handler

import (
	"github.com/gin-gonic/gin"
	"gogle-class/backend/internal/service"
	"gogle-class/backend/pkg/auth"
)

type Handler struct {
	services     *service.Service
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Service, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) InitRouter(port string) error {
	router := gin.Default()

	auth := router.Group("/auth")

	auth.POST("/register", h.register)
	auth.POST("/login", h.login)
	auth.POST("/refresh", h.refresh)

	//router.GET("/ping", h.authMiddleware, func(c *gin.Context) {
	//	id, _ := c.Get("userID")
	//	c.JSON(200, id)
	//})

	return router.Run(port)
}
