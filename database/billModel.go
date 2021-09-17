package database

import (
	"time"
)

type BillDetail struct {
	ID         int       `xorm:"not null pk autoincr INT"`
	BillNumber string    `xorm:"VARCHAR(20)"`
	Type       string    `xorm:"VARCHAR(100)"`
	BillName   string    `xorm:"VARCHAR(50)"`
	Account    string    `xorm:"DECIMAL(10,2)"`
	Date       time.Time `xorm:"DATETIME"`
	Remarks    string    `xorm:"VARCHAR(500)"`
}

type BillNameConfig struct {
	ID         int       `xorm:"not null pk autoincr INT"`
	BillName   string    `xorm:"VARCHAR(255)"`
	Count      int       `xorm:"INT"`
	Color      string    `xorm:"VARCHAR(255)"`
	Icon       string    `xorm:"VARCHAR(255)"`
	UpdateDate time.Time `xorm:"DATETIME"`
	CreatDate  time.Time `xorm:"DATETIME"`
}

type BillNameConfigAudit struct {
	ID         int       `xorm:"not null pk autoincr INT"`
	BillID     int       `xorm:"INT"`
	Billname   string    `xorm:"VARCHAR(255)"`
	Count      int       `xorm:"INT"`
	Color      string    `xorm:"VARCHAR(255)"`
	Icon       string    `xorm:"VARCHAR(255)"`
	UpdateDate time.Time `xorm:"DATETIME"`
	CreatDate  time.Time `xorm:"DATETIME"`
	Status     int       `xorm:"INT"`
}
