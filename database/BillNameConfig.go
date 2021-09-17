package database

import (
	"time"
)

type Billnameconfig struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Billname   string    `xorm:"VARCHAR(255)"`
	Count      int       `xorm:"INT"`
	Color      string    `xorm:"VARCHAR(255)"`
	Icon       string    `xorm:"VARCHAR(255)"`
	Updatedate time.Time `xorm:"DATETIME"`
	Creatdate  time.Time `xorm:"DATETIME"`
}
