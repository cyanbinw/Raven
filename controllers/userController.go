package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/application"
	"net/http"
)

//Login
// @Tags User
// @Summary 登陆
// @Description 描述信息
// @Param user body userModel.UserInfo true "User"
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true, "data":null, "Error":"", "Message":""}
// @Failure 400 {object} ReturnData {"Successful":true, "data":null, "Error":"", "Message":""}
// @Router /v1/Bill/GetBillsTable [post]
func Login(c *gin.Context) {
	var user = application.NewUser()

	err := c.ShouldBindJSON(&user)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}
	rd := new(ReturnData)
	rd.Successful, err = user.Login()
	if err != nil {
		rd.Error = err.Error()
	}
	rd.Data = user.TokenNum
	c.JSON(http.StatusOK, rd)
}

//ValidateToken
// @Tags User
// @Summary 登陆
// @Description 描述信息
// @Param user body userModel.Token true "Token"
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true, "data":null, "Error":"", "Message":""}
// @Failure 400 {object} ReturnData {"Successful":true, "data":null, "Error":"", "Message":""}
// @Router /v1/Bill/GetBillsTable [post]
func ValidateToken(c *gin.Context) {
	var user = application.NewUser()

	err := c.ShouldBindJSON(&user)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}
	rd := new(ReturnData)
	rd.Successful, err = user.ValidateToken()
	if err != nil {
		rd.Error = err.Error()
	}
	c.JSON(http.StatusOK, rd)
}
