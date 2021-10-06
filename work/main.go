package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/swirling-melodies/Raven/work/billNameWork"
	"github.com/swirling-melodies/Raven/work/investmentWork"
	"github.com/swirling-melodies/Raven/work/userWork"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func main() {
	billNameWork.SetBillName()
	userWork.SetUser()
	investmentWork.SetInvestmentItem()
}
