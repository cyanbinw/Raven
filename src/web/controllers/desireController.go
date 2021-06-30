package controllers

import (
	desireModels2 "Raven/src/models/desireModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDesire(c *gin.Context) {
	desireModels2.GetDesire()

	c.JSON(http.StatusOK, nil)
}
