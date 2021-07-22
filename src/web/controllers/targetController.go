package controllers

import (
	"github.com/WFallenDown/Raven/src/web/models/targetModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTarget(c *gin.Context) {
	targetModels.TargetInitDB()
	targetModels.GetTargetData()

	c.JSON(http.StatusOK, nil)
}
