package database

import "time"

type Liabilities struct {
	ID                int       `xorm:"not null pk autoincr INT"`
	Target            int       `xorm:"not null INT"`
	LiabilitiesName   string    `xorm:"not null VARCHAR(500)"`
	LiabilitiesAmount string    `xorm:"not null DECIMAL(10)"`
	Notes             string    `xorm:"not null VARCHAR(500)"`
	InstallDate       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdateDate        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}