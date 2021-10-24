package investmentWork

func SetInvestmentItem() (bool, error) {
	return RunTransaction(actionItem)
}

func SetInvestmentType() (bool, error) {
	return RunTransaction(actionInvestmentType)
}

func SetInvestmentServiceCharge() (bool, error) {
	return RunTransaction(actionInvestmentServiceCharge)
}
