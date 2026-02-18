package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jacobbananaldev/arcflow/internal/config"
)

func main() {
	r := chi.NewRouter()
	cfg := config.Load()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Arcflow API running"))
	})

	log.Printf("Server running on :%s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
