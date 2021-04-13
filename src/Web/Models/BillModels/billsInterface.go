package BillModels

type IBillData interface {
	NewBillData()
	BillsInitDB()
	BillsWriteToJSON()
	BillsGetYearData()
}
