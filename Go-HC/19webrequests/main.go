package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilError(err)
	content := string(databytes)
	fmt.Println(content)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
