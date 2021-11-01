package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swirling-melodies/Raven/models/desireModels"
	"net/http"
)

func (DesireReuter) GetDesireREPost(c *gin.Context) {
	desireModels.GetDesire()

	c.JSON(http.StatusOK, nil)
}
