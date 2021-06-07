package Controllers

import (
	"Raven/src/Web/Models/TargetModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTarget(c *gin.Context) {
	TargetModels.TargetInitDB()
	TargetModels.GetTargetData()

	c.JSON(http.StatusOK, nil)
}
