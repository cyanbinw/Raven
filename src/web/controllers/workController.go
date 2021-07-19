package controllers

import (
	"Raven/src/log"
	"Raven/src/models/billModels"
	service2 "Raven/src/service"
	"Raven/src/work/billNameWork"
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
	r.Error = err.Error()
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
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}
	log.Writer(log.Info, service2.ToJSON(bill))
	r := new(ReturnData)
	flag := billNameWork.UpdateBillName(&bill)
	r.Successful = flag
	c.JSON(http.StatusOK, r)
}
