package billModels

import "time"

//BillDetail is test
type BillDetail struct {
	ID         int64   `db:"ID"`
	BillNumber string  `db:"BillNumber"`
	Type       string  `db:"Type"`
	BillName   string  `db:"BillName"`
	Account    float64 `db:"Account"`
	Date       string  `db:"Date"`
	Remarks    string  `db:"Remarks"`
}

type BillTable struct {
	BillDetail []BillDetail
	PageSize   int
	Total      int64
	PageNumber int
	BillType   []string
	BillName   []string
	AccountMax float64
	AccountMin float64
	DateMax    time.Time
	DateMin    time.Time
}

type BillName struct {
	ID         int `xorm:"pk autoincr"`
	BillName   string
	Count      int
	Color      string
	Icon       string
	UpdateDate time.Time `xorm:"updated"` //修改后自动更新时间
	CreatDate  time.Time `xorm:"created"` //创建时间
}

type BillNameAudit struct {
	ID         int `xorm:"pk autoincr"`
	BillID     int
	BillName   string
	Count      int
	Color      string
	Icon       string
	UpdateDate time.Time `xorm:"updated"` //修改后自动更新时间
	CreatDate  time.Time `xorm:"created"` //创建时间
	Status     int
}

type BillData struct {
	BillDetail `xorm:"extends"`
	ID         int
	Count      int
	Color      string
	Icon       string
}

type BillDataByPage struct {
	BillData   []BillData
	PageSize   int
	PageNumber int
}

func (BillData) TableName() string {
	return "BillDetail"
}
