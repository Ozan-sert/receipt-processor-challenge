// routes/routes.go

package routes

import (
	"github.com/go-chi/chi"
	"github.com/Ozan-sert/receipt-processor-challenge/handlers"
)

// ReciptRoutes sets up the routes related to receipt processing
func ReciptRoutes() chi.Router {
	router := chi.NewRouter()

	// POST /receipts/process - Process receipts
	router.Post("/process", handlers.ProcessReceipts)

	// GET /receipts/{id}/points - Get points for a specific receipt
	router.Get("/{id}/points", handlers.GetPoints)

	return router
}
