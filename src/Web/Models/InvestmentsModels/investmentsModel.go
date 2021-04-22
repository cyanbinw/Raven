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
	investmentsInitDB()
}

func (data *InvestmentData) InvestmentGetAll() {
	investmentGetAll(data)
}

func (data *InvestmentData) SetInvestmentChartForAccount() InvestmentsChartModel {
	return investmentGetDataToChart()
}

func (data *InvestmentData) GetInvestmentTable() []InvestmentTable {
	return investmentGetTable()
}

func (data *InvestmentTable) InvestmentsInitDB() {
	investmentsInitDB()
}

func (data *InvestmentTable) AddInvestmentTable() (bool, error) {
	return investmentAddTable(*data)
}

func (data *InvestmentTable) UpdateInvestmentTable() (bool, error) {
	return investmentUpdateTable(*data)
}

func GetInvestmentDiagram() (map[string][]Investment, error) {
	investmentsInitDB()
	return investmentGetDiagram()
}

func GetInvestmentOption() ([]InvestmentType, []InvestmentActivity, error) {
	investmentsInitDB()
	return investmentGetOption()
}
