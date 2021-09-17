package database

import (
	"time"
)

type Investmenttype struct {
	Typeid     int       `xorm:"not null pk autoincr INT"`
	Typename   string    `xorm:"not null default '' VARCHAR(100)"`
	Insertdate time.Time `xorm:"default 'CURRENT_TIMESTAMP' DATETIME"`
}
