package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
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
	userId, err := h.services.ValidateToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "incorrect token value")
	}
	fmt.Println(userId)
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
