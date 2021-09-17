package database

import (
	"time"
)

type Liabilities struct {
	Id                int       `xorm:"not null pk autoincr INT"`
	Target            int       `xorm:"not null INT"`
	Liabilitiesname   string    `xorm:"not null VARCHAR(500)"`
	Liabilitiesamount string    `xorm:"not null DECIMAL(10)"`
	Notes             string    `xorm:"not null VARCHAR(500)"`
	Installdate       time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Updatedate        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
