//
// Copyright 2021 The Dapr Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createResponse() CoinDeskResponse {
	return CoinDeskResponse{
		Time:       CoinDeskResponseTime{Updated: "Apr 21, 2022 18:15:00 UTC", UpdatedISO: "2022-04-21T18:15:00+00:00", Updateduk: "Apr 21, 2022 at 19:15 BST"},
		Disclaimer: "This data was produced from the CoinDesk Bitcoin Price Index (USD). Non-USD currency data converted using hourly conversion rate from openexchangerates.org",
		ChartName:  "Bitcoin",
		Bpi: CoinDeskResponseBpi{
			Usd: CoinDeskResponseBpiContent{
				Code:        "USD",
				Symbol:      "&#36;",
				Rate:        "41,467.0555",
				Description: "United States Dollar",
				RateFloat:   41467.0555,
			},
			Gbp: CoinDeskResponseBpiContent{
				Code:        "GBP",
				Symbol:      "&pound;",
				Rate:        "31,701.9371",
				Description: "British Pound Sterling",
				RateFloat:   31701.9371,
			},
			Eur: CoinDeskResponseBpiContent{
				Code:        "EUR",
				Symbol:      "&euro;",
				Rate:        "38,034.5370",
				Description: "Euro",
				RateFloat:   38034.537,
			},
		},
	}
}

func TestGetResponseOk(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		validResponse := createResponse()
		validJson, _ := json.Marshal(validResponse)

		w.Write(validJson)
	}))
	defer ts.Close()

	response, err := getAndParseResponse(ts.URL)

	if err != nil {
		panic("err should be nil")
	}

	if response == nil {
		panic("err should not be nil")
	}
}
