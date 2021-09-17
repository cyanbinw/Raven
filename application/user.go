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
	modelUser := newDBUserInfo(*data.UserInfo)
	flag, err := database.Login(&modelUser)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return flag, err
	}

	if flag {
		modelToken := new(database.TokenInfo)
		modelToken.UserID = modelUser.ID
		err := database.CreateToken(modelToken)
		if err != nil {
			Helheim.Writer(Helheim.Error, err)
			return false, err
		}
		toDTOToken(*modelToken, data.TokenInfo)
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
	var model database.TokenInfo
	toDBToken(*data.TokenInfo, &model)
	err := database.CreateToken(&model)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	toDTOToken(model, data.TokenInfo)
	return false, nil
}

func (data *User) ValidateToken() (bool, error) {
	var model database.TokenInfo
	toDBToken(*data.TokenInfo, &model)

	flag, err := database.ValidateToken(model.TokenNum)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}

	toDTOToken(model, data.TokenInfo)
	return flag, nil
}

func (data *User) UpdateToken() (bool, error) {
	var model database.TokenInfo
	toDBToken(*data.TokenInfo, &model)
	err := database.UpdateToken(&model)
	if err != nil {
		Helheim.Writer(Helheim.Error, err)
		return false, err
	}
	toDTOToken(model, data.TokenInfo)
	return true, nil
}

func (data *User) GetToken() (bool, error) {
	panic("implement me")
}

func (data *User) GetTokens() ([]userModels.TokenInfo, error) {
	panic("implement me")
}

func NewUser() User {
	user := User{}
	user.UserInitDB()
	return user
}

func newDBUserInfo(data userModels.UserInfo) database.UserInfo {
	return database.UserInfo{
		ID:       data.ID,
		UserName: data.UserName,
		Password: data.Password,
		Status:   data.Status,
	}
}

func toDBToken(dto userModels.TokenInfo, model *database.TokenInfo) {
	*model = database.TokenInfo{
		UserID:          dto.UserID,
		TokenNum:        dto.TokenNum,
		UpdateTokenNum:  dto.UpdateTokenNum,
		StratTime:       dto.StratTime,
		EndTime:         dto.EndTime,
		UpdateTokenTime: dto.UpdateTokenTime,
	}
}

func toDTOToken(model database.TokenInfo, dto *userModels.TokenInfo) {
	*dto = userModels.TokenInfo{
		UserID:          model.UserID,
		TokenNum:        model.TokenNum,
		UpdateTokenNum:  model.UpdateTokenNum,
		StratTime:       model.StratTime,
		EndTime:         model.EndTime,
		UpdateTokenTime: model.UpdateTokenTime,
	}
}
