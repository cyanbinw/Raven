package investmentWork

import (
	"errors"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/models/investmentsModels"
	"xorm.io/xorm"
)

func SetInvestmentType() (bool, error) {
	return RunTransaction(actionInvestmentType)
}

func actionInvestmentType(session *xorm.Session) error {
	investmentType := new(investmentsModels.InvestmentType)

	var investmentTypeList []investmentsModels.InvestmentType

	err := session.Find(&investmentTypeList)
	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
	}

	if len(investmentTypeList) == 0 {
		session.Rollback()
		return errors.New("the size of the investment type is empty")
	}

	err = session.DropTable(investmentType)
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

	return nil
}
