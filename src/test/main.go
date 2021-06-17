package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//investmentData, err := investmentsModels.GetInvestmentDiagram()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(investmentData)

	bitXOR()
}

func bitXOR() {
	num1 := 1
	num2 := 2

	fmt.Println(num1)

	num1 = num1 ^ num2

	fmt.Println(num1)

	num2 = num1 ^ num2

	fmt.Println(num2)

	num1 = num1 ^ num2

	fmt.Println(num1)
}
