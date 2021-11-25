package application

import (
	"fmt"
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
	Type          []investmentsModels.InvestmentType              `json:"type"`
	Activity      []investmentsModels.InvestmentActivity          `json:"activity"`
	Item          []investmentsModels.InvestmentItem              `json:"item"`
	ServiceCharge []investmentsModels.InvestmentServiceChargeType `json:"serviceCharge"`
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
	InvestmentsInitDB()
	return database.InvestmentAddTable(*data)
}

func UpdateInvestmentTable(data *investmentsModels.InvestmentTable) (bool, error) {
	InvestmentsInitDB()
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
	option.Type, option.Activity, option.Item, option.ServiceCharge, err = database.InvestmentGetOption()
	return option, err
}

func GetInvestmentServiceCharge(itemID int) []investmentsModels.InvestmentServiceCharge {
	InvestmentsInitDB()
	return database.GetServiceChargeData(itemID)
}

func GetInvestmentReprotForm() investmentsModels.InvestmentReportForm {
	InvestmentsInitDB()
	data := database.InvestmentGetTable()
	setInvestmentReprotForm(data)
	return investmentsModels.InvestmentReportForm{}
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

func setInvestmentReprotForm(data []investmentsModels.InvestmentTable) {
	var itemList []investmentsModels.InvestmentReportForm
	item := new(investmentsModels.InvestmentReportForm)
	From(data).GroupBy(func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable).ItemID
	}, func(i interface{}) interface{} {
		return i.(investmentsModels.InvestmentTable)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		value := investmentsModels.InvestmentReportForm{}

		for _, item := range i.Group {
			if item.(investmentsModels.InvestmentTable).IsEmpty != true {
				if item.(investmentsModels.InvestmentTable).ActivityStatus == 1 {
					value.PositionInvement += item.(investmentsModels.InvestmentTable).Account
					value.PositionServiceCharge += item.(investmentsModels.InvestmentTable).ServiceCharge
				} else if item.(investmentsModels.InvestmentTable).ActivityStatus == 2 {
					value.PositionSell += item.(investmentsModels.InvestmentTable).Account
					value.PositionServiceCharge += item.(investmentsModels.InvestmentTable).ServiceCharge
				}
			} else {
				if item.(investmentsModels.InvestmentTable).ActivityStatus == 1 {
					value.ClearanceInvement += item.(investmentsModels.InvestmentTable).Account
					value.ClearanceServiceCharge += item.(investmentsModels.InvestmentTable).ServiceCharge
				} else if item.(investmentsModels.InvestmentTable).ActivityStatus == 2 {
					value.ClearanceSell += item.(investmentsModels.InvestmentTable).Account
					value.ClearanceServiceCharge += item.(investmentsModels.InvestmentTable).ServiceCharge
				}
			}
			if item.(investmentsModels.InvestmentTable).ActivityStatus == 1 {
				value.TotalPositionInvement += item.(investmentsModels.InvestmentTable).Account
				value.TotalServiceCharge += item.(investmentsModels.InvestmentTable).ServiceCharge
			} else if item.(investmentsModels.InvestmentTable).ActivityStatus == 2 {
				value.TotalClearanceInvement += item.(investmentsModels.InvestmentTable).Account
				value.TotalServiceCharge += item.(investmentsModels.InvestmentTable).ServiceCharge
			}
		}
		value.PositionInformation = i.Group[0].(investmentsModels.InvestmentTable).Name
		value.ClearanceInformation = i.Group[0].(investmentsModels.InvestmentTable).Name
		value.TotalInformation = i.Group[0].(investmentsModels.InvestmentTable).Name
		return value
	}).ToSlice(&itemList)
	for _, i := range itemList {
		if i.PositionInvement > 0 {
			item.PositionInvement += i.PositionInvement
			item.PositionServiceCharge += i.PositionServiceCharge
			item.PositionSell += i.PositionSell
			item.Position += 1
			if i.PositionInformation != "" {
				item.PositionInformation += i.PositionInformation + ";"
			}
			item.TotalPositionInvement += item.PositionInvement - item.PositionSell
		} else if i.ClearanceInvement > 0 {
			item.ClearanceInvement += i.ClearanceInvement
			item.ClearanceServiceCharge += i.ClearanceServiceCharge
			item.ClearanceSell += i.ClearanceSell
			item.Clearance += 1
			if i.ClearanceInformation != "" {
				item.ClearanceInformation += i.ClearanceInformation + ";"
			}
			item.TotalClearanceInvement += item.ClearanceInvement - item.ClearanceSell
		}

		item.TotalInvement += i.TotalInvement
		item.TotalServiceCharge += i.TotalServiceCharge
		item.TotalSell += i.TotalSell
		item.Total += 1
		if i.TotalInformation != "" {
			item.TotalInformation += i.TotalInformation + ";"
		}
		item.TotalTotalInvement += item.TotalPositionInvement + item.TotalClearanceInvement
	}
	fmt.Println(item)
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
