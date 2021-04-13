package BillModels

import (
	"Raven/src/Web/Service"
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"strconv"
	"strings"
	"time"
)

//Bills is test
type BillDetail struct {
	ID         int64     `db:"ID"`
	BillNumber string    `db:"BillNumber"`
	Type       string    `db:"Type"`
	BillName   string    `db:"BillName"`
	Account    string    `db:"Account"`
	Date       time.Time `db:"Date"`
	Remarks    string    `db:"Remarks"`
}

const (
	userName = "sa"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test_db"
)

var db *sql.DB
var engine *xorm.Engine

var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间

func BillsInitDB() {
	engine = Service.InitDB()
}

func BillsInitDBV1() {
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

func BillsGetYearData(data *BillData) {
	var bills []BillDetail

	var star, end string
	star = strconv.Itoa(data.Year)
	end = strconv.Itoa(data.Year + 1)
	star = star + "-01-01"
	end = end + "-01-01"
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	err := engine.Where("Date > ? and Date < ?", star, end).Find(&bills)
	if err != nil {
		fmt.Println(err)
	}
	data.Data = bills
}

func BillsGetYearDataV1(data *BillData) {
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
		fmt.Printf("Query failed,err:%v", err)
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
			fmt.Printf("scan failed, err:%v", err)
		}
		DefaultTimeLoc := time.Local

		billDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		Service.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data.Data = bills
}

func BillsGetFourMonthsData(data *BillData) {
	var bills []BillDetail

	err := engine.SQL("select * from BillDetail where Type = '支出' and Date > (select DATE_ADD(Max(date_format(Date,'%Y-%m-01') ),INTERVAL -3 Month) from BillDetail)").Find(&bills)
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
	}
	data.Data = bills
}

func BillsGetFourMonthsDataV1(data *BillData) {
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
		fmt.Printf("Query failed,err:%v", err)
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
			fmt.Printf("scan failed, err:%v", err)
		}
		DefaultTimeLoc := time.Local

		billDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		Service.CheckErr(err)
		bills = append(bills, *billDetail)
	}
	data.Data = bills
}
