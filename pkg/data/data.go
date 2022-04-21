package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

type geoApiClient struct {
	httpClient *http.Client
	urlPrefix  string
}

func New() *geoApiClient {
	gc := geoApiClient{
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
		},
		urlPrefix: "https://countriesnow.space/api/v0.1/countries/capital",
	}
	return &gc
}

func (gc *geoApiClient) newRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", gc.urlPrefix, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (gc *geoApiClient) RetrieveGeoData() (*Response, error) {
	var responseObject Response

	req, err := gc.newRequest()
	if err != nil {
		return nil, err
	}

	res, err := gc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	json.Unmarshal(responseData, &responseObject)

	return &responseObject, nil
}
