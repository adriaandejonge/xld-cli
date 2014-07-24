package http

import (
	"github.com/adriaandejonge/xld/login"
	"io"
	"io/ioutil"
	"net/http"
)

func Read(path string) (statusCode int, body []byte, err error) {

	return doHttp(path, "GET", nil)
}

func Post(path string, reader io.Reader) (statusCode int, body []byte, err error) {

	return doHttp(path, "POST", reader)
}

func doHttp(path string, method string, reader io.Reader) (statusCode int, body []byte, err error) {

	client := &http.Client{}

	loginData, err := login.Check()
	if err != nil {
		return
	}

	url := "http://" + loginData.Url + "/deployit" + path

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Basic "+loginData.Auth)
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return resp.StatusCode, body, err

}