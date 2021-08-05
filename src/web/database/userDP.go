package database

import (
	"github.com/WFallenDown/Helheim"
	"github.com/WFallenDown/Raven/src/web/models/userModels"
	"github.com/WFallenDown/Raven/src/web/service"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
	"time"
)

func UserInitDB() {
	engine = service.InitDB()
}

func Login(data *userModels.UserInfo) (bool, error) {

	flag, err := engine.Where("UserName = ?", data.UserName).And("Password = ?", data.Password).Get(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return flag, err
	}

	return flag, nil
}

func CreateToken(data *userModels.Token) error {
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	tokenAudit := new(userModels.TokenAudit)

	_, err = engine.Where("UserID = ?", data.UserID).Get(data)

	row, err := engine.ID(data.TokenNum).Delete(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}

	if row > 0 {
		tokenAudit.AuditType = delete
		tokenAudit.Token = *data
		_, err = engine.Insert(tokenAudit)
	}

	data.TokenNum, err = uuid.New()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}
	data.UpdateTokenNum, err = uuid.New()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}
	data.StratTime = time.Now().Local()
	data.EndTime = data.StratTime.AddDate(0, 0, userModels.EffectiveTime)
	data.UpdateTokenTime = data.StratTime.AddDate(0, 0, userModels.UpdateEffectiveTime)
	_, err = engine.Insert(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}

	tokenAudit = new(userModels.TokenAudit)
	tokenAudit.Token = *data
	tokenAudit.AuditType = insert
	_, err = engine.Insert(tokenAudit)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}

	err = session.Commit()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	return nil
}

func UpdateToken(data *userModels.Token) error {
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	tokenAudit := new(userModels.TokenAudit)

	_, err = engine.Where("TokenNum = ?", data.TokenNum).Get(data)

	data.StratTime = time.Now().Local()
	data.EndTime = data.StratTime.AddDate(0, 0, userModels.EffectiveTime)
	data.UpdateTokenTime = data.StratTime.AddDate(0, 0, userModels.UpdateEffectiveTime)
	row, err := engine.ID(data.TokenNum).Update(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}

	if row > 0 {
		tokenAudit.AuditType = update
		tokenAudit.Token = *data
		_, err = engine.Insert(tokenAudit)
	}

	err = session.Commit()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	return nil
}

func ValidateToken(data uuid.UUID) (bool, error) {
	token := new(userModels.Token)
	row, err := engine.Where("TokenNum = ?", data).Get(token)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	return row, nil
}
