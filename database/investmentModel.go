package database

import "time"

type Investment struct {
	ID             int64     `xorm:"not null pk autoincr INT"`
	ItemID         int       `xorm:"not null INT"`
	Code           string    `xorm:"VARCHAR(50)"`
	Name           string    `xorm:"VARCHAR(100)"`
	TypeID         int       `xorm:"INT"`
	Account        float64   `xorm:"DECIMAL(10,4)"`
	Share          float64   `xorm:"DECIMAL(10,4)"`
	NetWorth       float64   `xorm:"DECIMAL(10,4)"`
	Date           time.Time `xorm:"DATETIME"`
	ActivityStatus int       `xorm:"INT"`
	IsEmpty        bool      `xorm:"not null TINYINT(1)"`
}

type InvestmentTable struct {
	Investment   `xorm:"extends"`
	ActivityName string `json:"activityName"`
	TypeName     string `json:"typeName"`
}

func (InvestmentTable) TableName() string {
	return "Investment"
}

type InvestmentActivity struct {
	ActivityID   int       `xorm:"not null pk autoincr INT"`
	ActivityName string    `xorm:"not null default '' VARCHAR(100)"`
	InsertDate   time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}

type InvestmentItem struct {
	ItemID int    `xorm:"INT"`
	Name   string `xorm:"VARCHAR(255)"`
	Code   string `xorm:"VARCHAR(255)"`
}

type InvestmentType struct {
	TypeID     int       `xorm:"not null pk autoincr INT"`
	TypeName   string    `xorm:"not null default '' VARCHAR(100)"`
	InsertDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
