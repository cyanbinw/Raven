package billNameService

import (
	"github.com/WFallenDown/Helheim"
	. "github.com/WFallenDown/Raven/src/web/models/billModels"
	"github.com/WFallenDown/Raven/src/web/service"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

const (
	create = 1
	update = 2
	delect = 3
)

func initDB() {
	engine = service.InitDB()
}

func SetBillName() (bool, error) {
	var bill = new(BillNameConfig)
	var audit = new(BillNameConfigAudit)
	var bills BillDetail
	var name []BillNameConfig
	initDB()

	flag, err := engine.IsTableExist(bill)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(bill)
	}

	flag, err = engine.IsTableExist(audit)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	if !flag {
		engine.CreateTables(audit)
	}

	err = engine.Table(bills).GroupBy("BillName").OrderBy("BillName").Select("BillName, count(1) AS Count").Find(&name)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	err = addORUpdate(&name)
	if err != nil {
		Helheim.Writer(Helheim.Error, "Work error")
		return false, err
	}
	return true, nil
}

func addORUpdate(bills *[]BillNameConfig) error {
	var bill = new(BillNameConfig)

	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		return err
	}

	for _, data := range *bills {
		flag, err := engine.Where("BillName = ?", data.BillName).Get(bill)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			return err
		}
		if !flag {
			_, err := engine.Insert(&data)
			if err != nil {
				Helheim.Writer(Helheim.Error, err)
				return err
			}
			audit := setAudit(&data, create)
			_, err = engine.Insert(audit)
			if err != nil {
				Helheim.Writer(Helheim.Error, err)
				return err
			}
		} else {
			if bill.Count != data.Count {
				bill.Count = data.Count

				_, err := engine.ID(bill.ID).Update(bill)
				if err != nil {
					Helheim.Writer(Helheim.Error, err)
					return err
				}
				audit := setAudit(bill, update)
				_, err = engine.Insert(audit)
				if err != nil {
					Helheim.Writer(Helheim.Error, err)
					return err
				}
			}
		}
	}
	return session.Commit()
}

func setAudit(data *BillNameConfig, status int) *BillNameConfigAudit {
	var audit = new(BillNameConfigAudit)
	audit.ID = 0
	audit.BillID = data.ID
	audit.BillName = data.BillName
	audit.Count = data.Count
	audit.Color = data.Color
	audit.Icon = data.Icon
	audit.UpdateDate = data.UpdateDate
	audit.CreatDate = data.CreatDate
	audit.Status = status
	return audit
}

func GetBillNameList() *[]BillNameConfig {
	initDB()
	var data []BillNameConfig
	err := engine.Find(&data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
	}
	return &data
}

func UpdateBillName(data *BillNameConfig) bool {
	initDB()
	_, err := engine.ID(data.ID).Cols("Color", "Icon").Update(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false
	}
	return true
}
