package controllers

type ReturnData struct {
	Successful bool        `json:"successful"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
