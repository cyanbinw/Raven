package controllers

import (
	"github.com/WFallenDown/Helheim"
	"github.com/WFallenDown/Raven/application"
	"github.com/WFallenDown/Raven/models/billModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Login
// @Tags User
// @Summary 登陆
// @Description 描述信息
// @Param user body billModels.BillTable true "investmentData"
// @Security Bearer
// @Produce  json
// @Success 200 {object} billModels.BillTable
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Bill/GetBillsTable [post]
func Login(c *gin.Context) {
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
