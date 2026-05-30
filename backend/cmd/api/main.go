package main

import (
	"MetalCalc/backend/internal/handlers"
	"MetalCalc/backend/internal/storage"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	storage.Seed()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	r.Post("/api/v1/forging", handlers.ForgingHandler)
	r.Post("/api/v1/forging/pdf", handlers.ForgingPDFHandler)

	r.Get("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	r.Get("/api/v1/materials", handlers.GetMaterialsHandler)
	r.Post("/api/v1/materials", handlers.AddMaterialHandler)
	r.Delete("/api/v1/materials/{id}", handlers.DeleteMaterialHandler)

	r.Get("/api/v1/shapes", handlers.GetShapesHandler)
	r.Post("/api/v1/shapes", handlers.AddShapeHandler)
	r.Delete("/api/v1/shapes/{id}", handlers.DeleteShapeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
