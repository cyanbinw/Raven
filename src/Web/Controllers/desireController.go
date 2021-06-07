package Controllers

import (
	"Raven/src/Web/Models/DesireModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDesire(c *gin.Context) {
	DesireModels.GetDesire()

	c.JSON(http.StatusOK, nil)
}
