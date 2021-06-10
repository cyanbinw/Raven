package controllers

import (
	"Raven/src/web/models/billModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetYearAllData(c *gin.Context) {
	var billData = billModels.BillDataByDate{}

	billData.BillsInitDB()
	billData.BillsGetYearData()
	c.JSON(http.StatusOK, billData.Data)
}

func GetDataByMonth(c *gin.Context) {
	var billData = billModels.BillDataByDate{}

	billData.BillsInitDB()
	billData.BillsGetDataByMonth()
	c.JSON(http.StatusOK, billData.Data)
}

func GetAllData(c *gin.Context) {
	c.JSON(http.StatusOK, billModels.BillsGetAll())
}
