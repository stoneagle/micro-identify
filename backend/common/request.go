package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func DoHttpPost(url string, params []byte) (body []byte, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(params))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
