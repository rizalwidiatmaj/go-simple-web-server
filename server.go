package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			message := &Message{
				Status:  200,
				Message: "Hello World",
			}
			jsonMessage, _ := json.Marshal(message)
			w.Write(jsonMessage)
		}
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
