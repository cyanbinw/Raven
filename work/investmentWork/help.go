package investmentWork

import (
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/common"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type Transaction func(session *xorm.Session) error

func initDB() {
	engine = common.InitDB()
}

func RunTransaction(transaction Transaction) (bool, error) {
	initDB()
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		Helheim.Writer(Helheim.Error, err)
	}

	err := transaction(session)

	if err != nil {
		session.Rollback()
		Helheim.Writer(Helheim.Error, err)
	}

	session.Commit()

	return err == nil, err
}
