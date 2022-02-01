package investmentWork

import (
	. "github.com/ahmetb/go-linq/v3"
	"github.com/shopspring/decimal"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/database"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"xorm.io/xorm"
)

type isEmptyModel struct {
	ItemID  int
	Share   float64
	IsEmpty bool
}

func actionIsEmpty(session *xorm.Session) error {

	setDataIsEmpty(session)

	return nil
}

func setDataIsEmpty(session *xorm.Session) error {

	database.InvestmentsInitDB()
	data := database.InvestmentGetTable()
	item := selectModifiedData(data)
	for _, i := range *item {
		if i.Share == 0 {
			_, err := session.Where("ItemID", i.ItemID).Update(i.IsEmpty)
			if err != nil {
				session.Rollback()
				Helheim.Writer(Helheim.Error, err)
			}
		}
	}

	return nil
}

func selectModifiedData(data []investmentsModels.InvestmentTable) *[]isEmptyModel {
	var item []isEmptyModel
	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable).ItemID
	}, func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		expenditure := 0.0
		revenue := 0.0
		value := 0.0
		for _, item := range i.Group {
			if item.(investmentsModels.InvestmentTable).ActivityStatus == 1 {
				expenditure += item.(investmentsModels.InvestmentTable).Share
			} else if item.(investmentsModels.InvestmentTable).ActivityStatus == 2 {
				revenue += item.(investmentsModels.InvestmentTable).Share
			}
		}

		value, _ = decimal.NewFromFloat(expenditure - revenue).Round(4).Float64()
		isEmpty := false
		if value == 0.0 {
			isEmpty = true
		}
		return isEmptyModel{ItemID: i.Key.(int), Share: value, IsEmpty: isEmpty}
	}).ToSlice(&item)

	return &item
}
