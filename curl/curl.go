package curl

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func HttpGet(url string) (reply []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	reply, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
//在线验证银行卡信息
func ValidateOnlineBankCode(bankCard string) (b bool) {
	type ValidateBankCard struct {
		CardType  string        `json:"cardType"`
		Bank      string        `json:"bank"`
		Key       string        `json:"key"`
		Messages  []interface{} `json:"messages"`
		Validated bool          `json:"validated"`
		Stat      string        `json:"stat"`
	}
	url := "https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardNo=%s&cardBinCheck=true"
	url = fmt.Sprintf(url, bankCard)
	resp, err := HttpGet(url)
	if err != nil {
		return
	}
	result := ValidateBankCard{}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.Validated == true && result.Bank != "" {
		return true
	}
	return
}