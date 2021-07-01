package database

import (
	"database/sql"
	"github.com/go-xorm/xorm"
)

const (
	userName = ""
	password = ""
	ip       = ""
	port     = ""
	dbName   = ""
	month    = 12
)

var db *sql.DB
var engine *xorm.Engine

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
