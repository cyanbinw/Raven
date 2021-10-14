package investmentsModels

import "time"

type Investment struct {
	ID             int64     `db:"ID" json:"id"`
	ItemID         int       `db:"ItemID" json:"itemID"`
	Code           string    `db:"Code" json:"code"`
	Name           string    `db:"Name" json:"name"`
	TypeID         int       `db:"TypeID" json:"type"`
	Account        float64   `db:"Account" json:"account"`
	Share          float64   `db:"Share" json:"share"`
	NetWorth       float64   `db:"NetWorth" json:"netWorth"`
	ServiceCharge  float64   `db:"ServiceCharge" json:"serviceCharge"`
	Date           time.Time `db:"Date" json:"date"`
	ActivityStatus int       `db:"ActivityStatus" json:"activity"`
	IsEmpty        bool      `db:"IsEmpty" json:"isEmpty"`
}

type InvestmentTable struct {
	Investment   `xorm:"extends"`
	ActivityName string `json:"activityName"`
	TypeName     string `json:"typeName"`
}

type InvestmentActivity struct {
	ActivityID   int
	ActivityName string
	InsertDate   time.Time
}

type InvestmentItem struct {
	ItemID int
	Name   string
	Code   string
}

type InvestmentType struct {
	TypeID     int `xorm:"pk autoincr"`
	TypeName   string
	GroupID    int
	GroupName  string
	InsertDate time.Time `xorm:"created"`
}

type InvestmentChartModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (InvestmentTable) TableName() string {
	return "Investment"
}
