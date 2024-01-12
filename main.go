package main

import (
	"C0lliNN/better-handlers-blog/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type TransactionRequest struct {
	CardNumber string `json:"card_number"`
	CVV string `json:"cvv"`
	ExpirationDate string `json:"expiration_date"`
	Amount int `json:"amount"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req TransactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		card, err := repository.FindCardByNumber(req.CardNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if card.CVV != req.CVV || card.ExpirationDate != req.ExpirationDate  {
			http.Error(w, "invalid cvv or exp date", http.StatusUnauthorized)
			return
		}

		if card.AvaliableLimit < req.Amount {
			http.Error(w, "insufficient funds", http.StatusUnauthorized)
			return
		}

		card.AvaliableLimit -= req.Amount
		if err := repository.SaveCard(card); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "transaction completed")
	})

	http.ListenAndServe(":8080", nil)
}

