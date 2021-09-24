package investmentWork

import (
	"encoding/json"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"io/ioutil"
	"xorm.io/xorm"
)

func SetInvestmentType() (bool, error) {
	return RunTransaction(actionInvestmentType)
}

func actionInvestmentType(session *xorm.Session) error {
	investmentType := new(investmentsModels.InvestmentType)

	err := session.DropTable(investmentType)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	err = session.CreateTable(investmentType)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	err = addInitialData(session)
	if err != nil {
		return err
	}
	return nil
}

func addInitialData(session *xorm.Session) error {
	var list []investmentsModels.InvestmentType
	err := ReadJSON("../initialData/investmentTypeinitialData.json", &list)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}
	_, err = session.Insert(&list)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}
	return nil
}

func ReadJSON(address string, data *[]investmentsModels.InvestmentType) error {
	bytes, err := ioutil.ReadFile(address)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return err
	}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return err
	}
	return nil
}
