package dummydata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const response = `{"data":123, "data1": 22}`

type httpBinIpResponse struct {
	Origin string
}

type Requester func(string) []byte

func GetResponse() string {
	return response
}

func ZendeskRequest() {

}

func GetRequester(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request error %s \n", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	return body
}

func SimpleGet(requester Requester, url string) *httpBinIpResponse {
	body := requester(url)
	respObj := httpBinIpResponse{}
	if err := json.Unmarshal(body, &respObj); err != nil {
		panic(err)
	}
	return &respObj
}
