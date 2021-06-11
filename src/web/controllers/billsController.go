package controllers

import (
	"Raven/src/log"
	"Raven/src/web/models/billModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBillsYearAllData(c *gin.Context) {
	var billData = billModels.BillDataByDate{}

	billData.BillsInitDB()
	billData.BillsGetYearData()
	c.JSON(http.StatusOK, billData.Data)
}

func GetBillsDataByMonth(c *gin.Context) {
	var billData = billModels.BillDataByDate{}

	billData.BillsInitDB()
	billData.BillsGetDataByMonth()
	c.JSON(http.StatusOK, billData.Data)
}

func GetBillsTable(c *gin.Context) {
	var bill = billModels.BillTable{}

	err := c.ShouldBindJSON(&bill)
	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数错误", "error": err})
		return
	}

	c.JSON(http.StatusOK, billModels.BillsGetTable(&bill))
}

func GetBillsDiagram(c *gin.Context) {
	// c.JSON(http.StatusOK, billModels.BillsGetAll())
}
