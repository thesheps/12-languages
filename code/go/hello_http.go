package main

import (
	"net/http"
	"io/ioutil"
)

func main() {
	response, error := http.Get("http://www.google.com")

	if error != nil {
		// handle error
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		
		if err != nil {
			println(err)
		}

		println(string(contents))
	}
}