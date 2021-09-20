package test

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/swirling-melodies/Raven/application"
	"github.com/swirling-melodies/Raven/models/billModels"
	"testing"
)

func TestBillsGetDiagram(t *testing.T) {
	var bill = billModels.BillTable{}

	data, err := application.BillsGetDiagram(&bill)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if data.Total < 0 {
		t.Log(data)
		t.FailNow()
	}
	t.Log(data)
}

func TestBillDataByDate(t *testing.T) {
	bill := new(application.BillDataByDate)
	bill.BillsInitDB()
	bill.BillsGetDataByMonth()
	for _, i := range bill.Data {
		t.Log(i.Date)
	}
}
