package handler

import (
	"net/http"

	"github.com/Astemirdum/article-app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SingUp(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "singUp invalid input")
		return
	}
	userId, err := h.services.CreateUser(user)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "singUp "+err.Error())
		return
	}
	h.log.Infof("sign up user id = %d", userId)
	ctx.JSON(http.StatusOK, gin.H{"userId": userId})
}

func (h *Handler) SingIn(ctx *gin.Context) {
	var user models.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "singIn invalid input")
		return
	}
	token, err := h.services.GenerateToken(user.Email, user.Password)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "singIn "+err.Error())
		return
	}

	h.log.Infof("signIn user %d with token %s", user.Id, token)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
