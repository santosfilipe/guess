package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var api = "https://countriesnow.space/api/v0.1/countries/capital"

type Response struct {
	Data  []Data `json:"data"`
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}

type Data struct {
	Capital string `json:"capital"`
	Name    string `json:"name"`
}

type Page struct {
	Title   string
	Country string
	Capital string
}

func connectToApi(api string) []byte {
	response, err := http.Get(api)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func generateRandomNumber(responseData []byte) int {
	var randomNumber int
	var responseObject Response

	rand.Seed(time.Now().UnixNano())
	json.Unmarshal(responseData, &responseObject)

	randomNumber = rand.Intn(len(responseObject.Data))

	return randomNumber
}

func printCapitalAndCountry(randomNumber int, responseData []byte) (string, string) {
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	var guessCountry = "Country: " + responseObject.Data[randomNumber].Name
	var guessCapital = "Capital: " + responseObject.Data[randomNumber].Capital

	return guessCountry, guessCapital
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	var apiResponse = connectToApi(api)
	var randomNumber = generateRandomNumber(apiResponse)

	guessCountry, guessCapital := printCapitalAndCountry(randomNumber, apiResponse)

	data := Page{
		Title:   "Guess",
		Country: guessCountry,
		Capital: guessCapital,
	}

	tmpl := template.Must(template.ParseFiles("guess.html"))
	tmpl.Execute(w, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/guess/", guessHandler)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
