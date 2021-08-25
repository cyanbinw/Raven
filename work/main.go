package main

import (
	"github.com/WFallenDown/Raven/work/billNameWork"
	"github.com/WFallenDown/Raven/work/userWork"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func main() {
	billNameWork.SetBillName()
	userWork.SetUser()
}
