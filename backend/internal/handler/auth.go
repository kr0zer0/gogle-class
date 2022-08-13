package handler

import (
	"github.com/gin-gonic/gin"
	"gogle-class/backend/internal/command"
	"gogle-class/backend/internal/domain"
	"net/http"
)

func (h *Handler) register(c *gin.Context) {
	var input command.Registration

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user := domain.User{
		Email:    input.Email,
		Password: input.Password,
	}

	err = h.services.Auth.Registration(c, &user)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{Status: "created"})
}

func (h *Handler) login(c *gin.Context) {
	var input command.Login

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	tokens, err := h.services.Auth.Login(c, input.Email, input.Password)
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
	var input command.RefreshInput

	err := c.BindJSON(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	tokens, err := h.services.Auth.RefreshUserTokens(c, input.Token)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	})
}
