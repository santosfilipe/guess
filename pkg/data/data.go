package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Data  []Data `json:"data"`
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}

type Data struct {
	Capital string `json:"capital"`
	Name    string `json:"name"`
}

func ConnectToApi(api string) *Response {
	var responseObject Response

	response, err := http.Get(api)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &responseObject)

	return &responseObject
}
