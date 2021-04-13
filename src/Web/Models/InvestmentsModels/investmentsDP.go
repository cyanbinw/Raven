package InvestmentsModels

import (
	"Raven/src/Web/Service"
	"database/sql"
	"fmt"
	"github.com/go-xorm/xorm"
	"strings"
	"time"
)

type Investment struct {
	ID             int64     `db:"ID" json:"id"`
	Name           string    `db:"Name" json:"name"`
	TypeID         int       `db:"TypeID" json:"type"`
	Account        float32   `db:"Account" json:"account"`
	Share          float32   `db:"Share" json:"share"`
	NetWorth       float32   `db:"NetWorth" json:"netWorth"`
	Date           time.Time `db:"Date" json:"date"`
	ActivityStatus int       `db:"ActivityStatus" json:"activity"`
}

type InvestmentTable struct {
	Investment   `xorm:"extends"`
	ActivityName string `json:"activityName"`
	TypeName     string `json:"typeName"`
}

type InvestmentActivity struct {
	ActivityID   int
	ActivityName string
	InsertDate   time.Time
}

type InvestmentType struct {
	TypeID     int
	TypeName   string
	InsertDate time.Time
}

type InvestmentGroup struct {
	Data  []Investment
	Count int
	Name  string
}

type InvestmentGroupList []InvestmentGroup

const (
	userName = "sa"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "test_db"
)

var db *sql.DB
var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
var engine *xorm.Engine

func InvestmentsInitDB() {
	engine = Service.InitDB()
}

func InvestmentsInitDBV1() {
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

func InvestmentGetAll(data *InvestmentData) {
	var investments []Investment

	//pEveryOne := make([]Inengine = {*github.com/go-xorm/xorm.Engine | 0xc0003161c0} vestment, 0)
	err := engine.Find(&investments)
	if err != nil {
		fmt.Println(err)
	}
	data.Data = investments
}

func InvestmentGetAllV1(data *InvestmentData) {
	var investments []Investment

	var investment []Investment

	//pEveryOne := make([]Investment, 0)
	err := engine.Find(&investment)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(investment)

	investmentDetail := new(Investment)
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	/*	_, err := engine.Get(&investments)
		if err != nil {
			fmt.Println(err)
		}*/
	row, err := db.Query("select * from Investment")
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
			&investmentDetail.ID,
			&investmentDetail.Name,
			&investmentDetail.TypeID,
			&investmentDetail.Account,
			&investmentDetail.Share,
			&investmentDetail.NetWorth,
			&lastLoginTime,
			&investmentDetail.ActivityStatus,
		); err != nil {
			fmt.Printf("scan failed, err:%v", err)
		}
		DefaultTimeLoc := time.Local

		investmentDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		Service.CheckErr(err)
		investments = append(investments, *investmentDetail)
	}
	data.Data = investments
}

func (InvestmentTable) TableName() string {
	return "Investment"
}

func InvestmentGetDataToChart() InvestmentsChartModel {
	var investmentsChartModel InvestmentsChartModel
	err := engine.SQL("select Name, sum(Account) Value from Investment group by Name").Find(&investmentsChartModel.Account)
	if err != nil {
		fmt.Println(err)
	}

	/*	err := engine.Table("Investment").GroupBy("Name").Select("select Name, sum(Account) Value").Find(&investmentsChartModel.Account)
		if err != nil {
			fmt.Println(err)
		}*/

	err = engine.SQL("select Name, sum(Share) Value from Investment group by Name").Find(&investmentsChartModel.Share)
	if err != nil {
		fmt.Println(err)
	}

	err = engine.SQL("select Name, avg(NetWorth) Value from Investment where ID != 23 group by Name").Find(&investmentsChartModel.NetWorth)

	return investmentsChartModel
}

func InvestmentGetDataToChartV1() InvestmentsChartModel {
	investmentDetail := new(InvestmentChartModel)
	var investmentsChartModel InvestmentsChartModel
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	row1, err := db.Query("select Name, sum(Account) from Investment group by Name")
	defer func() {
		if row1 != nil {
			row1.Close()
		}
	}()

	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
	}

	for row1.Next() {
		if err := row1.Scan(
			&investmentDetail.Name,
			&investmentDetail.Value,
		); err != nil {
			fmt.Printf("scan failed, err:%v", err)
		}

		Service.CheckErr(err)
		investmentsChartModel.Account = append(investmentsChartModel.Account, *investmentDetail)
	}

	row2, err := db.Query("select Name, sum(Share) from Investment group by Name")
	defer func() {
		if row2 != nil {
			row2.Close()
		}
	}()

	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
	}

	for row2.Next() {
		if err := row2.Scan(
			&investmentDetail.Name,
			&investmentDetail.Value,
		); err != nil {
			fmt.Printf("scan failed, err:%v", err)
		}

		Service.CheckErr(err)
		investmentsChartModel.Share = append(investmentsChartModel.Share, *investmentDetail)
	}

	row3, err := db.Query("select Name, avg(NetWorth) from Investment where ID != 23 group by Name")
	defer func() {
		if row3 != nil {
			row3.Close()
		}
	}()

	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
	}

	for row3.Next() {
		if err := row3.Scan(
			&investmentDetail.Name,
			&investmentDetail.Value,
		); err != nil {
			fmt.Printf("scan failed, err:%v", err)
		}

		Service.CheckErr(err)
		investmentsChartModel.NetWorth = append(investmentsChartModel.NetWorth, *investmentDetail)
	}

	return investmentsChartModel
}

func InvestmentGetTable() []InvestmentTable {
	var investments []InvestmentTable
	err := engine.Join("INNER", "InvestmentActivity",
		"InvestmentActivity.ActivityID = Investment.ActivityStatus").
		Join("INNER", "InvestmentType",
			"InvestmentType.TypeID = Investment.TypeID").Find(&investments)
	if err != nil {
		fmt.Println(err)
	}
	return investments
}

func InvestmentGetTableV1() []InvestmentTable {
	var investments []InvestmentTable
	investmentDetail := new(InvestmentTable)

	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	row, err := db.Query("select ID,Name,i.TypeID,Account,Share,NetWorth,Date,ActivityStatus,TypeName,ActivityName from Investment i join InvestmentActivity iA on i.ActivityStatus = iA.ActivityID join InvestmentType iT on i.TypeID = iT.TypeID")
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
			&investmentDetail.ID,
			&investmentDetail.Name,
			&investmentDetail.Investment.TypeID,
			&investmentDetail.Account,
			&investmentDetail.Share,
			&investmentDetail.NetWorth,
			&lastLoginTime,
			&investmentDetail.ActivityStatus,
			&investmentDetail.TypeName,
			&investmentDetail.ActivityName,
		); err != nil {
			fmt.Printf("scan failed, err:%v", err)
		}
		DefaultTimeLoc := time.Local
		if lastLoginTime != "" {
			investmentDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)
		}

		Service.CheckErr(err)
		investments = append(investments, *investmentDetail)
	}
	return investments
}

func InvestmentAddTable(data InvestmentTable) (bool, error) {
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()

	_, err = engine.Insert(&data.Investment)
	if err != nil {
		session.Rollback()
		return false, err
	}

	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func InvestmentAddTableV1(data InvestmentTable) (bool, error) {
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	date := data.Date.Local().Format("2006-01-02")

	result, err := db.Exec("INSERT INTO Investment (`Name`,Account,Share,NetWorth,`Date`,TypeID, ActivityStatus)VALUES(?,?,?,?,?,?,?)", data.Name, data.Account, data.Share, data.NetWorth, date, data.TypeID, data.ActivityStatus)
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return false, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return false, err
	}
	if id > 0 {
		return true, nil
	}
	return false, err
}

func InvestmentUpdateTable(data InvestmentTable) (bool, error) {
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()

	_, err = engine.ID(data.ID).Update(&data.Investment)
	if err != nil {
		return false, err
	}

	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func InvestmentUpdateTableV1(data InvestmentTable) (bool, error) {
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	insterDate := data.Date.Local().Format("2006-01-02")

	result, err := db.Exec("Update Investment Set `Name` = ?, Account = ?, Share = ?, NetWorth = ?, `Date` = ?, TypeID = ?, ActivityStatus = ? Where ID = ?",
		data.Name, data.Account, data.Share, data.NetWorth, insterDate, data.TypeID, data.ActivityStatus, data.ID)
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return false, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return false, err
	}
	if id > 0 {
		return true, nil
	}
	return false, err
}

func InvestmentGetDiagram() (map[string][]Investment, error) {
	var table = new(InvestmentData)
	data := make(map[string][]Investment)

	InvestmentGetAll(table)
	for _, index := range table.Data {
		data[index.Name] = append(data[index.Name], index)
	}
	return data, nil
}

func InvestmentGetOption() ([]InvestmentType, []InvestmentActivity, error) {
	var itype []InvestmentType
	var iactivity []InvestmentActivity

	err := engine.Find(&itype)
	if err != nil {
		fmt.Println(err)
	}

	err = engine.Find(&iactivity)

	return itype, iactivity, nil
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
