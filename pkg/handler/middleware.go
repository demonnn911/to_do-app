package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userId, err := h.services.ValidateToken(ctx, headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect token value")
	}
	c.Set(userCtx, userId)
}

func (h *Handler) getUserId(c *gin.Context) (int64, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int64)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "can not set int64 type to user's id(invalid type)")
		return 0, errors.New("can not set int64 type to user's id(invalid type)")
	}
	return idInt, nil
}
