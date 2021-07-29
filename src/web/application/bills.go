package application

import (
	"encoding/json"
	"github.com/WFallenDown/Helheim"
	"github.com/WFallenDown/Raven/src/web/database"
	"github.com/WFallenDown/Raven/src/web/models/billModels"
	"github.com/WFallenDown/Raven/src/web/service"
	. "github.com/ahmetb/go-linq/v3"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"os"
	"strconv"
)

type BillOption struct {
	BillName []string
	BillType []string
}

type BillChartModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type BillDataByDate struct {
	Data []billModels.BillDetail
	Year int `json:"Year" form:"Year"`
}

type BillCharts []BillChartModel
type BillChartsData struct {
	BillCharts `json:"billCharts"`
	Total      float64 `json:"total"`
}

type IBillData interface {
	NewBillData()
	BillsInitDB()
	BillsWriteToJSON()
	BillsGetYearData()
}

func (data *BillDataByDate) NewBillData() {

}

func (data *BillDataByDate) BillsInitDB() {
	database.BillsInitDB()
}

func (data *BillDataByDate) BillsWriteToJSON() {
	var f *os.File
	src := strconv.Itoa(data.Year) + ".json"
	val, err := json.MarshalIndent(data.Data, "", "	") // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	if service.CheckFileIsExist(src) { //如果文件存在
		f, err = os.OpenFile(src, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(src) //创建文件

	}

	err = ioutil.WriteFile(src, val, 0777)
	service.CheckErr(err)
	f.Close()
}

func (data *BillDataByDate) BillsGetYearData() {
	database.BillsGetYearData(&data.Data, data.Year)
}

func (data *BillDataByDate) BillsGetDataByMonth() {
	database.BillsGetDataByMonth(&data.Data)
}

func BillsGetTable(bill *billModels.BillTable) *billModels.BillTable {
	database.BillsInitDB()
	database.BillsGetTable(bill)
	return bill
}

func BillsGetTableOption() *BillOption {
	var option = new(BillOption)
	database.BillsInitDB()
	option.BillName, option.BillType = database.BillsGetTableOption()
	return option
}

func BillsGetDiagram(bill *billModels.BillTable) (*BillChartsData, error) {
	var data = new(BillChartsData)

	database.BillsInitDB()
	database.BillsGetDiagram(bill)

	From(bill.BillDetail).GroupBy(func(i interface{}) interface{} {
		return i.(billModels.BillDetail).BillName
	}, func(i interface{}) interface{} {
		return i.(billModels.BillDetail)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(billModels.BillDetail).Account
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return BillChartModel{i.Key.(string), m}
	}).ToSlice(&data.BillCharts)
	expend := From(bill.BillDetail).Where(func(i interface{}) bool {
		return i.(billModels.BillDetail).Type == "支出"
	}).Select(func(i interface{}) interface{} {
		return i.(billModels.BillDetail).Account
	}).SumFloats()
	income := From(bill.BillDetail).Where(func(i interface{}) bool {
		return i.(billModels.BillDetail).Type == "收入"
	}).Select(func(i interface{}) interface{} {
		return i.(billModels.BillDetail).Account
	}).SumFloats()
	data.Total, _ = decimal.NewFromFloat(expend - income).Round(4).Float64()
	return data, nil
}

func BillsGetDataByPage(bill *billModels.BillDataByPage) *billModels.BillDataByPage {
	database.BillsInitDB()
	database.BillsGetDataByPage(bill)
	return bill
}
