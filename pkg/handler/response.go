package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type errorResp struct {
	Mess string `json:"message"`
}

func ErrorResponse(ctx *gin.Context, stCode int, message string) {
	log.Error(message)
	ctx.AbortWithStatusJSON(stCode, errorResp{message})
}
