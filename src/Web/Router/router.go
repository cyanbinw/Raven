package Router

import (
	"Raven/src/Web/Controllers"
	"Raven/src/Web/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.CrossDomain())

	test := router.Group("test")
	testGroup(test)

	v1 := router.Group("v1")
	billGroup(v1)
	desireGroup(v1)
	investmentGroup(v1)

	router.Run()
}

func testGroup(c *gin.RouterGroup) {
	c.GET("/testInsert", Controllers.TestHome)
}

func billGroup(c *gin.RouterGroup) {
	c.POST("/GetYearAllData", Controllers.GetYearAllData)
	c.POST("/GetFourMonthAllData", Controllers.GetFourMonthData)
}

func investmentGroup(c *gin.RouterGroup) {
	c.POST("/GetInvestments", Controllers.GetInvestments)
	c.POST("/GetInvestmentsTable", Controllers.GetInvestmentsTable)
	c.POST("/AddInvestmentsTable", Controllers.AddInvestmentsTable)
	c.POST("/UpdateInvestmentsTable", Controllers.UpdateInvestmentsTable)
	c.POST("/GetInvestmentDiagram", Controllers.GetInvestmentDiagram)
	c.POST("/GetInvestmentOption", Controllers.GetInvestmentOption)
}

func desireGroup(c *gin.RouterGroup) {
	c.POST("/GetDesire", Controllers.GetDesire)
}
