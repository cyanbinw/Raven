package main

import (
	"Raven/src/work/billNameWork/billNameService"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	billNameService.SetBillName()
}
