package database

import (
	"time"
)

type Investmentactivity struct {
	Activityid   int       `xorm:"not null pk autoincr INT"`
	Activityname string    `xorm:"not null default '' VARCHAR(100)"`
	Insertdate   time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
