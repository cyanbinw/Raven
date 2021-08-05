package userModels

import (
	"time"
)

type Token struct {
	UserID          int
	TokenNum        [16]byte `xorm:"pk"`
	UpdateTokenNum  [16]byte
	StratTime       time.Time
	EndTime         time.Time
	UpdateTokenTime time.Time
}

type TokenAudit struct {
	AuditID   int `xorm:"pk autoincr"`
	AuditType int
	AuditTime time.Time `xorm:"created"`
	Token
}

type IToken interface {
	GetToken() (bool, error)
	GetTokens() ([]Token, error)
	ValidateToken() (bool, error)
	UpdateToken() (bool, error)
	CreateToken() (bool, error)
}

const (
	EffectiveTime       = 7  //天
	UpdateEffectiveTime = 30 //天
)
