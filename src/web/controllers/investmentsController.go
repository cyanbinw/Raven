package controllers

import (
	"Raven/src/application"
	"Raven/src/log"
	"Raven/src/models/investmentsModels"
	service2 "Raven/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetInvestments
// @Tags Investment
// @Summary 获取Investments的金额分类
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} application.InvestmentsChartModel {}
// @Router /v1/Investment/GetInvestments [post]
func GetInvestments(c *gin.Context) {
	var investmentData = application.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.InvestmentChartForAccount())
}

//GetInvestmentsTable
// @Tags Investment
// @Summary 获取Investments的表数据
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []investmentsModels.InvestmentTable {}
// @Router /v1/Investment/GetInvestmentsTable [post]
func GetInvestmentsTable(c *gin.Context) {
	var investmentData = application.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.GetInvestmentTable())
}

//AddInvestmentsTable
// @Tags Investment
// @Summary 添加新数据
// @Description 描述信息
// @Param user body investmentsModels.InvestmentTable true "investmentData"
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/AddInvestmentsTable [post]
func AddInvestmentsTable(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	log.Writer(log.Info, service2.ToJSON(investmentData.Investment))

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	application.InvestmentsInitDB()
	flag, err := application.AddInvestmentTable(&investmentData)

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, ReturnData{Successful: flag})
}

//UpdateInvestmentsTable
// @Tags Investment
// @Summary 更新一条数据
// @Description 描述信息
// @Param user body investmentsModels.InvestmentTable true "investmentData"
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Failure 400 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/UpdateInvestmentsTable [post]
func UpdateInvestmentsTable(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	log.Writer(log.Info, service2.ToJSON(investmentData.Investment))

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	application.InvestmentsInitDB()
	flag, err := application.UpdateInvestmentTable(&investmentData)

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, ReturnData{Successful: flag})
}

//GetInvestmentDiagram
// @Tags Investment
// @Summary 获取图表信息
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} map[string][]investmentsModels.Investment
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/GetInvestmentDiagram [post]
func GetInvestmentDiagram(c *gin.Context) {

	data, err := application.GetInvestmentDiagram()

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{
			Successful: false,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

//GetInvestmentOption
// @Tags Investment
// @Summary 获取查询条件信息(table page)
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} application.InvestmentOption
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/GetInvestmentDiagram [post]
func GetInvestmentOption(c *gin.Context) {
	data, err := application.GetInvestmentOption()

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, data)
}
