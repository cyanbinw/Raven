package controllers

import (
	"github.com/WFallenDown/Raven/application"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TestHome(c *gin.Context) {
	var billData = application.BillDataByDate{}
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数错误"})
	}
	billData.Year = year
	billData.BillsInitDB()
	billData.BillsGetYearData()
	c.JSON(http.StatusOK, gin.H{"data": billData})
}
