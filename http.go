package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HTTPGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get %=v,statuscode=%v", url, resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
