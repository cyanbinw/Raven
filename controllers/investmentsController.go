package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/application"
	"github.com/swirling-melodies/Raven/common"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"net/http"
)

//GetInvestmentsREPost
// @Tags Investment
// @Summary 获取Investments的金额分类
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} application.InvestmentsChartModel {}
// @Router /v1/Investment/GetInvestments [post]
func (InvestmentRouters) GetInvestmentsREPost(c *gin.Context) {
	var investmentData = application.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.InvestmentChartForAccount())
}

//GetInvestmentsTableREPost
// @Tags Investment
// @Summary 获取Investments的表数据
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} []investmentsModels.InvestmentTable {}
// @Router /v1/Investment/GetInvestmentsTable [post]
func (InvestmentRouters) GetInvestmentsTableREPost(c *gin.Context) {
	var investmentData = application.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.GetInvestmentTable())
}

//AddInvestmentsTableREPost
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
func (InvestmentRouters) AddInvestmentsTableREPost(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	Helheim.Writer(Helheim.Info, common.ToJSON(investmentData.Investment))

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	flag, err := application.AddInvestmentTable(&investmentData)

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, ReturnData{Successful: flag})
}

//UpdateInvestmentsTableREPost
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
func (InvestmentRouters) UpdateInvestmentsTableREPost(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	Helheim.Writer(Helheim.Info, common.ToJSON(investmentData.Investment))

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusBadRequest, ReturnData{Message: "参数错误", Error: err.Error(), Successful: false})
		return
	}

	flag, err := application.UpdateInvestmentTable(&investmentData)

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, ReturnData{Successful: flag})
}

//GetInvestmentDiagramREPost
// @Tags Investment
// @Summary 获取图表信息
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} map[string][]investmentsModels.Investment
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/GetInvestmentDiagram [post]
func (InvestmentRouters) GetInvestmentDiagramREPost(c *gin.Context) {

	data, err := application.GetInvestmentDiagram()

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{
			Successful: false,
			Error:      err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

//GetInvestmentOptionREPost
// @Tags Investment
// @Summary 获取查询条件信息(table page)
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} application.InvestmentOption
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/GetInvestmentDiagram [post]
func (InvestmentRouters) GetInvestmentOptionREPost(c *gin.Context) {
	data, err := application.GetInvestmentOption()

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, data)
}

//GetInvestmentServiceChargeREPost
// @Tags Investment
// @Summary 获取手续费(table page)
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":[]investmentsModels.InvestmentServiceCharge,"Error":"", Message:""}
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/GetInvestmentDiagram [post]
func (InvestmentRouters) GetInvestmentServiceChargeREPost(c *gin.Context) {
	item := struct {
		ItemID int
	}{}
	err := c.ShouldBindJSON(&item)
	data := application.GetInvestmentServiceCharge(item.ItemID)

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		c.JSON(http.StatusInternalServerError, ReturnData{Error: err.Error(), Successful: false})
		return
	}

	c.JSON(http.StatusOK, ReturnData{Error: "", Successful: true, Data: data})
}

//GetInvestmentReportFormREPost
// @Tags Investment
// @Summary 投资报表
// @Description 描述信息
// @Security Bearer
// @Produce  json
// @Success 200 {object} ReturnData {"Successful":true,"data":[]investmentsModels.InvestmentReportForm,"Error":"", Message:""}
// @Failure 500 {object} ReturnData {"Successful":true,"data":null,"Error":"", Message:""}
// @Router /v1/Investment/GetInvestmentDiagram [post]
func (InvestmentRouters) GetInvestmentReportFormREPost(c *gin.Context) {

	c.JSON(http.StatusOK, ReturnData{Error: "", Successful: true, Data: application.GetInvestmentReportForm()})
}
