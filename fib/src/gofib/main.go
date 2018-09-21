package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FibonacciRequest struct {
	FibonacciNumber int `json:"fibonacci_number"`
}

type FibonacciResponse struct {
	FibonacciNumber int    `json:"fibonacci_number"`
	Value           uint64 `json:"value"`
}

type StatusResponse struct {
	CacheLength  int `json:"cache_length"`
	RequestCount int `json:"request_count"`
}

var call_count = 0
var fibonacci_numbers = []uint64{0, 1}

func main() {
	http.HandleFunc("/status", StatusHandler)
	http.HandleFunc("/fibonacci", FibonacciHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		sr := StatusResponse{len(fibonacci_numbers), call_count}
		jsonbytes, err := json.Marshal(sr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to process request"))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonbytes))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported Method"))
	}
}

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to read request body."))
		} else {
			call_count += 1
			fr := FibonacciRequest{}
			err = json.Unmarshal(body, &fr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Unable to unmarshal request JSON"))
			} else {
				if fr.FibonacciNumber > 92 {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("Sorry, I can only get to F(92)"))
				} else {
					result, err := calc_fibonacci(fr.FibonacciNumber)
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte(err.Error()))
					} else {
						fresp := FibonacciResponse{fr.FibonacciNumber, result}
						jsonbytes, err := json.Marshal(fresp)
						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
							w.Write([]byte("Unable to process request"))
						}
						w.Header().Set("Content-Type", "application/json")
						fmt.Fprintf(w, string(jsonbytes))
					}
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported Method."))
	}
}

func calc_fibonacci(number int) (uint64, error) {
	if number < 0 {
		return 0, errors.New("Negative, really?")
	}
	if number >= len(fibonacci_numbers) {
		result := uint64(0)
		for i := len(fibonacci_numbers); i <= number; i++ {
			result = fibonacci_numbers[i-2] + fibonacci_numbers[i-1]
			fibonacci_numbers = append(fibonacci_numbers, result)
		}
	}
	return fibonacci_numbers[number], nil
}
