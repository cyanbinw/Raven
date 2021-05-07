package Controllers

import (
	"Raven/src/Log"
	"Raven/src/Web/Models/InvestmentsModels"
	"Raven/src/Web/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInvestments(c *gin.Context) {
	var investmentData = InvestmentsModels.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.SetInvestmentChartForAccount())
}

func GetInvestmentsTable(c *gin.Context) {
	var investmentData = InvestmentsModels.InvestmentData{}

	investmentData.InvestmentsInitDB()

	c.JSON(http.StatusOK, investmentData.GetInvestmentTable())
}

func AddInvestmentsTable(c *gin.Context) {
	var investmentData = InvestmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	Log.Writer(Log.Info, Service.ToJSON(investmentData.Investment))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数错误", "error": err})
		return
	}

	investmentData.InvestmentsInitDB()
	flag, err := investmentData.AddInvestmentTable()

	if err != nil {
		Log.Writer(Log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": flag})
}

func UpdateInvestmentsTable(c *gin.Context) {
	var investmentData = InvestmentsModels.InvestmentTable{}

	err := c.ShouldBindJSON(&investmentData)
	Log.Writer(Log.Info, Service.ToJSON(investmentData.Investment))

	if err != nil {
		Log.Writer(Log.Error, err)
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数错误", "error": err})
		return
	}

	investmentData.InvestmentsInitDB()
	flag, err := investmentData.UpdateInvestmentTable()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": flag})
}

func GetInvestmentDiagram(c *gin.Context) {

	data, err := InvestmentsModels.GetInvestmentDiagram()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetInvestmentOption(c *gin.Context) {
	Type, Activity, err := InvestmentsModels.GetInvestmentOption()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"type": Type, "activity": Activity})
}
