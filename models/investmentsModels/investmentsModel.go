package investmentsModels

import (
	"github.com/swirling-melodies/Raven/database"
	"time"
)

type Investment struct {
	ID             int64     `db:"ID" json:"id"`
	ItemID         int       `db:"ItemID" json:"itemID"`
	Code           string    `db:"Code" json:"code"`
	Name           string    `db:"Name" json:"name"`
	TypeID         int       `db:"TypeID" json:"type"`
	Account        float64   `db:"Account" json:"account"`
	Share          float64   `db:"Share" json:"share"`
	NetWorth       float64   `db:"NetWorth" json:"netWorth"`
	Date           time.Time `db:"Date" json:"date"`
	ActivityStatus int       `db:"ActivityStatus" json:"activity"`
	IsEmpty        bool      `db:"IsEmpty" json:"isEmpty"`
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
	TypeID     int
	TypeName   string
	InsertDate time.Time
}

type InvestmentChartModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (data *Investment) Set(i database.Investment) {
	data.ID = i.ID
	data.ItemID = i.ItemID
	data.Code = i.Code
	data.Name = i.Name
	data.TypeID = i.TypeID
	data.Account = i.Account
	data.Share = i.Share
	data.NetWorth = i.NetWorth
	data.Date = i.Date
	data.ActivityStatus = i.ActivityStatus
	data.IsEmpty = i.IsEmpty
}

func (data *Investment) Get(i database.Investment) *database.Investment {
	i.ID = data.ID
	i.ItemID = data.ItemID
	i.Code = data.Code
	i.Name = data.Name
	i.TypeID = data.TypeID
	i.Account = data.Account
	i.Share = data.Share
	i.NetWorth = data.NetWorth
	i.Date = data.Date
	i.ActivityStatus = data.ActivityStatus
	i.IsEmpty = data.IsEmpty

	return &i
}

func SetInvestmentList(i []database.Investment) []Investment {
	var list []Investment
	data := new(Investment)
	for _, item := range i {
		data.ID = item.ID
		data.ItemID = item.ItemID
		data.Code = item.Code
		data.Name = item.Name
		data.TypeID = item.TypeID
		data.Account = item.Account
		data.Share = item.Share
		data.NetWorth = item.NetWorth
		data.Date = item.Date
		data.ActivityStatus = item.ActivityStatus
		data.IsEmpty = item.IsEmpty
		list = append(list, *data)
	}
	return list
}

func SetInvestmentTypeList(i []database.InvestmentType) {
	var list []InvestmentType
	data := new(InvestmentType)
	for _, item := range i {
		data.TypeID = item.TypeID
		data.TypeName = item.TypeName
		data.InsertDate = item.InsertDate
		list = append(list, *data)
	}
}
