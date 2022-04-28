package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		return nil, fmt.Errorf("failed to get JSON: %v", err)
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
	if len(sourceEndpoint) == 0 {
		log.Fatal("COIN_DESK_ENDPOINT must be set")
	}

	pubSubName := os.Getenv("PUB_SUB_NAME")
	if len(pubSubName) == 0 {
		log.Fatal("PUB_SUB_NAME must be set")
	}

	topicName := os.Getenv("TOPIC_NAME")
	if len(topicName) == 0 {
		log.Fatal("TOPIC_NAME must be set")
	}

	queryInterval, err := strconv.Atoi(os.Getenv("QUERY_INTERVAL_SECONDS"))
	if err != nil || queryInterval <= 0 {
		queryInterval = 10
		log.Printf("QUERY_INTERVAL_SECONDS not set, using default value: %d\n", queryInterval)
	}

	//client, err := dapr.NewClient()
	//if err != nil {
	//	panic(err)
	//}
	//
	//ctx := context.Background()

	client := http.Client{}

	ticker := time.Tick(time.Duration(queryInterval) * time.Second)
	for next := range ticker {
		result, err := getAndParseResponse(sourceEndpoint)
		if err != nil {
			log.Printf("failed parsing response: %v", err)
		}

		log.Printf("%v %+v\n", next, result)

		priceEvent := PriceEvent{
			Symbol: result.Symbol,
			Usd:    result.Usd,
			Eur:    result.Eur,
		}

		priceEventJson, err := json.Marshal(priceEvent)

		//log.Printf("wat")
		//if client == nil {
		//	panic("client is null")
		//}
		//err = client.PublishEvent(ctx, pubSubName, topicName, priceEvent)
		//if err != nil {
		//	panic(err)
		//}

		req, err := http.NewRequest("POST", "http://localhost:"+daprPort+"/v1.0/publish/"+pubSubName+"/"+pubSubName, bytes.NewBuffer(priceEventJson))
		if err != nil {
			log.Fatal(err)
		}

		// Publish an event using Dapr pub/sub
		if _, err = client.Do(req); err != nil {
			log.Fatal(err)
		}

	}
}
