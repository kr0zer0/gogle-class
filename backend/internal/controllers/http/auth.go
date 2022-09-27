package http

import (
	"github.com/gin-gonic/gin"
	"gogle-class/backend/internal/controllers/http/dto"
	"gogle-class/backend/internal/domain"
	"net/http"
)

func (h *Handler) register(c *gin.Context) {
	var input dto.Registration

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user := domain.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err = h.useCases.RegisterUseCase.Register(c, &user)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "created"})
}

func (h *Handler) login(c *gin.Context) {
	var input dto.Login

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.useCases.LoginUseCase.Login(c, input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}

func (h *Handler) refresh(c *gin.Context) {
	var input dto.RefreshInput

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	tokens, err := h.useCases.RefreshTokensUseCase.RefreshUserTokens(c, input.Token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
