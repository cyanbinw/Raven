package controllers

import (
	"github.com/WFallenDown/Raven/models/desireModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDesire(c *gin.Context) {
	desireModels.GetDesire()

	c.JSON(http.StatusOK, nil)
}
