package userModels

import (
	"time"
)

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
	TokenNum        string `xorm:"pk"`
	UpdateTokenNum  string
	StratTime       time.Time
	EndTime         time.Time
	UpdateTokenTime time.Time
}

type IToken interface {
	GetToken() (bool, error)
	GetTokens() ([]TokenInfo, error)
	ValidateToken() (bool, error)
	UpdateToken() (bool, error)
	CreateToken() (bool, error)
}

const (
	EffectiveTime       = 7  //天
	UpdateEffectiveTime = 30 //天
)
