package database

import (
	"database/sql"
	"github.com/swirling-melodies/Raven/common"
	"xorm.io/xorm"
)

const (
	userName = ""
	password = ""
	ip       = ""
	port     = ""
	dbName   = ""
	month    = 12
	insert   = 1
	update   = 2
	delete   = 3
)

var db *sql.DB
var engine *xorm.Engine

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间

func InitDB() {
	engine = common.InitDB()
}
