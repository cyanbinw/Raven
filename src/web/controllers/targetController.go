package controllers

import (
	targetModels2 "Raven/src/models/targetModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTarget(c *gin.Context) {
	targetModels2.TargetInitDB()
	targetModels2.GetTargetData()

	c.JSON(http.StatusOK, nil)
}
