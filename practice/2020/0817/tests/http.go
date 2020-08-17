package tests

import (
	"io/ioutil"
	"log"
	"net/http"
)

func DoRequest(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(data))
	return nil
}
