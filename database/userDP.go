package database

import (
	"github.com/satori/go.uuid"
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/common"
	"github.com/swirling-melodies/Raven/models/userModels"
	"time"
)

func UserInitDB() {
	engine = common.InitDB()
}

func Login(data *UserInfo) (bool, error) {
	flag, err := engine.Where("UserName = ?", data.UserName).And("Password = ?", data.Password).Get(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return flag, err
	}

	return flag, nil
}

func CreateToken(data *TokenInfo) error {
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	tokenAudit := new(userModels.TokenInfoAudit)

	flag, err := session.Where("UserID = ?", data.UserID).Get(data)

	if flag {
		row, err := session.ID(data.TokenNum).Delete(data)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			if err = session.Rollback(); err != nil {
				Helheim.Writer(Helheim.Error, err)
			}
			return err
		}
		if row > 0 {
			tokenAudit = setTokenAudit(data)
			tokenAudit.AuditType = delete
			_, err = session.Insert(tokenAudit)
		}
	}

	data.TokenNum = uuid.NewV4().String()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}

	data.UpdateTokenNum = uuid.NewV4().String()
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
	_, err = session.Insert(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}

	tokenAudit = setTokenAudit(data)
	tokenAudit.AuditType = insert
	_, err = session.Insert(tokenAudit)
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

func UpdateToken(data *TokenInfo) error {
	session := engine.NewSession()
	defer session.Close()

	err := session.Begin()
	tokenAudit := new(userModels.TokenInfoAudit)

	_, err = session.Where("TokenNum = ?", data.TokenNum).Get(data)

	data.StratTime = time.Now().Local()
	data.EndTime = data.StratTime.AddDate(0, 0, userModels.EffectiveTime)
	data.UpdateTokenTime = data.StratTime.AddDate(0, 0, userModels.UpdateEffectiveTime)
	row, err := session.ID(data.TokenNum).Update(data)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		if err = session.Rollback(); err != nil {
			Helheim.Writer(Helheim.Error, err)
		}
		return err
	}

	if row > 0 {
		tokenAudit = setTokenAudit(data)
		tokenAudit.AuditType = update
		_, err = session.Insert(tokenAudit)
	}

	err = session.Commit()
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return err
	}

	return nil
}

func ValidateToken(data string) (bool, error) {
	token := new(userModels.TokenInfo)
	row, err := engine.Where("TokenNum = ?", data).Get(token)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	return row, nil
}

func setTokenAudit(data *TokenInfo) *userModels.TokenInfoAudit {
	audit := userModels.TokenInfoAudit{
		UserID:          data.UserID,
		TokenNum:        data.TokenNum,
		UpdateTokenNum:  data.UpdateTokenNum,
		StratTime:       data.StratTime,
		EndTime:         data.EndTime,
		UpdateTokenTime: data.UpdateTokenTime,
	}

	return &audit
}
