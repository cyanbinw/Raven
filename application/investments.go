package application

import (
	. "github.com/ahmetb/go-linq/v3"
	"github.com/shopspring/decimal"
	"github.com/swirling-melodies/Raven/database"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
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
	Account    []investmentsModels.InvestmentChartModel
	Share      []investmentsModels.InvestmentChartModel
	NetWorth   []investmentsModels.InvestmentChartModel
	Proportion []investmentsModels.InvestmentChartModel
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

func createChart(data []investmentsModels.InvestmentTable) InvestmentsChartModel {
	var item InvestmentsChartModel

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable).Name
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
				expenditure += item.(investmentsModels.InvestmentTable).Account
			} else if item.(investmentsModels.InvestmentTable).ActivityStatus == 2 {
				revenue += item.(investmentsModels.InvestmentTable).Account
			}
		}

		value, _ = decimal.NewFromFloat(expenditure - revenue).Round(4).Float64()

		return investmentsModels.InvestmentChartModel{Name: i.Key.(string), Value: value}
	}).ToSlice(&item.Account)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable).Name
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
				expenditure += item.(investmentsModels.InvestmentTable).NetWorth
			} else if item.(investmentsModels.InvestmentTable).ActivityStatus == 2 {
				revenue += item.(investmentsModels.InvestmentTable).NetWorth
			}
		}

		value, _ = decimal.NewFromFloat(expenditure - revenue).Round(4).Float64()

		return investmentsModels.InvestmentChartModel{Name: i.Key.(string), Value: value}
	}).ToSlice(&item.NetWorth)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable).Name
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

		return investmentsModels.InvestmentChartModel{Name: i.Key.(string), Value: value}
	}).ToSlice(&item.Share)

	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable).TypeName
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
				expenditure += item.(investmentsModels.InvestmentTable).Account
			} else if item.(investmentsModels.InvestmentTable).ActivityStatus == 2 {
				revenue += item.(investmentsModels.InvestmentTable).Account
			}
		}

		value, _ = decimal.NewFromFloat(expenditure - revenue).Round(4).Float64()

		return investmentsModels.InvestmentChartModel{Name: i.Key.(string), Value: value}
	}).ToSlice(&item.Proportion)

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
