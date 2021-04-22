package DesireModels

import (
	"Raven/src/Web/Service"
	"github.com/go-xorm/xorm"
)

type Desire struct {
}

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
var engine *xorm.Engine

func desireInitDB() {
	engine = Service.InitDB()
}
