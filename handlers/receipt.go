package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/Ozan-sert/receipt-processor-challenge/models"
)

// In-memory storage for receipts
var receipts = make(map[string]models.Receipt)

// processReceipts handles the /receipts/process endpoint.
func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	receipts[id] = receipt

	response := struct {
		ID string `json:"id"`
	}{
		ID: id,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// getPoints handles the /receipts/{id}/points endpoint.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	//  lookup based on the provided ID
	receipt, found := receipts[id]
	if !found {
		http.NotFound(w, r)
		return
	}

	points := receipt.CalculatePoints()

	response := struct {
		Points int `json:"points"`
	}{
		Points: points,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


