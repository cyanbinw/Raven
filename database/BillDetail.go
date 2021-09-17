package database

import (
	"time"
)

type Billdetail struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	Billnumber string    `xorm:"VARCHAR(20)"`
	Type       string    `xorm:"VARCHAR(100)"`
	Billname   string    `xorm:"VARCHAR(50)"`
	Account    string    `xorm:"DECIMAL(10,2)"`
	Date       time.Time `xorm:"DATETIME"`
	Remarks    string    `xorm:"VARCHAR(500)"`
}
