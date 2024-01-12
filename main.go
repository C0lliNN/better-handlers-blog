package main

import (
	"C0lliNN/better-handlers-blog/processor"
	"C0lliNN/better-handlers-blog/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	repository := repository.NewCardRepository()
	cardProcessor := processor.NewProcessor(repository)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req processor.TransactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := cardProcessor.ProcessTransaction(req); err != nil {
			// map errors the desired status code
			http.Error(w, err.Error(), http.StatusForbidden)
		}

		fmt.Fprintf(w, "transaction completed")
	})

	http.ListenAndServe(":8080", nil)
}

