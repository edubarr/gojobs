package handler

import (
	"github.com/edubarr/gojobs/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := make([]schemas.Opening, 0)

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(context, "list-openings", openings)
}
