package billNameWork

import (
	"Raven/src/log"
	. "Raven/src/models/billModels"
	"Raven/src/service"
	"fmt"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func initDB() {
	engine = service.InitDB()
}

func SetBillName() (bool, error) {
	var bill = new(BillName)
	var bills BillDetail
	var name []BillName
	initDB()

	flag, err := engine.IsTableExist(bill)
	if err != nil {
		log.Writer(log.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(bill)
	}

	err = engine.Table(bills).GroupBy("BillName").OrderBy("BillName").Select("BillName").Find(&name)
	if err != nil {
		log.Writer(log.Error, err)
		return false, err
	}
	err = addORUpdate(name)
	if err != nil {
		log.Writer(log.Error, "Work error")
		return false, err
	}
	return true, nil
}

func addORUpdate(bills []BillName) error {
	var bill = new(BillName)

	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return err
	}

	for _, data := range bills {
		flag, err := engine.Where("BillName = ?", data.BillName).Get(bill)
		if err != nil {
			log.Writer(log.Error, err)
			return err
		}
		if !flag {
			row, err := engine.Insert(data)
			if err != nil {
				log.Writer(log.Error, err)
				return err
			}
			fmt.Println(row)
		} else {
			row, err := engine.ID(bill.ID).Update(bill)
			if err != nil {
				log.Writer(log.Error, err)
				return err
			}
			fmt.Println(row)
		}
	}
	return session.Commit()
}
