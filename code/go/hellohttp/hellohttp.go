package hellohttp

import (
	"net/http"
	"io/ioutil"
)

// Hello grabs the string Hello World over the wire
func Hello() string {
	response, error := http.Get("https://thesheps.github.io/12-languages/hello-world.html")

	if error != nil {
		// handle error
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		
		if err != nil {
			println(err)
		}

	    return string(contents)
	}

	return ""
}