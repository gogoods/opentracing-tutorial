package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Do executes an HTTP request and returns the response body.
// Any errors or non-200 status code result in an error.
func Do(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, body)
	}

	return body, nil
}

func Post(url string, body []byte) ([]byte, error){

	buf := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}
	return Do(req)
}

func Get(url string)([]byte, error){
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return Do(req)
}