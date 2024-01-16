package handler

import (
	"fmt"
	"github.com/edubarr/gojobs/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(context *gin.Context) {
	id := context.Query("id")
	request := UpdateOpeningRequest{}

	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	if err := context.BindJSON(&request); err != nil {
		sendError(context, http.StatusBadRequest, err.Error())
		logger.Error("JSON binding error")
		return
	}

	if err := request.Validate(); err != nil {
		sendError(context, http.StatusBadRequest, err.Error())
		logger.Errorf("validation error: %v", err.Error())
		return
	}

	opening := schemas.Opening{}

	// Find opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, fmt.Sprintf("opening with id %v not found", id))
		return
	}

	// Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, fmt.Sprintf("error updating opening with id: %v", id))
		logger.Errorf("error updating opening with id: %v", id)
		return
	}
	sendSuccess(context, "update-opening", opening)
}
