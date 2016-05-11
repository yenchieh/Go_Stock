package common

import (
	"net/url"
	"log"
	"io/ioutil"
	"net/http"
)


func RequestService(Url *url.URL) ([]byte, error) {
	log.Printf("Request URL: %s\n", Url.String())
	resp, err := http.Get(Url.String())

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Resposne: %s", body)

	return body, nil

}
