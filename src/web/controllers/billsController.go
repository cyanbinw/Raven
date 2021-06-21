package controllers

import (
	"Raven/src/log"
	"Raven/src/web/models/billModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetBillsYearAllData
// @Tags Bill
// @Summary 获取最近一年的bills
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []billModels.BillDetail
// @Router /v1/Bill/GetBillsYearAllData [post]
func GetBillsYearAllData(c *gin.Context) {
	var billData = billModels.BillDataByDate{}

	billData.BillsInitDB()
	billData.BillsGetYearData()
	c.JSON(http.StatusOK, billData.Data)
}

//GetBillsDataByMonth
// @Tags Bill
// @Summary 获取最近四个月的bills
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []billModels.BillDetail
// @Router /v1/Bill/GetBillsDataByMonth [post]
func GetBillsDataByMonth(c *gin.Context) {
	var billData = billModels.BillDataByDate{}

	billData.BillsInitDB()
	billData.BillsGetDataByMonth()
	c.JSON(http.StatusOK, billData.Data)
}

//GetBillsTable
// @Tags Bill
// @Summary 获取bills表信息
// @Description 描述信息
// @Param user body billModels.BillTable true "investmentData"
// @Security Bearer
// @Produce  json
// @Success 200 {object} billModels.BillTable
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Bill/GetBillsTable [post]
func GetBillsTable(c *gin.Context) {
	var bill = billModels.BillTable{}

	err := c.ShouldBindJSON(&bill)
	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	if !bill.DateMin.IsZero() {
		bill.DateMin = bill.DateMin.Local()
	}
	if !bill.DateMax.IsZero() {
		bill.DateMax = bill.DateMax.Local()
	}

	c.JSON(http.StatusOK, billModels.BillsGetTable(&bill))
}

//GetBillsTableOption
// @Tags Bill
// @Summary 获取bills表查询条件
// @Description 描述信息
// @Param user body billModels.BillTable true "investmentData"
// @Security Bearer
// @Produce  json
// @Success 200 {object} billModels.BillOption
// @Router /v1/Bill/GetBillsTable [post]
func GetBillsTableOption(c *gin.Context) {
	c.JSON(http.StatusOK, billModels.BillsGetTableOption())
}

func GetBillsDiagram(c *gin.Context) {
	// c.JSON(http.StatusOK, billModels.BillsGetAll())
}
