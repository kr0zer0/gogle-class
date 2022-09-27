package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) authMiddleware(c *gin.Context) {
	id, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set("userID", id)
	c.Next()
}

func (h *Handler) parseAuthHeader(c *gin.Context) (string, error) {
	header := c.GetHeader("Authorization")

	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("empty auth header")
	}

	return h.tokenManager.ParseToken(headerParts[1])

}
