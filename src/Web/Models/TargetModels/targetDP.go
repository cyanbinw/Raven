package TargetModels

import (
	"Raven/src/Web/Service"
	"fmt"
	"github.com/go-xorm/xorm"
	"time"
)

type Target struct {
	FiveYearTarget string `xorm:"5YearTarget"`
	Revenue        float32
	Expend         float32
	Skill          string
	Other          string
	InstallDate    time.Time
	UpdateDate     time.Time
}

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
var engine *xorm.Engine

func targetInitDB() {
	engine = Service.InitDB()
}

func getTarget() []Target {
	if engine == nil {
		fmt.Println("so so")
	}
	return nil
}
