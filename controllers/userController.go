package controllers

import (
	"github.com/WFallenDown/Helheim"
	"github.com/WFallenDown/Raven/application"
	"github.com/gin-gonic/gin"
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
	var user = application.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}
	user.UserInitDB()
	rd := new(ReturnData)
	rd.Successful, err = user.Login()
	rd.Error = err.Error()
	rd.Data = struct {
		Token       [16]byte
		UpdateToken [16]byte
	}{
		Token:       user.TokenNum,
		UpdateToken: user.UpdateTokenNum,
	}
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
	var user = application.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}
	user.UserInitDB()
	rd := new(ReturnData)
	rd.Successful, err = user.ValidateToken()
	rd.Error = err.Error()
	c.JSON(http.StatusOK, rd)
}
