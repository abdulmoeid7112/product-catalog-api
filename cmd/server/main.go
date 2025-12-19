package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/mytheresa/go-hiring-challenge/app/catalog"
	category "github.com/mytheresa/go-hiring-challenge/app/category"
	"github.com/mytheresa/go-hiring-challenge/app/database"
	"github.com/mytheresa/go-hiring-challenge/repositories"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// signal handling for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initialize database connection
	db, close := database.New(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	defer close()

	// Initialize handlers
	prodRepo := repositories.NewGormProductRepository(db)
	cat := catalog.NewCatalogHandler(prodRepo)

	catRepo := repositories.NewGormCategoryRepository(db)
	catHandler := category.NewCategoryHandler(catRepo)

	// Use Gorilla Mux
	r := mux.NewRouter()

	// Define routes
	// Catalog routes
	r.HandleFunc("/catalog", cat.HandleList).Methods("GET")
	r.HandleFunc("/catalog/{code}", cat.HandleDetail).Methods("GET")

	// Category routes
	r.HandleFunc("/categories", catHandler.HandleList).Methods("GET")
	r.HandleFunc("/categories", catHandler.HandleCreate).Methods("POST")

	// Set up the HTTP server
	srv := &http.Server{
		Addr:    ":" + os.Getenv("HTTP_PORT"),
		Handler: r,
	}

	// Start the server
	go func() {
		log.Printf("Starting server on http://%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %s", err)
		}

		log.Println("Server stopped gracefully")
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
	stop()
}
