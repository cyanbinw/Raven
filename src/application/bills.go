package application

import (
	database2 "Raven/src/database"
	"Raven/src/log"
	billModels2 "Raven/src/models/billModels"
	service2 "Raven/src/service"
	"encoding/json"
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
	Data []billModels2.BillDetail
	Year int `json:"Year" form:"Year"`
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
	database2.BillsInitDB()
}

func (data *BillDataByDate) BillsWriteToJSON() {
	var f *os.File
	src := strconv.Itoa(data.Year) + ".json"
	val, err := json.MarshalIndent(data.Data, "", "	") // 第二个表示每行的前缀，这里不用，第三个是缩进符号，这里用tab
	if err != nil {
		log.Writer(log.Error, err)
	}

	if service2.CheckFileIsExist(src) { //如果文件存在
		f, err = os.OpenFile(src, os.O_APPEND, 0666) //打开文件
	} else {
		f, err = os.Create(src) //创建文件

	}

	err = ioutil.WriteFile(src, val, 0777)
	service2.CheckErr(err)
	f.Close()
}

func (data *BillDataByDate) BillsGetYearData() {
	database2.BillsGetYearData(&data.Data, data.Year)
}

func (data *BillDataByDate) BillsGetDataByMonth() {
	database2.BillsGetDataByMonth(&data.Data)
}

func BillsGetTable(bill *billModels2.BillTable) *billModels2.BillTable {
	database2.BillsInitDB()
	database2.BillsGetTable(bill)
	return bill
}

func BillsGetTableOption() *BillOption {
	var option = new(BillOption)
	database2.BillsInitDB()
	option.BillName, option.BillType = database2.BillsGetTableOption()
	return option
}

func BillsGetDiagram(bill *billModels2.BillTable) ([]BillChartModel, error) {
	var data []BillChartModel

	database2.BillsInitDB()
	database2.BillsGetDiagram(bill)

	From(bill.BillDetail).GroupBy(func(i interface{}) interface{} {
		return i.(billModels2.BillDetail).BillName
	}, func(i interface{}) interface{} {
		return i.(billModels2.BillDetail)
	}).OrderBy(func(i interface{}) interface{} {
		return i.(Group).Key
	}).Select(func(group interface{}) interface{} {
		i := group.(Group)
		m := 0.0
		for _, item := range i.Group {
			m += item.(billModels2.BillDetail).Account
		}

		m, _ = decimal.NewFromFloat(m).Round(4).Float64()

		return BillChartModel{i.Key.(string), m}
	}).ToSlice(&data)
	return data, nil
}
