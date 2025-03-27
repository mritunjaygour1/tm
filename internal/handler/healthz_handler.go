package handler

import (
	"net/http"
	"task-manager/internal/model"

	"github.com/gin-gonic/gin"
)

type (
	IHealthzHandler interface {
		GetHealthz(*gin.Context)
	}

	HealthzHandler struct {
	}
)

func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

/*
	Handler functions
*/

func (h *HealthzHandler) GetHealthz(c *gin.Context) {
	c.JSON(http.StatusOK, &model.Response{Message: "OK"})
}

/*
	Normal Suporting functions
*/
