package billModels

import (
	"Raven/src/log"
	"Raven/src/web/service"
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"strconv"
	"strings"
	"time"
)

//BillDetail is test
type BillDetail struct {
	ID         int64     `db:"ID"`
	BillNumber string    `db:"BillNumber"`
	Type       string    `db:"Type"`
	BillName   string    `db:"BillName"`
	Account    float64   `db:"Account"`
	Date       time.Time `db:"Date"`
	Remarks    string    `db:"Remarks"`
}

type BillTable struct {
	BillDetail []BillDetail
	PageSize   int
	Total      int64
	PageNumber int
	BillType   []string
	BillName   []string
	AccountMax float64
	AccountMin float64
	DateMax    time.Time
	DateMin    time.Time
}

type BillOption struct {
	BillName []string
	BillType []string
}

const (
	userName = ""
	password = ""
	ip       = ""
	port     = ""
	dbName   = ""
	month    = 12
)

var db *sql.DB
var engine *xorm.Engine

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间

func billsInitDB() {
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

func billsGetYearData(data *BillDataByDate) {
	var bills []BillDetail

	var star, end string
	star = strconv.Itoa(data.Year)
	end = strconv.Itoa(data.Year + 1)
	star = star + "-01-01"
	end = end + "-01-01"
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	err := engine.Where("Date > ? and Date < ?", star, end).Desc("Date").Find(&bills)
	if err != nil {
		fmt.Println(err)
	}
	data.Data = bills
}

func billsGetYearDataV1(data *BillDataByDate) {
	var bills []BillDetail
	billDetail := new(BillDetail)
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	var star, end string
	star = strconv.Itoa(data.Year)
	end = strconv.Itoa(data.Year + 1)
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

		service.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data.Data = bills
}

func billsGetDataByMonth(data *BillDataByDate) {

	err := engine.SQL("select * from BillDetail where Type = '支出' and Date > (select DATE_ADD(Max(date_format(Date,'%Y-%m-01') ),INTERVAL -? Month) from BillDetail) ORDER BY Date DESC", month-1).Find(&data.Data)
	if err != nil {
		log.Writer(log.Error, err)
	}
}

func billsGetFourMonthsDataV1(data *BillDataByDate) {
	var bills []BillDetail
	billDetail := new(BillDetail)
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	var star, end string
	star = strconv.Itoa(data.Year)
	end = strconv.Itoa(data.Year + 1)
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

		service.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data.Data = bills
}

func billsGetTable(bill *BillTable) {
	data := new(BillDetail)

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

func setBillsGetTableOption(search *xorm.Session, bill *BillTable) {
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

func billsGetTableOption() *BillOption {

	var option BillOption

	err := engine.Table("BillDetail").GroupBy("BillName").Find(&option.BillName)
	if err != nil {
		log.Writer(log.Error, err)
	}
	err = engine.Table("BillDetail").GroupBy("Type").Find(&option.BillType)
	if err != nil {
		log.Writer(log.Error, err)
	}
	return &option
}

func billsGetDiagram(bill *BillTable) {
	row := engine.Table("BillDetail")

	setBillsGetTableOption(row, bill)

	err := row.Find(&bill.BillDetail)
	if err != nil {
		log.Writer(log.Error, err)
	}
}
