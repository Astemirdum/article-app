package handler

import (
	"github.com/Astemirdum/article-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SingUp(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "SingUp invalid input")
		return
	}
	userId, err := h.services.CreateUser(user)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "SingUp "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"userId": userId})
}

func (h *Handler) SingIn(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "SingIn invalid input")
		return
	}
	token, err := h.services.GenerateToken(user.Email, user.Password)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "SingIn "+err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
