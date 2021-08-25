package userModels

import (
	"time"
)

type UserInfo struct {
	ID       int `xorm:"pk autoincr"`
	UserName string
	Password string
	Status   int
}

type UserInfoAudit struct {
	AuditID    int `xorm:"pk autoincr"`
	AuditType  int
	AuditTime  time.Time `xorm:"created"`
	UserID     int
	UserName   string
	Password   string
	InsertTime time.Time `xorm:"updated"` //修改后自动更新时间
	UpdateTime time.Time `xorm:"created"` //创建时间
	Status     int
}

type IUserInfo interface {
	Login() (bool, error)
	Register() (bool, error)
	GetUser() (bool, error)
	GetUsers() ([]UserInfo, error)
	SetUserInfo() (bool, error)
	Delete() (bool, error)
}
