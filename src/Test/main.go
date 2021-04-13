package main

import (
	"Raven/src/Web/Models/InvestmentsModels"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	investmentData, err := InvestmentsModels.GetInvestmentDiagram()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(investmentData)
}
