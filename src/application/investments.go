package application

import (
	"Raven/src/database"
	"Raven/src/models/investmentsModels"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/shopspring/decimal"
)

type InvestmentData struct {
	Data []investmentsModels.Investment
}

type InvestmentAccountModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type InvestmentGroup struct {
	Data  []investmentsModels.Investment
	Count int
	Name  string
}

type InvestmentsChartModel struct {
	Account  []investmentsModels.InvestmentChartModel
	Share    []investmentsModels.InvestmentChartModel
	NetWorth []investmentsModels.InvestmentChartModel
}

type InvestmentOption struct {
	Type     []investmentsModels.InvestmentType     `json:"type"`
	Activity []investmentsModels.InvestmentActivity `json:"activity"`
	Item     []investmentsModels.InvestmentItem     `json:"item"`
}

type InvestmentGroupList []InvestmentGroup

func (data *InvestmentData) InvestmentsInitDB() {
	database.InvestmentsInitDB()
}

func (data *InvestmentData) InvestmentGetAll() {
	data.Data = database.InvestmentGetAll()
}

func (data *InvestmentData) InvestmentChartForAccount() InvestmentsChartModel {
	item := database.InvestmentGetChart()
	return createChart(item)

	//return investmentGetDataToChart()
}

func (data *InvestmentData) GetInvestmentTable() []investmentsModels.InvestmentTable {
	return database.InvestmentGetTable()
}

func InvestmentsInitDB() {
	database.InvestmentsInitDB()
}

func AddInvestmentTable(data *investmentsModels.InvestmentTable) (bool, error) {
	return database.InvestmentAddTable(*data)
}

func UpdateInvestmentTable(data *investmentsModels.InvestmentTable) (bool, error) {
	return database.InvestmentUpdateTable(*data)
}

func GetInvestmentDiagram() (map[string][]investmentsModels.Investment, error) {
	database.InvestmentsInitDB()
	return investmentGetDiagram()
}

func GetInvestmentOption() (*InvestmentOption, error) {
	var option = new(InvestmentOption)
	var err error
	database.InvestmentsInitDB()
	option.Type, option.Activity, option.Item, err = database.InvestmentGetOption()
	return option, err
}

func createChart(data []investmentsModels.Investment) InvestmentsChartModel {
	var item InvestmentsChartModel

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.Investment).Name
	}, func(i interface{}) interface{} {
		return i.(investmentsModels.Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(investmentsModels.Investment).Account
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return investmentsModels.InvestmentChartModel{Name: i.Key.(string), Value: m}
	}).ToSlice(&item.Account)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.Investment).Name
	}, func(i interface{}) interface{} {
		return i.(investmentsModels.Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(investmentsModels.Investment).NetWorth
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return investmentsModels.InvestmentChartModel{Name: i.Key.(string), Value: m}
	}).ToSlice(&item.NetWorth)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.Investment).Name
	}, func(i interface{}) interface{} {
		return i.(investmentsModels.Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(investmentsModels.Investment).Share
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return investmentsModels.InvestmentChartModel{Name: i.Key.(string), Value: m}
	}).ToSlice(&item.Share)

	return item
}

func investmentGetDiagram() (map[string][]investmentsModels.Investment, error) {
	data := make(map[string][]investmentsModels.Investment)

	investments := database.InvestmentGetDateOrderbyDate()
	for _, index := range *investments {
		data[index.Name] = append(data[index.Name], index)
	}
	return data, nil
}

func shareOutBonus() {

}

func (data InvestmentGroupList) Len() int {
	return len(data)
}

func (data InvestmentGroupList) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func (data InvestmentGroupList) Less(i, j int) bool {
	return data[i].Count > data[j].Count
}
