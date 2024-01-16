package handler

import (
	"github.com/edubarr/gojobs/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListOpeningsHandler
// @BasePath /api/v1
// @Summary List all job openings
// @Description List all job openings
// @Tags openings
// @Accept  json
// @Produce  json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(context *gin.Context) {
	openings := make([]schemas.Opening, 0)

	if err := db.Find(&openings).Error; err != nil {
		sendError(context, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(context, "list-openings", openings)
}
