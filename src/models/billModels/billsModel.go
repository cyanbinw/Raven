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
