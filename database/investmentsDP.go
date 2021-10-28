package database

import (
	"database/sql"
	"fmt"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/common"
	. "github.com/swirling-melodies/Raven/models/investmentsModels"
	"strings"
)

func InvestmentsInitDB() {
	engine = common.InitDB()
}

func investmentsInitDBV1() {
	var err error
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err = sql.Open("mysql", path)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err = db.Ping(); err != nil {
		Helheim.Writer(Helheim.Error, err)
		return
	}
}

func InvestmentGetAll() []Investment {
	var investments []Investment

	err := engine.Find(&investments)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	return investments
}

func investmentGetAllV1() {
	var investments []Investment

	var investment []Investment

	//pEveryOne := make([]Investment, 0)
	err := engine.Find(&investment)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	investmentDetail := new(Investment)
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	/*	_, err := engine.Get(&investments)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
		}*/
	row, err := db.Query("select * from Investment")
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	defer func() {
		if row != nil {
			if err = row.Close(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
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
		//DefaultTimeLoc := time.Local

		//investmentDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)

		common.CheckErr(err)
		investments = append(investments, *investmentDetail)
	}
}

func InvestmentGetDataToChart() (*[]InvestmentChartModel, *[]InvestmentChartModel, *[]InvestmentChartModel) {
	var account, share, netWorth []InvestmentChartModel
	err := engine.SQL("select Name, sum(Account) Value from Investment where IsEmpty <> 1 group by Name").Find(&account)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	/*	err := engine.Table("Investment").GroupBy("Name").Select("select Name, sum(Account) Value").Find(&investmentsChartModel.Account)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
		}*/

	err = engine.SQL("select Name, sum(Share) Value from Investment where IsEmpty <> 1 group by Name").Find(&share)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	err = engine.SQL("select Name, avg(NetWorth) Value from Investment where where IsEmpty <> 1 and TypeID <> 3 group by Name").Find(&netWorth)

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	return &account, &share, &netWorth
}

func InvestmentGetChart() []InvestmentTable {
	var Item []InvestmentTable
	err := engine.Join("INNER", "InvestmentType",
		"InvestmentType.TypeID = Investment.TypeID").
		Where("IsEmpty <> ?", 1).And("ActivityStatus <> ?", 4).Find(&Item)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	return Item
}

func investmentGetDataToChartV1() {
	investmentDetail := new(InvestmentChartModel)
	var account, share, netWorth []InvestmentChartModel
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	row1, err := db.Query("select Name, sum(Account) from Investment group by Name")
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	defer func() {
		if row1 != nil {
			if err = row1.Close(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
		}
	}()

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	for row1.Next() {
		if err := row1.Scan(
			&investmentDetail.Name,
			&investmentDetail.Value,
		); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}

		common.CheckErr(err)
		account = append(account, *investmentDetail)
	}

	row2, err := db.Query("select Name, sum(Share) from Investment group by Name")
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	defer func() {
		if row2 != nil {
			if err = row2.Close(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
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

		common.CheckErr(err)
		share = append(share, *investmentDetail)
	}

	row3, err := db.Query("select Name, avg(NetWorth) from Investment where ID != 23 group by Name")
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	defer func() {
		if row3 != nil {
			if err = row3.Close(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
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

		common.CheckErr(err)
		netWorth = append(netWorth, *investmentDetail)
	}
}

func InvestmentGetTable() []InvestmentTable {
	var investments []InvestmentTable
	err := engine.Join("INNER", "InvestmentActivity",
		"InvestmentActivity.ActivityID = Investment.ActivityStatus").
		Join("INNER", "InvestmentType",
			"InvestmentType.TypeID = Investment.TypeID").Find(&investments)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	return investments
}

func investmentGetTableV1() []InvestmentTable {
	var investments []InvestmentTable
	investmentDetail := new(InvestmentTable)

	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	row, err := db.Query("select ID,Name,i.TypeID,Account,Share,NetWorth,Date,ActivityStatus,TypeName,ActivityName from Investment i join InvestmentActivity iA on i.ActivityStatus = iA.ActivityID join InvestmentType iT on i.TypeID = iT.TypeID")
	defer func() {
		if row != nil {
			if err = row.Close(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
		}
	}()

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
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
			Helheim.Writer(Helheim.Error, err)
		}
		//DefaultTimeLoc := time.Local
		if lastLoginTime != "" {
			//investmentDetail.Date, err = time.ParseInLocation(timeLayoutStr, lastLoginTime, DefaultTimeLoc)
		}

		common.CheckErr(err)
		investments = append(investments, *investmentDetail)
	}
	return investments
}

func InvestmentAddTable(data InvestmentTable) (bool, error) {

	if data.ActivityStatus == 4 {
		data.IsEmpty = true
	}

	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()

	if data.Investment.ItemID == 0 {
		var num int
		_, err = session.Table("Investment").Select("MAX(ItemID)").Get(&num)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		data.Investment.ItemID = num + 1
	}
	_, err = session.Insert(&data.Investment)

	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}

		return false, err
	}

	if data.Investment.IsEmpty == true {
		_, err = session.Exec("UPDATE Investment SET IsEmpty = 1 WHERE ItemID = ?", data.ItemID)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			if err = session.Rollback(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
			return false, err
		}
	}

	if data.ServiceChargeList != nil && len(data.ServiceChargeList) > 0 {
		for i := range data.ServiceChargeList {
			data.ServiceChargeList[i].ItemID = data.ID
		}
		_, err = session.Insert(&data.ServiceChargeList)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			if err = session.Rollback(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
			return false, err
		}
		for _, i := range data.ServiceChargeList {
			data.Investment.ServiceCharge += i.Cost
		}

		_, err = session.ID(data.ID).Cols("ServiceCharge").Update(&data.Investment)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			if err = session.Rollback(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
			return false, err
		}
	}

	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func investmentAddTableV1(data InvestmentTable) (bool, error) {
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	date := data.Date

	result, err := db.Exec("INSERT INTO Investment (`Name`,Account,Share,NetWorth,`Date`,TypeID, ActivityStatus)VALUES(?,?,?,?,?,?,?)", data.Name, data.Account, data.Share, data.NetWorth, date, data.TypeID, data.ActivityStatus)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if id > 0 {
		return true, nil
	}
	return false, err
}

func InvestmentUpdateTable(data InvestmentTable) (bool, error) {

	if data.ActivityStatus == 4 {
		data.IsEmpty = true
	}

	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()

	_, err = session.ID(data.ID).Cols("Code", "Name", "Account", "Share", "NetWorth", "Date", "TypeID", "ActivityStatus").Update(&data.Investment)
	if err != nil {
		return false, err
	}

	if data.Investment.IsEmpty == true {
		_, err = session.Exec("UPDATE Investment SET IsEmpty = 1 WHERE ItemID = ?", data.ItemID)
		if err != nil {
			if err = session.Rollback(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
			Helheim.Writer(Helheim.Error, err)
			return false, err
		}
	}

	if data.ServiceChargeList != nil {
		data.Investment.ServiceCharge = 0
		for _, i := range data.ServiceChargeList {
			data.Investment.ServiceCharge += i.Cost
		}

		_, err = session.ID(data.ID).Cols("ServiceCharge").Update(&data.Investment)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			if err = session.Rollback(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
			return false, err
		}

		for _, i := range data.ServiceChargeList {
			count, err := session.Where("ItemID = ?", i.ItemID).And("TypeID = ?", i.TypeID).Cols("Cost").Update(i)
			if err != nil {
				if err = session.Rollback(); err != nil {
					Helheim.Writer(Helheim.Error, err)
				}
				Helheim.Writer(Helheim.Error, err)
				return false, err
			}
			if count == 0 {
				_, err = session.Insert(i)
				if err != nil {
					if err = session.Rollback(); err != nil {
						Helheim.Writer(Helheim.Error, err)
					}
					Helheim.Writer(Helheim.Error, err)
					return false, err
				}
			}
		}
	}

	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func investmentUpdateTableV1(data InvestmentTable) (bool, error) {
	// db.QueryRow()调用完毕后会将连接传递给sql.Row类型，当.Scan()方法调用之后把连接释放回到连接池。

	// 查询单行数据
	// starDate, _ := time.Parse(timeLayoutStr, star)
	// endDate,_ := time.Parse(timeLayoutStr, end)

	insterDate := data.Date

	result, err := db.Exec("Update Investment Set `Name` = ?, Account = ?, Share = ?, NetWorth = ?, `Date` = ?, TypeID = ?, ActivityStatus = ? Where ID = ?",
		data.Name, data.Account, data.Share, data.NetWorth, insterDate, data.TypeID, data.ActivityStatus, data.ID)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if id > 0 {
		return true, nil
	}
	return false, err
}

func InvestmentGetOption() ([]InvestmentType, []InvestmentActivity, []InvestmentItem, []InvestmentServiceChargeType, error) {
	var itype []InvestmentType
	var iactivity []InvestmentActivity
	var item []InvestmentItem
	var serviceCharge []InvestmentServiceChargeType

	err := engine.Find(&itype)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	err = engine.Find(&iactivity)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	err = engine.Find(&item)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	err = engine.Find(&serviceCharge)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	return itype, iactivity, item, serviceCharge, nil
}

func InvestmentGetDateOrderbyDate() *[]Investment {
	var investments []Investment
	err := engine.OrderBy("Date").Find(&investments)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	return &investments
}

func GetServiceChargeData(itemID int) []InvestmentServiceCharge {
	var list []InvestmentServiceCharge
	err := engine.Where("ItemID = ?", itemID).Find(&list)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	return list
}
