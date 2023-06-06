package main

import (

	"log"
	"net/http"

	"github.com/Ozan-sert/receipt-processor-challenge/middlewares"
	"github.com/Ozan-sert/receipt-processor-challenge/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func main() {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middlewares.RecoverMiddleware)
	router.Use(cors.Default().Handler)

	// Custom 404 Not Found handler
	router.NotFound(middlewares.NotFound)

	// Routes
	router.Mount("/receipts", routes.ReciptRoutes())

	

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
