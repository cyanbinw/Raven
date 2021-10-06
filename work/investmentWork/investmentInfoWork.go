package investmentWork

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type InvestmentInfo struct {
	gid            string  /*股票编号*/
	increPer       float64 /*涨跌百分比*/
	increase       float64 /*涨跌额*/
	name           string  /*股票名称*/
	todayStartPri  float64 /*今日开盘价*/
	yestodEndPri   float64 /*昨日收盘价*/
	nowPri         float64 /*当前价格*/
	todayMax       float64 /*今日最高价*/
	todayMin       float64 /*今日最低价*/
	competitivePri float64 /*竞买价*/
	reservePri     float64 /*竞卖价*/
	traNumber      int64   /*成交量*/
	traAmount      int64   /*成交金额*/
	buyOne         int     /*买一*/
	buyOnePri      float64 /*买一报价*/
	buyTwo         int     /*买二*/
	buyTwoPri      float64 /*买二报价*/
	buyThree       int     /*买三*/
	buyThreePri    float64 /*买三报价*/
	buyFour        int     /*买四*/
	buyFourPri     float64 /*买四报价*/
	buyFive        int     /*买五*/
	buyFivePri     float64 /*买五报价*/
	sellOne        int     /*卖一*/
	sellOnePri     float64 /*卖一报价*/
	sellTwo        int     /*卖二*/
	sellTwoPri     float64 /*卖二报价*/
	sellThree      int     /*卖三*/
	sellThreePri   float64 /*卖三报价*/
	sellFour       int     /*卖四*/
	sellFourPri    float64 /*卖四报价*/
	sellFive       int     /*卖五*/
	sellFivePri    float64 /*卖五报价*/
	date           string  /*日期*/
	time           string  /*时间*/
}

type Dapandata struct {
	Dot       float64 `json:"dot"`
	Name      string  `json:"name"`
	NowPic    float64 `json:"nowPic"`
	Rate      float64 `json:"rate"`
	TraAmount float64 `json:"traAmount"`
	TraNumber float64 `json:"traNumber"`
}

func GetInvestmentInfo() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://web.juhe.cn:8080/finance/stock/hs", nil)
	

	response,_ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
	//    reqest, _ = http.NewRequest("POST","http:/127.0.0.1/", bytes.NewBufferString(data.Encode()))
	//    respet1,_ := http.NewRequest("POST","http://127.0.0.1/",url.Values{"key":"Value"})
	//    reqest1.Header.Set("User-Agent","chrome 100")
	//    client.Do(reqest1)
}