package investmentsModels

import (
	"Raven/src/log"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/shopspring/decimal"
)

type InvestmentData struct {
	Data []Investment
}

type InvestmentAccountModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type InvestmentChartModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type InvestmentsChartModel struct {
	Account  []InvestmentChartModel
	Share    []InvestmentChartModel
	NetWorth []InvestmentChartModel
}

func (data *InvestmentData) InvestmentsInitDB() {
	investmentsInitDB()
}

func (data *InvestmentData) InvestmentGetAll() {
	investmentGetAll(data)
}

func (data *InvestmentData) InvestmentChartForAccount() InvestmentsChartModel {
	item := investmentGetChart()
	return createChart(item)

	//return investmentGetDataToChart()
}

func (data *InvestmentData) GetInvestmentTable() []InvestmentTable {
	return investmentGetTable()
}

func (data *InvestmentTable) InvestmentsInitDB() {
	investmentsInitDB()
}

func (data *InvestmentTable) AddInvestmentTable() (bool, error) {
	return investmentAddTable(*data)
}

func (data *InvestmentTable) UpdateInvestmentTable() (bool, error) {
	return investmentUpdateTable(*data)
}

func GetInvestmentDiagram() (map[string][]Investment, error) {
	investmentsInitDB()
	return investmentGetDiagram()
}

func GetInvestmentOption() ([]InvestmentType, []InvestmentActivity, []InvestmentItem, error) {
	investmentsInitDB()
	return investmentGetOption()
}

func createChart(data []Investment) InvestmentsChartModel {
	var item InvestmentsChartModel

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(Investment).Name
	}, func(i interface{}) interface{} {
		return i.(Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(Investment).Account
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return InvestmentChartModel{i.Key.(string), m}
	}).ToSlice(&item.Account)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(Investment).Name
	}, func(i interface{}) interface{} {
		return i.(Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(Investment).NetWorth
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return InvestmentChartModel{i.Key.(string), m}
	}).ToSlice(&item.NetWorth)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(Investment).Name
	}, func(i interface{}) interface{} {
		return i.(Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(Investment).Share
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return InvestmentChartModel{i.Key.(string), m}
	}).ToSlice(&item.Share)

	return item
}

func investmentGetDiagram() (map[string][]Investment, error) {
	data := make(map[string][]Investment)

	var investments []Investment

	err := engine.OrderBy("Date").Find(&investments)
	if err != nil {
		log.Writer(log.Error, err)
	}
	for _, index := range investments {
		data[index.Name] = append(data[index.Name], index)
	}
	return data, nil
}

func shareOutBonus() {

}
