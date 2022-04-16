package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type StateData struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func main() {
	http.HandleFunc("/increment",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				log.Printf("Error: Got method %s instead of POST\n", r.Method)
				w.WriteHeader(400)
				return
			}

			resp, _ := http.Get("http://localhost:8089/v1.0/state/statestore/counter")
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
				"http://localhost:8089/v1.0/state/statestore",
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

			resp, _ := http.Get("http://localhost:8089/v1.0/state/statestore/counter")
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			strVal := string(body)
			log.Printf("Current counter is: %s\n", strVal)
			w.WriteHeader(200)
			fmt.Fprintf(w, strVal)
		})

	log.Fatal(http.ListenAndServe(":8088", nil))
}
