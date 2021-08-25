package controllers

import (
	"github.com/WFallenDown/Helheim"
	"github.com/WFallenDown/Raven/models/billModels"
	"github.com/WFallenDown/Raven/service"
	"github.com/WFallenDown/Raven/work/billNameWork"
	"github.com/WFallenDown/Raven/work/userWork"
	"github.com/gin-gonic/gin"
	"net/http"
)

//BillNameSetWork
// @Tags Work
// @Summary 每月更新数据库后更新BillName表
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/BillNameSetWork [post]
func BillNameSetWork(c *gin.Context) {
	r := new(ReturnData)
	flag, err := billNameWork.SetBillName()
	r.Successful = flag
	if err != nil {
		r.Error = err.Error()
	}
	c.JSON(http.StatusOK, r)
}

//GetBillNameList
// @Tags Work
// @Summary 获取所有BillName
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []billModels.BillNameConfig
// @Router /v1/Work/GetBillNameList [post]
func GetBillNameList(c *gin.Context) {
	c.JSON(http.StatusOK, billNameWork.GetBillNameList())
}

//UpdateBillName
// @Tags Work
// @Summary 更新BillName
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/UpdateBillName [post]
func UpdateBillName(c *gin.Context) {
	var bill = billModels.BillNameConfig{}

	err := c.ShouldBindJSON(&bill)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}
	Helheim.Writer(Helheim.Info, service.ToJSON(bill))
	r := new(ReturnData)
	flag := billNameWork.UpdateBillName(&bill)
	r.Successful = flag
	c.JSON(http.StatusOK, r)
}

//UserSetWork
// @Tags Work
// @Summary 重置用户表
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/UserSetWork [post]
func UserSetWork(c *gin.Context) {
	r := new(ReturnData)
	flag, err := userWork.SetUser()
	r.Successful = flag
	if err != nil {
		r.Error = err.Error()
	}
	c.JSON(http.StatusOK, r)
}
