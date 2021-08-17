package main

import (
	"github.com/WFallenDown/Raven/work/billNameWork/billNameService"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	billNameService.SetBillName()
}
