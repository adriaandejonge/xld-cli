package http

import (
	"github.com/adriaandejonge/xld/util/login"
	"io"
	"io/ioutil"
	"net/http"
	"errors"
	"fmt"
)

func Read(path string) (body []byte, err error) {
	statusCode, body, err := Get(path)
	err = checkStatusCode(statusCode, body, err)
	return 
}

func Create(path string, reader io.Reader) (body []byte, err error) {
	statusCode, body, err := Post(path, reader)
	err = checkStatusCode(statusCode, body, err)
	return
}

func Update(path string, reader io.Reader) (body []byte, err error) {
	statusCode, body, err := Put(path, reader)
	err = checkStatusCode(statusCode, body, err)
	return 
}

func Remove(path string) (body []byte, err error) {

	// TODO Check for existance first
	
	statusCode, body, err := Delete(path)
	err = checkStatusCode(statusCode, body, err)
	return 
}

func Get(path string) (statusCode int, body []byte, err error) {
	return doHttp(path, "GET", nil)
}

func Post(path string, reader io.Reader) (statusCode int, body []byte, err error) {
	return doHttp(path, "POST", reader)
}

func Put(path string, reader io.Reader) (statusCode int, body []byte, err error) {
	return doHttp(path, "PUT", reader)
}

func Delete(path string) (statusCode int, body []byte, err error) {
	return doHttp(path, "DELETE", nil)
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

	// TODO Make content type variable in HTTP util
	req.Header.Add("Content-Type", "application/xml")
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

func checkStatusCode(statusCode int, body []byte, err error) error {
	if  err != nil {
		return err
	} else if statusCode < 200 || statusCode >= 300 {
		return errors.New(fmt.Sprintf("HTTP status code %d: %s", statusCode, body))
		
		// TODO if message type is XML (validation-message), then read and display nicely
	} 
	return nil
}