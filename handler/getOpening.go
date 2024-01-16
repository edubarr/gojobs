package handler

import (
	"fmt"
	"github.com/edubarr/gojobs/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetOpeningHandler
// @BasePath /api/v1
// @Summary Get a job opening
// @Description Get the job opening with the id supplied
// @Tags openings
// @Accept  json
// @Produce  json
// @Param id query string true "Opening identification"
// @Success 200 {object} GetOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func GetOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id %v not found", id))
		return
	}

	sendSuccess(context, "get-opening", opening)
}
