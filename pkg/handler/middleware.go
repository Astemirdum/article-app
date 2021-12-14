package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	idKey      = "userId"
	headerAuth = "Authorization"
)

func (h *Handler) UserIdentification(ctx *gin.Context) {
	headerToken := ctx.GetHeader(headerAuth)
	if headerToken == "" {
		ErrorResponse(ctx, http.StatusUnauthorized, "empty header Authorization")
		return
	}
	//headerToken = strings.Trim(headerToken, "\"")
	headerTokenParts := strings.Split(headerToken, " ")
	//log.Println(headerTokenParts)
	if len(headerTokenParts) != 2 || headerTokenParts[0] != "Bearer" {
		ErrorResponse(ctx, http.StatusUnauthorized, "invalid header Authorization")
		return
	}
	token := headerTokenParts[1]
	if len(token) == 0 {
		ErrorResponse(ctx, http.StatusUnauthorized, "empty header Authorization Token")
		return
	}
	userId, err := h.services.ParseToken(token)
	if err != nil {
		ErrorResponse(ctx, http.StatusUnauthorized, "token invalid "+err.Error())
		return
	}
	ctx.Set(idKey, userId)
}

func getUserId(ctx *gin.Context) (int, error) {
	key, ok := ctx.Get(idKey)
	if !ok {
		return 0, errors.New("user doesnt exist with this token")
	}
	userId, ok := key.(int)
	if !ok {
		return 0, errors.New("userId not int")
	}
	return userId, nil
}
