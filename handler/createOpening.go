package handler

import (
	"github.com/edubarr/gojobs/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

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

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		sendError(context, http.StatusInternalServerError, err.Error())
		logger.Errorf("error creating opening: %v", err.Error())
		return
	}

	sendSuccess(context, "create-opening", opening)
}
