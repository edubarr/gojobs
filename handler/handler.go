package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET Openings",
	})
}
func GetOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
func CreateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "POST Opening",
	})
}
func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "DELETE Opening",
	})
}
func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
