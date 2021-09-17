package database

type Investmentitem struct {
	Itemid int    `xorm:"INT"`
	Name   string `xorm:"VARCHAR(255)"`
	Code   string `xorm:"VARCHAR(255)"`
}
