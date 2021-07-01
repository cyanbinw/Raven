package investmentsModels

import "time"

type Investment struct {
	ID             int64   `db:"ID" json:"id"`
	ItemID         int     `db:"ItemID" json:"itemID"`
	Name           string  `db:"Name" json:"name"`
	TypeID         int     `db:"TypeID" json:"type"`
	Account        float64 `db:"Account" json:"account"`
	Share          float64 `db:"Share" json:"share"`
	NetWorth       float64 `db:"NetWorth" json:"netWorth"`
	Date           string  `db:"Date" json:"date"`
	ActivityStatus int     `db:"ActivityStatus" json:"activity"`
	IsEmpty        bool    `db:"IsEmpty" json:"isEmpty"`
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
}

type InvestmentType struct {
	TypeID     int
	TypeName   string
	InsertDate time.Time
}

type InvestmentChartModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (InvestmentTable) TableName() string {
	return "Investment"
}
