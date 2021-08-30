package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/swirling-melodies/Raven/work/billNameWork"
	"github.com/swirling-melodies/Raven/work/userWork"
)

var engine *xorm.Engine

func main() {
	billNameWork.SetBillName()
	userWork.SetUser()
}
