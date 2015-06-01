package main

import (
	"fmt"
	"github.com/jeffail/gabs"
	"io/ioutil"
	"net/http"
)

const URLPrefix = "https://api.uwaterloo.ca/v2/"

func callAPI(key, url string) (*gabs.Container, error) {
	var empty *gabs.Container

	// send the get request to the UW API...
	res, err := http.Get(url)
	if err != nil {
		return empty, err
	}
	defer res.Body.Close()

	// read the response (it's a byte array)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return empty, err
	}

	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return empty, err
	}

	return jsonParsed, nil

}

func main() {
	fmt.Println("Boop")
}
