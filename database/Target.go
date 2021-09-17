package database

import "time"

type Target struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	Yeartarget5 string    `xorm:"5YearTarget not null VARCHAR(500)"`
	Revenue     string    `xorm:"not null DECIMAL(10)"`
	Expend      string    `xorm:"not null DECIMAL(10)"`
	Skill       string    `xorm:"not null VARCHAR(500)"`
	Other       string    `xorm:"VARCHAR(500)"`
	Installdate time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Updatedate  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
