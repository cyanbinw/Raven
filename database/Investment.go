package database

import (
	"time"
)

type Investment struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	Itemid         int       `xorm:"not null INT"`
	Code           string    `xorm:"VARCHAR(50)"`
	Name           string    `xorm:"VARCHAR(100)"`
	Typeid         int       `xorm:"INT"`
	Account        string    `xorm:"DECIMAL(10,4)"`
	Share          string    `xorm:"DECIMAL(10,4)"`
	Networth       string    `xorm:"DECIMAL(10,4)"`
	Date           time.Time `xorm:"DATETIME"`
	Activitystatus int       `xorm:"INT"`
	Isempty        int       `xorm:"not null TINYINT(1)"`
}
