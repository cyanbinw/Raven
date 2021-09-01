package application

import (
	"github.com/swirling-melodies/Helheim"
	"github.com/swirling-melodies/Raven/database"
	"github.com/swirling-melodies/Raven/models/userModels"
)

type User struct {
	*userModels.UserInfo
	*userModels.TokenInfo
}

func (User) UserInitDB() {
	database.UserInitDB()
}

func (data *User) Login() (bool, error) {
	flag, err := database.Login(data.UserInfo)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return flag, err
	}

	if flag {
		data.TokenInfo = new(userModels.TokenInfo)
		data.UserID = data.ID
		err := database.CreateToken(data.TokenInfo)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			return false, err
		}
		return flag, nil
	}
	return false, nil
}

func (data *User) Register() (bool, error) {
	return false, nil
}

func (data *User) GetUser() (bool, error) {
	panic("implement me")
}

func (data *User) GetUsers() ([]userModels.UserInfo, error) {
	panic("implement me")
}

func (data *User) SetUserInfo() (bool, error) {
	panic("implement me")
}

func (data *User) Delete() (bool, error) {
	panic("implement me")
}

func (data *User) CreateToken() (bool, error) {
	err := database.CreateToken(data.TokenInfo)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	return false, nil
}

func (data *User) ValidateToken() (bool, error) {
	flag, err := database.ValidateToken(data.TokenNum)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	return flag, nil
}

func (data *User) UpdateToken() (bool, error) {
	err := database.UpdateToken(data.TokenInfo)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	return true, nil
}

func (data *User) GetToken() (bool, error) {
	panic("implement me")
}

func (data *User) GetTokens() ([]userModels.TokenInfo, error) {
	panic("implement me")
}
