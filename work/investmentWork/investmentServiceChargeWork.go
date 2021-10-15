package investmentWork

import (
	"encoding/json"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"io/ioutil"
	"xorm.io/xorm"
)

func actionInvestmentServiceCharge(session *xorm.Session) error {
	err := setInvestmentServiceCharge(session)
	if err != nil {
		return err
	}
	err = setInvestmentServiceChargeType(session)

	if err != nil {
		return err
	}
	return nil
}

func setInvestmentServiceCharge(session *xorm.Session) error {
	investmentServiceCharge := new(investmentsModels.InvestmentServiceCharge)

	flag, err := session.IsTableExist(investmentServiceCharge)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}
	if !flag {
		err = session.CreateTable(investmentServiceCharge)
		if err != nil {
			session.Rollback()
			Helheim.Writer(Helheim.Error, err)
			return err
		}
	}

	return nil
}

func setInvestmentServiceChargeType(session *xorm.Session) error {
	investmentServiceChargeType := new(investmentsModels.InvestmentServiceChargeType)

	flag, err := session.IsTableExist(investmentServiceChargeType)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	if flag {
		err = session.DropTable(investmentServiceChargeType)
		if err != nil {
			session.Rollback()
			Helheim.Writer(Helheim.Error, err)
			return err
		}

		err = session.CreateTable(investmentServiceChargeType)
		if err != nil {
			session.Rollback()
			Helheim.Writer(Helheim.Error, err)
			return err
		}
	} else {
		err = session.CreateTable(investmentServiceChargeType)
		if err != nil {
			session.Rollback()
			Helheim.Writer(Helheim.Error, err)
			return err
		}
	}

	addServiceChargeTypeInitialData(session)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	return nil
}

func addServiceChargeTypeInitialData(session *xorm.Session) error {
	var list []investmentsModels.InvestmentServiceChargeType
	err := ReadServiceChargeJSON("../initialData/investmentServiceChargeInitialData.json", &list)
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

func ReadServiceChargeJSON(address string, data *[]investmentsModels.InvestmentServiceChargeType) error {
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
