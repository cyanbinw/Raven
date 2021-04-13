package InvestmentsModels

type InvestmentData struct {
	Data []Investment
}

type InvestmentAccountModel struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type InvestmentChartModel struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type InvestmentsChartModel struct {
	Account  []InvestmentChartModel
	Share    []InvestmentChartModel
	NetWorth []InvestmentChartModel
}

func (data *InvestmentData) InvestmentsInitDB() {
	InvestmentsInitDB()
}

func (data *InvestmentData) InvestmentGetAll() {
	InvestmentGetAll(data)
}

func (data *InvestmentData) SetInvestmentChartForAccount() InvestmentsChartModel {
	return InvestmentGetDataToChart()
}

func (data *InvestmentData) GetInvestmentTable() []InvestmentTable {
	return InvestmentGetTable()
}

func (data *InvestmentTable) InvestmentsInitDB() {
	InvestmentsInitDB()
}

func (data *InvestmentTable) AddInvestmentTable() (bool, error) {
	return InvestmentAddTable(*data)
}

func (data *InvestmentTable) UpdateInvestmentTable() (bool, error) {
	return InvestmentUpdateTable(*data)
}

func GetInvestmentDiagram() (map[string][]Investment, error) {
	InvestmentsInitDB()
	return InvestmentGetDiagram()
}

func GetInvestmentOption() ([]InvestmentType, []InvestmentActivity, error) {
	InvestmentsInitDB()
	return InvestmentGetOption()
}
