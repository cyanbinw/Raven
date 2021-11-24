package investmentsModels

import "time"

type Investment struct {
	ID             int64     `db:"ID" json:"id"`
	ItemID         int       `db:"ItemID" json:"itemID"`
	Code           string    `db:"Code" json:"code"`
	Name           string    `db:"Name" json:"name"`
	TypeID         int       `db:"TypeID" json:"type"`
	Account        float64   `db:"Account" json:"account"`
	Share          float64   `db:"Share" json:"share"`
	NetWorth       float64   `db:"NetWorth" json:"netWorth"`
	ServiceCharge  float64   `db:"ServiceCharge" json:"serviceCharge"`
	Date           time.Time `db:"Date" json:"date"`
	ActivityStatus int       `db:"ActivityStatus" json:"activity"`
	IsEmpty        bool      `db:"IsEmpty" json:"isEmpty"`
}

type InvestmentTable struct {
	Investment        `xorm:"extends"`
	ServiceChargeList []InvestmentServiceCharge `json:"serviceChargeList"`
	ActivityName      string                    `json:"activityName"`
	TypeName          string                    `json:"typeName"`
}

type InvestmentActivity struct {
	ActivityID   int
	ActivityName string
	InsertDate   time.Time
}

type InvestmentItem struct {
	ItemID int
	Name   string
	Code   string
}

type InvestmentType struct {
	TypeID     int `xorm:"pk autoincr"`
	TypeName   string
	GroupID    int
	GroupName  string
	InsertDate time.Time `xorm:"created"`
}

type InvestmentServiceCharge struct {
	ItemID     int64     `xorm:"pk" json:"itemID"`
	TypeID     int       `xorm:"pk" json:"typeID"`
	Cost       float64   `json:"cost"`
	InsertDate time.Time `xorm:"created"`
	UpdateDate time.Time `xorm:"updated"`
}

type InvestmentServiceChargeType struct {
	TypeID     int `xorm:"pk autoincr"`
	TypeName   string
	InsertDate time.Time `xorm:"created"`
	UpdateDate time.Time `xorm:"updated"`
}

type InvestmentChartModel struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (InvestmentTable) TableName() string {
	return "Investment"
}

type InvestmentReportForm struct {
	Position               int     //持仓
	PositionInvement       float64 //持仓投资
	TotalPositionInvement  float64 //持仓总投资
	PositionSell           float64 //持仓卖出
	PositionServiceCharge  float64 // 持仓手续费
	PositionInformation    string  // 持仓信息
	Clearance              int     //清仓
	ClearanceInvement      float64 //清仓投资
	TotalClearanceInvement float64 //清仓总投资
	ClearanceSell          float64 //清仓卖出
	ClearanceServiceCharge float64 // 清仓手续费
	ClearanceInformation   string  // 清仓信息
	Total                  int     //总体
	TotalInvement          float64 //总体投资
	TotalTotalInvement     float64 //总体总投资
	TotalSell              float64 //总体卖出
	TotalServiceCharge     float64 // 总体手续费
	TotalInformation       string  // 总体信息
}
