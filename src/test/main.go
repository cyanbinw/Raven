package main

import (
	"Raven/src/web/models/investmentsModels"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	investmentData, err := investmentsModels.GetInvestmentDiagram()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(investmentData)
}
