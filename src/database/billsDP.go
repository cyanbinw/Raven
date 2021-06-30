package database

import (
	"Raven/src/log"
	billModels2 "Raven/src/models/billModels"
	service2 "Raven/src/service"
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"strconv"
	"strings"
	"time"
)

func BillsInitDB() {
	engine = service2.InitDB()
}

func billsInitDBV1() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
}

func BillsGetYearData(data *[]billModels2.BillDetail, year int) {
	var bills []billModels2.BillDetail

	var star, end string
	star = strconv.Itoa(year)
	end = strconv.Itoa(year + 1)
	star = star + "-01-01"
	end = end + "-01-01"
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	err := engine.Where("Date > ? and Date < ?", star, end).Desc("Date").Find(bills)
	if err != nil {
		fmt.Println(err)
	}
	data = &bills
}

func billsGetYearDataV1(data *[]billModels2.BillDetail, year int) {
	var bills []billModels2.BillDetail
	billDetail := new(billModels2.BillDetail)
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	var star, end string
	star = strconv.Itoa(year)
	end = strconv.Itoa(year + 1)
	star = star + "-01-01"
	end = end + "-01-01"
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	row, err := db.Query("select * from BillDetail where Date > ? and Date < ?", star, end)
	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	if err != nil {
		log.Writer(log.Error, err)
	}

	for row.Next() {
		var lastLoginTime string
		if err := row.Scan(
			&billDetail.ID,
			&billDetail.BillNumber,
			&billDetail.Type,
			&billDetail.BillName,
			&billDetail.Account,
			&lastLoginTime,
			&billDetail.Remarks,
		); err != nil {
			log.Writer(log.Error, err)
		}
		DefaultTimeLoc := time.Local

		billDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		service2.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data = &bills
}

func BillsGetDataByMonth(data *[]billModels2.BillDetail) {

	err := engine.SQL("select * from BillDetail where Type = '支出' and Date > (select DATE_ADD(Max(date_format(Date,'%Y-%m-01') ),INTERVAL -? Month) from BillDetail) ORDER BY Date DESC", month-1).Find(data)
	if err != nil {
		log.Writer(log.Error, err)
	}
}

func billsGetFourMonthsDataV1(data *[]billModels2.BillDetail, year int) {
	var bills []billModels2.BillDetail
	billDetail := new(billModels2.BillDetail)
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	var star, end string
	star = strconv.Itoa(year)
	end = strconv.Itoa(year + 1)
	star = star + "-01-01"
	end = end + "-01-01"
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	row, err := db.Query("select * from BillDetail where Type = '支出' and Date > (select DATE_ADD(Max(date_format(Date,'%Y-%m-01') ),INTERVAL -3 Month) from BillDetail)")
	defer func() {
		if row != nil {
			row.Close()
		}
	}()

	if err != nil {
		log.Writer(log.Error, err)
	}

	for row.Next() {
		var lastLoginTime string
		if err := row.Scan(
			&billDetail.ID,
			&billDetail.BillNumber,
			&billDetail.Type,
			&billDetail.BillName,
			&billDetail.Account,
			&lastLoginTime,
			&billDetail.Remarks,
		); err != nil {
			log.Writer(log.Error, err)
		}
		DefaultTimeLoc := time.Local

		billDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		service2.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data = &bills
}

func BillsGetTable(bill *billModels2.BillTable) {
	data := new(billModels2.BillDetail)

	row := engine.Table("BillDetail")

	setBillsGetTableOption(row, bill)

	err := row.Limit(bill.PageSize, (bill.PageNumber-1)*bill.PageSize).Find(&bill.BillDetail)
	if err != nil {
		log.Writer(log.Error, err)
	}

	count := engine.Table("BillDetail")

	setBillsGetTableOption(count, bill)

	bill.Total, _ = count.Count(data)
}

func setBillsGetTableOption(search *xorm.Session, bill *billModels2.BillTable) {
	if len(bill.BillType) == 0 && len(bill.BillName) == 0 &&
		bill.AccountMax == 0.0 && bill.AccountMin == 0.0 &&
		bill.DateMin.IsZero() && bill.DateMax.IsZero() {
		return
	}

	search = search.Where("")
	if len(bill.BillType) > 0 {
		for _, i := range bill.BillType {
			search = search.And("Type = ?", i)
		}
	}

	if len(bill.BillName) > 0 {
		for _, i := range bill.BillName {
			search = search.And("BillName = ?", i)
		}

	}

	if bill.AccountMax != 0.0 || bill.AccountMin != 0.0 {

		if bill.AccountMin > bill.AccountMax {
			num := bill.AccountMax
			bill.AccountMax = bill.AccountMin
			bill.AccountMin = num
		}

		if bill.AccountMax != 0.0 {
			search = search.And("Account <= ?", bill.AccountMax)
		}

		if bill.AccountMin != 0.0 {
			search = search.And("Account >= ?", bill.AccountMin)
		}
	}

	if !bill.DateMin.IsZero() {
		search = search.And("Date >= ?", bill.DateMin)
	}

	if !bill.DateMin.IsZero() {
		search = search.And("Date <= ?", bill.DateMax)
	}
}

func BillsGetTableOption() ([]string, []string) {

	var billName, billType []string

	err := engine.Table("BillDetail").GroupBy("BillName").Find(&billName)
	if err != nil {
		log.Writer(log.Error, err)
	}
	err = engine.Table("BillDetail").GroupBy("Type").Find(&billType)
	if err != nil {
		log.Writer(log.Error, err)
	}
	return billName, billType
}

func BillsGetDiagram(bill *billModels2.BillTable) {
	row := engine.Table("BillDetail")

	setBillsGetTableOption(row, bill)

	err := row.Find(&bill.BillDetail)
	if err != nil {
		log.Writer(log.Error, err)
	}
}
