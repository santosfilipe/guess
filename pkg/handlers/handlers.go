package handlers

import (
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/santosfilipe/guess/pkg/data"
)

type Page struct {
	Title   string
	Country string
	Capital string
}

func GeneratePseudoRandomIndex(responseData *data.Response) int {
	var randomNumber int

	rand.Seed(time.Now().UnixNano())

	randomNumber = rand.Intn(len(responseData.Data))

	return randomNumber
}

func printCapitalAndCountry(randomNumber int, responseData *data.Response) (string, string) {
	var guessCountry = "Country: " + responseData.Data[randomNumber].Name
	var guessCapital = "Capital: " + responseData.Data[randomNumber].Capital

	return guessCountry, guessCapital
}

func GuessHandler(w http.ResponseWriter, r *http.Request) {
	gc := data.New()
	geoData, err := gc.RetrieveGeoData()
	if err != nil {
		log.Fatal(err)
	}

	randomNumber := GeneratePseudoRandomIndex(geoData)
	guessCountry, guessCapital := printCapitalAndCountry(randomNumber, geoData)

	data := Page{
		Title:   "Guess",
		Country: guessCountry,
		Capital: guessCapital,
	}

	tmpl := template.Must(template.ParseFiles("guess.html"))
	tmpl.Execute(w, data)
}
