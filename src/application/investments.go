package application

import (
	database2 "Raven/src/database"
	investmentsModels2 "Raven/src/models/investmentsModels"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/shopspring/decimal"
)

type InvestmentData struct {
	Data []investmentsModels2.Investment
}

type InvestmentAccountModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type InvestmentGroup struct {
	Data  []investmentsModels2.Investment
	Count int
	Name  string
}

type InvestmentsChartModel struct {
	Account  []investmentsModels2.InvestmentChartModel
	Share    []investmentsModels2.InvestmentChartModel
	NetWorth []investmentsModels2.InvestmentChartModel
}

type InvestmentOption struct {
	Type     []investmentsModels2.InvestmentType     `json:"type"`
	Activity []investmentsModels2.InvestmentActivity `json:"activity"`
	Itme     []investmentsModels2.InvestmentItem     `json:"item"`
}

type InvestmentGroupList []InvestmentGroup

func (data *InvestmentData) InvestmentsInitDB() {
	database2.InvestmentsInitDB()
}

func (data *InvestmentData) InvestmentGetAll() {
	data.Data = database2.InvestmentGetAll()
}

func (data *InvestmentData) InvestmentChartForAccount() InvestmentsChartModel {
	item := database2.InvestmentGetChart()
	return createChart(item)

	//return investmentGetDataToChart()
}

func (data *InvestmentData) GetInvestmentTable() []investmentsModels2.InvestmentTable {
	return database2.InvestmentGetTable()
}

func InvestmentsInitDB() {
	database2.InvestmentsInitDB()
}

func AddInvestmentTable(data *investmentsModels2.InvestmentTable) (bool, error) {
	return database2.InvestmentAddTable(*data)
}

func UpdateInvestmentTable(data *investmentsModels2.InvestmentTable) (bool, error) {
	return database2.InvestmentUpdateTable(*data)
}

func GetInvestmentDiagram() (map[string][]investmentsModels2.Investment, error) {
	database2.InvestmentsInitDB()
	return investmentGetDiagram()
}

func GetInvestmentOption() (*InvestmentOption, error) {
	var option = new(InvestmentOption)
	var err error
	database2.InvestmentsInitDB()
	option.Type, option.Activity, option.Itme, err = database2.InvestmentGetOption()
	return option, err
}

func createChart(data []investmentsModels2.Investment) InvestmentsChartModel {
	var item InvestmentsChartModel

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels2.Investment).Name
	}, func(i interface{}) interface{} {
		return i.(investmentsModels2.Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(investmentsModels2.Investment).Account
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return investmentsModels2.InvestmentChartModel{Name: i.Key.(string), Value: m}
	}).ToSlice(&item.Account)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels2.Investment).Name
	}, func(i interface{}) interface{} {
		return i.(investmentsModels2.Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(investmentsModels2.Investment).NetWorth
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return investmentsModels2.InvestmentChartModel{Name: i.Key.(string), Value: m}
	}).ToSlice(&item.NetWorth)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels2.Investment).Name
	}, func(i interface{}) interface{} {
		return i.(investmentsModels2.Investment)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(investmentsModels2.Investment).Share
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return investmentsModels2.InvestmentChartModel{Name: i.Key.(string), Value: m}
	}).ToSlice(&item.Share)

	return item
}

func investmentGetDiagram() (map[string][]investmentsModels2.Investment, error) {
	data := make(map[string][]investmentsModels2.Investment)

	investments := database2.InvestmentGetDateOrderbyDate()
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
