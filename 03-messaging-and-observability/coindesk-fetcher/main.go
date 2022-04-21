package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type CoinDeskResponseTime struct {
	Updated    string `json:"updated"`
	UpdatedISO string `json:"updatedISO"`
	Updateduk  string `json:"updateduk"`
}

type CoinDeskResponseBpiContent struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rate_float"`
}

type CoinDeskResponseBpi struct {
	Usd CoinDeskResponseBpiContent `json:"USD"`
	Gbp CoinDeskResponseBpiContent `json:"GBP"`
	Eur CoinDeskResponseBpiContent `json:"EUR"`
}

type CoinDeskResponse struct {
	Time       CoinDeskResponseTime `json:"time"`
	Disclaimer string               `json:"disclaimer"`
	ChartName  string               `json:"chartName"`
	Bpi        CoinDeskResponseBpi  `json:"bpi"`
}

type PriceEvent struct {
	Symbol string  `json:"symbol"`
	Usd    float64 `json:"usd"`
	Eur    float64 `json:"eur"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func getAndParseResponse(url string) (*PriceEvent, error) {
	response := new(CoinDeskResponse)

	err := getJson(url, response)
	if err != nil {
		return nil, errors.New("failed to get JSON")
	}

	return &PriceEvent{
		Symbol: "BTC",
		Usd:    response.Bpi.Usd.RateFloat,
		Eur:    response.Bpi.Eur.RateFloat,
	}, nil
}

func main() {
	daprPort := os.Getenv("DAPR_HTTP_PORT")
	if len(daprPort) == 0 {
		log.Fatal("DAPR_HTTP_PORT must be set")
	}

	sourceEndpoint := os.Getenv("COIN_DESK_ENDPOINT")
	if len(daprPort) == 0 {
		log.Fatal("COIN_DESK_ENDPOINT must be set")
	}

	queryInterval, err := strconv.Atoi(os.Getenv("QUERY_INTERVAL_SECONDS"))
	if err != nil || queryInterval <= 0 {
		queryInterval = 10
		log.Printf("QUERY_INTERVAL_SECONDS not set, using default value: %d\n", queryInterval)
	}

	ticker := time.Tick(time.Duration(queryInterval) * time.Second)
	for next := range ticker {
		result, err := getAndParseResponse(sourceEndpoint)
		if err != nil {
			log.Printf("failed parsing response: %e", err)
		}

		log.Printf("%v %+v\n", next, result)
	}
}
