package Controllers

import (
	"Raven/src/Web/Models/BillModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetYearAllData(c *gin.Context) {
	var billData = BillModels.BillData{}

	billData.BillsInitDB()
	billData.BillsGetYearData()
	c.JSON(http.StatusOK, billData.Data)
}

func GetFourMonthData(c *gin.Context) {
	var billData = BillModels.BillData{}

	billData.BillsInitDB()
	billData.BillsGetFourMonthsData()
	c.JSON(http.StatusOK, billData.Data)
}
