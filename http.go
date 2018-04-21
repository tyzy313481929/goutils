package utils

import (
	"bytes"
	"encoding/json"
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

func POSTJSon(url string, obj interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(obj)
	fmt.Println(jsonData)
	if err != nil {
		return nil, err
	}
	jsonData = bytes.Replace(jsonData, []byte("\\u003c"), []byte("<"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u003e"), []byte(">"), -1)
	jsonData = bytes.Replace(jsonData, []byte("\\u0026"), []byte("&"), -1)
	body := bytes.NewBuffer(jsonData)
	response, err := http.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : uri=%v , statusCode=%v", url, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
