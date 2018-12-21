package three_service

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeKing2018/goutil/curl"
)

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
	resp, err := curl.HttpGet(url)
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
