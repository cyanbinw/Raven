package database

import (
	"database/sql"
	"fmt"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/models/billModels"
	"github.com/swirling-melodies/Raven/service"
	"strconv"
	"strings"
	"xorm.io/xorm"
)

func BillsInitDB() {
	engine = service.InitDB()
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

func BillsGetYearData(data *[]billModels.BillDetail, year int) {
	var bills []billModels.BillDetail

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

func billsGetYearDataV1(data *[]billModels.BillDetail, year int) {
	var bills []billModels.BillDetail
	billDetail := new(billModels.BillDetail)
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
		Helheim.Writer(Helheim.Error, err)
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
			Helheim.Writer(Helheim.Error, err)
		}
		//DefaultTimeLoc := time.Local

		//billDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		service.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data = &bills
}

func BillsGetDataByMonth(data *[]billModels.BillDetail) {

	err := engine.SQL("select * from BillDetail where Type = '支出' and Date > (select DATE_ADD(Max(date_format(Date,'%Y-%m-01') ),INTERVAL -? Month) from BillDetail) ORDER BY Date DESC", month-1).Find(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
}

func billsGetFourMonthsDataV1(data *[]billModels.BillDetail, year int) {
	var bills []billModels.BillDetail
	billDetail := new(billModels.BillDetail)
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
		Helheim.Writer(Helheim.Error, err)
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
			Helheim.Writer(Helheim.Error, err)
		}
		//DefaultTimeLoc := time.Local

		//billDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		service.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data = &bills
}

func BillsGetTable(bill *billModels.BillTable) {
	data := new(billModels.BillDetail)

	row := engine.Table("BillDetail")

	setBillsGetTableOption(row, bill)

	err := row.Limit(bill.PageSize, (bill.PageNumber-1)*bill.PageSize).Find(&bill.BillDetail)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	count := engine.Table("BillDetail")

	setBillsGetTableOption(count, bill)

	bill.Total, _ = count.Count(data)
}

func setBillsGetTableOption(search *xorm.Session, bill *billModels.BillTable) {
	if len(bill.BillType) == 0 && len(bill.BillName) == 0 &&
		bill.AccountMax == 0.0 && bill.AccountMin == 0.0 &&
		bill.DateMin.IsZero() && bill.DateMax.IsZero() {
		search = search.Desc("Date")
		return
	}

	search = search.Where("")
	if len(bill.BillType) > 0 {
		var str string
		for _, i := range bill.BillType {
			str += "'" + i + "',"
		}
		str = str[:len(str)-1]
		search = search.And("Type in (" + str + ")")
	}

	if len(bill.BillName) > 0 {
		var str string
		for _, i := range bill.BillName {
			str += "'" + i + "',"
		}
		str = str[:len(str)-1]

		search = search.And("BillName in (" + str + ")")
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

	search = search.Desc("Date")
}

func BillsGetTableOption() ([]string, []string) {

	var billName, billType []string

	err := engine.Table("BillDetail").GroupBy("BillName").Find(&billName)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	err = engine.Table("BillDetail").GroupBy("Type").Find(&billType)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	return billName, billType
}

func BillsGetDiagram(bill *billModels.BillTable) {
	row := engine.Table("BillDetail")

	setBillsGetTableOption(row, bill)

	err := row.Find(&bill.BillDetail)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
}

func BillsGetDataByPage(bill *billModels.BillDataByPage) {
	err := engine.Join("INNER", "BillNameConfig",
		"BillNameConfig.BillName = BillDetail.BillName").
		Desc("Date").
		Limit(bill.PageSize, (bill.PageNumber-1)*bill.PageSize).Find(&bill.BillData)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
}
