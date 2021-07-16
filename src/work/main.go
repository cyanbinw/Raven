package main

import (
	"Raven/src/work/billNameWork"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	billNameWork.SetBillName()
}
