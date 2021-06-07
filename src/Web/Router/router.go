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
	billGroupV1(v1)
	desireGroupV1(v1)
	investmentGroupV1(v1)
	targetGroupV1(v1)

	router.Run()
}

func testGroup(c *gin.RouterGroup) {
	c.GET("/testInsert", Controllers.TestHome)
}

func billGroupV1(c *gin.RouterGroup) {
	c.POST("/Bill/GetYearAllData", Controllers.GetYearAllData)
	c.POST("/Bill/GetFourMonthAllData", Controllers.GetFourMonthData)
}

func investmentGroupV1(c *gin.RouterGroup) {
	c.POST("/Investment/GetInvestments", Controllers.GetInvestments)
	c.POST("/Investment/GetInvestmentsTable", Controllers.GetInvestmentsTable)
	c.POST("/Investment/AddInvestmentsTable", Controllers.AddInvestmentsTable)
	c.POST("/Investment/UpdateInvestmentsTable", Controllers.UpdateInvestmentsTable)
	c.POST("/Investment/GetInvestmentDiagram", Controllers.GetInvestmentDiagram)
	c.POST("/Investment/GetInvestmentOption", Controllers.GetInvestmentOption)
}

func desireGroupV1(c *gin.RouterGroup) {
	c.POST("/Desire/GetDesire", Controllers.GetDesire)
}

func targetGroupV1(c *gin.RouterGroup) {
	c.POST("/Target/GetTarget", Controllers.GetTarget)
}
