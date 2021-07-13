package controllers

import (
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
