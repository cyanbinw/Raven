package controllers

import (
	"github.com/WFallenDown/Helheim"
	"github.com/WFallenDown/Raven/src/web/application"
	"github.com/WFallenDown/Raven/src/web/models/billModels"
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
	var billData = application.BillDataByDate{}

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
	var billData = application.BillDataByDate{}

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
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	if !bill.DateMin.IsZero() {
		bill.DateMin = bill.DateMin.Local()
	}
	if !bill.DateMax.IsZero() {
		bill.DateMax = bill.DateMax.Local()
	}

	c.JSON(http.StatusOK, application.BillsGetTable(&bill))
}

//GetBillsDataByPage
// @Tags Bill
// @Summary 根据页面获取bill的数据
// @Description 描述信息
// @Param user body billModels.BillTable true "investmentData"
// @Security Bearer
// @Produce  json
// @Success 200 {object} billModels.BillDataByPage
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Bill/GetBillsAllData [post]
func GetBillsDataByPage(c *gin.Context) {
	var bill = billModels.BillDataByPage{}

	err := c.ShouldBindJSON(&bill)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, application.BillsGetDataByPage(&bill))
}

//GetBillsTableOption
// @Tags Bill
// @Summary 获取bills表查询条件
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} application.BillOption
// @Router /v1/Bill/GetBillsTable [post]
func GetBillsTableOption(c *gin.Context) {
	c.JSON(http.StatusOK, application.BillsGetTableOption())
}

//GetBillsDiagram
// @Tags Bill
// @Summary 获取bills表信息
// @Description 描述信息
// @Param user body billModels.BillTable true "investmentData"
// @Security Bearer
// @Produce  json
// @Success 200 {object} billModels.BillTable
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Bill/GetBillsDiagram [post]
func GetBillsDiagram(c *gin.Context) {
	var bill = billModels.BillTable{}

	err := c.ShouldBindJSON(&bill)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	data, err := application.BillsGetDiagram(&bill)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Message: "查询错误", Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, data)
}
