package database

import "time"

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

type TokenInfo struct {
	UserID          int
	TokenNum        string `xorm:"pk"`
	UpdateTokenNum  string
	StratTime       time.Time
	EndTime         time.Time
	UpdateTokenTime time.Time
}

type TokenInfoAudit struct {
	AuditID         int `xorm:"pk autoincr"`
	AuditType       int
	AuditTime       time.Time `xorm:"created"`
	UserID          int
	TokenNum        string
	UpdateTokenNum  string
	StratTime       time.Time
	EndTime         time.Time
	UpdateTokenTime time.Time
}
