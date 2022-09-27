package http

import (
	"github.com/gin-gonic/gin"
	"gogle-class/backend/internal/usecase"
	"gogle-class/backend/pkg/auth"
)

type Handler struct {
	useCases     *usecase.UseCases
	tokenManager auth.TokenManager
}

func NewHandler(services *usecase.UseCases, tokenManager auth.TokenManager) *Handler {
	return &Handler{
		useCases:     services,
		tokenManager: tokenManager,
	}
}

func (h *Handler) InitRouter(port string) error {
	router := gin.Default()

	auth := router.Group("/auth")

	auth.POST("/register", h.register)
	auth.POST("/login", h.login)
	auth.POST("/refresh", h.refresh)

	return router.Run(port)
}
