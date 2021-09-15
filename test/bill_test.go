package test

import (
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
	if data.Total == 0 {
		t.FailNow()
	}
}
