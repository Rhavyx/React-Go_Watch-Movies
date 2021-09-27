package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) statusHandler(c *gin.Context) {

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     app.config.version,
	}

	c.JSON(http.StatusOK, currentStatus)

}
