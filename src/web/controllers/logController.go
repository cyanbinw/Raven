package controllers

import (
	"fmt"
	"github.com/WFallenDown/Helheim"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetLogTable
// @Tags Log
// @Summary 查看Log
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []Helheim.Record
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Bill/GetLogTable [post]
func GetLogTable(c *gin.Context) {
	var data = Helheim.RecordList{}
	if err := c.ShouldBindJSON(&data); err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	if err := Helheim.GetLog(&data); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, data.Data)
}
