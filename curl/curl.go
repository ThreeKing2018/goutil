package curl

import (
	"net/http"
	"io/ioutil"
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