package controllers

import (
	"Raven/src/log"
	"Raven/src/web/models/investmentsModels"
	"Raven/src/web/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInvestments(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.InvestmentChartForAccount())
}

func GetInvestmentsTable(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.GetInvestmentTable())
}

func AddInvestmentsTable(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	log.Writer(log.Info, service.ToJSON(investmentData.Investment))

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数错误", "error": err})
		return
	}

	investmentData.InvestmentsInitDB()
	flag, err := investmentData.AddInvestmentTable()

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": flag})
}

func UpdateInvestmentsTable(c *gin.Context) {
	var investmentData = investmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	log.Writer(log.Info, service.ToJSON(investmentData.Investment))

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数错误", "error": err})
		return
	}

	investmentData.InvestmentsInitDB()
	flag, err := investmentData.UpdateInvestmentTable()

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": flag})
}

func GetInvestmentDiagram(c *gin.Context) {

	data, err := investmentsModels.GetInvestmentDiagram()

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetInvestmentOption(c *gin.Context) {
	Type, Activity, Item, err := investmentsModels.GetInvestmentOption()

	if err != nil {
		log.Writer(log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"type": Type, "activity": Activity, "item": Item})
}
