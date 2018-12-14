package curl

import (
	"testing"
	"fmt"
)



func TestHttpGet(t *testing.T) {
	url := "http://www.sgfoot.com"
	resp, err := HttpGet(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
}
