package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type StateData struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func main() {
	daprPort := os.Getenv("DAPR_HTTP_PORT")
	stateStoreUrl := "http://localhost:" + daprPort + "/v1.0/state/statestore"

	log.Println(stateStoreUrl)

	http.HandleFunc("/increment",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				log.Printf("Error: Got method %s instead of POST\n", r.Method)
				w.WriteHeader(400)
				return
			}

			resp, _ := http.Get(stateStoreUrl + "/counter")
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			strVal := string(body)
			count := 1
			if strVal != "" {
				count, _ = strconv.Atoi(strVal)
				count++
			}

			log.Printf("Current counter is: %s, now incremented to: %d\n", strVal, count)
			stateObj := []StateData{{Key: "counter", Value: count}}
			stateData, _ := json.Marshal(stateObj)
			_, _ = http.Post(
				stateStoreUrl,
				"application/json",
				bytes.NewBuffer(stateData),
			)
		})

	http.HandleFunc("/counter",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" {
				log.Printf("Error: Got method %s instead of GET\n", r.Method)
				w.WriteHeader(400)
				return
			}

			resp, _ := http.Get(stateStoreUrl + "/counter")
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			strVal := string(body)
			log.Printf("Current counter is: %s\n", strVal)
			w.WriteHeader(200)
			fmt.Fprintf(w, strVal)
		})

	log.Printf("Starting counter-app...")
	log.Fatal(http.ListenAndServe(":8088", nil))
}
