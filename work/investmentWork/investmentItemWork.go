package investmentWork

import (
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"xorm.io/xorm"
)

func actionItem(session *xorm.Session) error {
	item := new(investmentsModels.InvestmentItem)

	flag, err := session.IsTableExist(item)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return err
	}
	if !flag {
		engine.CreateTables(item)
	}

	addDataToItem(session)

	return nil
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
