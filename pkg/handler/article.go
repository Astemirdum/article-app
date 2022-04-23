package handler

import (
	"net/http"
	"strconv"

	"github.com/Astemirdum/article-app/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllArticles(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		h.log.Errorf("getUserId fetch userId %v", err)
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	articles, err := h.services.SelectAll(userId)
	if err != nil {
		h.log.Errorf("selectAll articles %v", err)
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"all articles": articles})
}

func (h *Handler) GetOneArticle(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	artId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	art, err := h.services.Get(userId, artId)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, art)
}

func (h *Handler) CreateArticle(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	var art models.Article

	if err := ctx.BindJSON(&art); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "CreateArticle json "+err.Error())
		return
	}
	artId, err := h.services.Create(userId, art)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	h.log.Infof("article has been created %d by user %d", artId, userId)
	ctx.JSON(http.StatusOK, gin.H{"article Id": artId})
}

func (h *Handler) DeleteArticle(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	artId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	if err := h.services.Delete(userId, artId); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.log.Infof("article has been deleted %d by user %d", artId, userId)
	ctx.JSON(http.StatusOK, gin.H{"delete": "ok"})
}

func (h *Handler) UpdateArticle(ctx *gin.Context) {
	userId, err := getUserId(ctx)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	artId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}
	var upArt models.UpdateArticle

	if err := ctx.BindJSON(&upArt); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(userId, artId, upArt); err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.log.Infof("article has been updated %d by user %d", artId, userId)
	ctx.JSON(http.StatusOK, gin.H{"update": "ok"})
}
