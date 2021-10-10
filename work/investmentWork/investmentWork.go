package investmentWork

import (
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"xorm.io/xorm"
)

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
	return RunTransaction(addDataToItem)
}

func addDataToItem(session *xorm.Session) error {
	var item []investmentsModels.InvestmentItem

	_, err := session.Exec("truncate table InvestmentItem;")
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	err = session.Table("Investment").GroupBy("ItemID,Name,Code").Find(&item)

	_, err = session.Insert(&item)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	return nil
}
