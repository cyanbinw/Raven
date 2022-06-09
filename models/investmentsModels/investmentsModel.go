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
	Position                 int     `json:"position"`                 //持仓
	PositionInvestment       float64 `json:"positionInvestment"`       //持仓投资
	TotalPositionInvestment  float64 `json:"totalPositionInvestment"`  //持仓总投资
	PositionSell             float64 `json:"positionSell"`             //持仓卖出
	PositionServiceCharge    float64 `json:"positionServiceCharge"`    // 持仓手续费
	PositionInformation      string  `json:"positionInformation"`      // 持仓信息
	Clearance                int     `json:"clearance"`                //清仓
	ClearanceInvestment      float64 `json:"clearanceInvestment"`      //清仓投资
	TotalClearanceInvestment float64 `json:"totalClearanceInvestment"` //清仓总投资
	ClearanceSell            float64 `json:"clearanceSell"`            //清仓卖出
	ClearanceServiceCharge   float64 `json:"clearanceServiceCharge"`   // 清仓手续费
	ClearanceInformation     string  `json:"clearanceInformation"`     // 清仓信息
	Total                    int     `json:"total"`                    //总体
	TotalInvestment          float64 `json:"totalInvestment"`          //总体投资
	TotalTotalInvestment     float64 `json:"totalTotalInvestment"`     //总体总投资
	TotalSell                float64 `json:"totalSell"`                //总体卖出
	TotalServiceCharge       float64 `json:"totalServiceCharge"`       // 总体手续费
	TotalInformation         string  `json:"totalInformation"`         // 总体信息
}
