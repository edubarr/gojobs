package handler

import (
	"fmt"
	"github.com/edubarr/gojobs/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteOpeningHandler
// @BasePath /api/v1
// @Summary Delete a job opening
// @Description Delete the job opening with the id supplied
// @Tags openings
// @Accept  json
// @Produce  json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(context *gin.Context) {
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

	// Delete opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %v", id))
		return
	}

	sendSuccess(context, "delete-opening", opening)
}
