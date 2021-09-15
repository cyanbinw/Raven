package investmentWork

import (
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/common"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func initDB() {
	engine = common.InitDB()
}

func SetInvestmentItem() (bool, error) {
	item := new(investmentsModels.InvestmentItem)

	initDB()

	flag, err := engine.IsTableExist(item)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(item)
	}
	return addDataToItem()
}

func addDataToItem() (bool, error) {
	var item []investmentsModels.InvestmentItem

	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}

	_, err := session.Exec("truncate table InvestmentItem;")
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}

	err = session.Table("Investment").GroupBy("ItemID,Name,Code").Find(&item)

	_, err = session.Insert(&item)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}

	session.Commit()

	return true, nil
}
