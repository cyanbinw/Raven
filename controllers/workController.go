package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/common"
	"github.com/swirling-melodies/Raven/models/billModels"
	"github.com/swirling-melodies/Raven/work/billNameWork"
	"github.com/swirling-melodies/Raven/work/investmentWork"
	"github.com/swirling-melodies/Raven/work/userWork"
	"net/http"
)

//BillNameSetWorkREPost
// @Tags Work
// @Summary 每月更新数据库后更新BillName表
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/BillNameSetWork [post]
func (WorkRouters) BillNameSetWorkREPost(c *gin.Context) {
	r := new(ReturnData)
	flag, err := billNameWork.SetBillName()
	r.Successful = flag
	if err != nil {
		r.Error = err.Error()
	}
	c.JSON(http.StatusOK, r)
}

//GetBillNameListREPost
// @Tags Work
// @Summary 获取所有BillName
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []billModels.BillNameConfig
// @Router /v1/Work/GetBillNameList [post]
func (WorkRouters) GetBillNameListREPost(c *gin.Context) {
	c.JSON(http.StatusOK, billNameWork.GetBillNameList())
}

//UpdateBillNameREPost
// @Tags Work
// @Summary 更新BillName
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/UpdateBillName [post]
func (WorkRouters) UpdateBillNameREPost(c *gin.Context) {
	var bill = billModels.BillNameConfig{}

	err := c.ShouldBindJSON(&bill)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}
	Helheim.Writer(Helheim.Info, common.ToJSON(bill))
	r := new(ReturnData)
	flag := billNameWork.UpdateBillName(&bill)
	r.Successful = flag
	c.JSON(http.StatusOK, r)
}

//UserSetWorkREPost
// @Tags Work
// @Summary 重置用户表
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/UserSetWork [post]
func (WorkRouters) UserSetWorkREPost(c *gin.Context) {
	r := new(ReturnData)
	flag, err := userWork.SetUser()
	r.Successful = flag
	if err != nil {
		r.Error = err.Error()
	}
	c.JSON(http.StatusOK, r)
}

//InvestmentItemSetWorkREPost
// @Tags Work
// @Summary 添加Investment的下拉菜单
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/UserSetWork [post]
func (WorkRouters) InvestmentItemSetWorkREPost(c *gin.Context) {
	r := new(ReturnData)
	flag, err := investmentWork.SetInvestmentItem()
	r.Successful = flag
	if err != nil {
		r.Error = err.Error()
	}
	c.JSON(http.StatusOK, r)
}

//InvestmentTypeSetWorkREPost
// @Tags Work
// @Summary 添加Investment的下拉菜单
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/UserSetWork [post]
func (WorkRouters) InvestmentTypeSetWorkREPost(c *gin.Context) {
	r := new(ReturnData)
	flag, err := investmentWork.SetInvestmentType()
	r.Successful = flag
	if err != nil {
		r.Error = err.Error()
	}
	c.JSON(http.StatusOK, r)
}

//InvestmentServiceChargeSetWorkREPost
// @Tags Work
// @Summary 添加Investment的下拉菜单
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Work/UserSetWork [post]
func (WorkRouters) InvestmentServiceChargeSetWorkREPost(c *gin.Context) {
	r := new(ReturnData)
	flag, err := investmentWork.SetInvestmentServiceCharge()
	r.Successful = flag
	if err != nil {
		r.Error = err.Error()
	}
	c.JSON(http.StatusOK, r)
}
