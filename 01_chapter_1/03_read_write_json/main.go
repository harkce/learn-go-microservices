package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloHandler)

	log.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := helloResponse{"Hello world!"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Whops")
	}

	fmt.Fprintf(w, string(data))
}
