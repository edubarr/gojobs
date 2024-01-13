package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
