package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// A Response struct to map the Response
type Response struct {
	Quote string `json:"quote"`
}

func main() {
	// define the URL
	APIURL := "https://api.kanye.rest/"

	// store the response of a GET request in a variable called response
	response, err := http.Get(APIURL)

	// if the error contains something, exit
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// store the body of the GET request in resopnse to a variable called responseData
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Quote)

}
