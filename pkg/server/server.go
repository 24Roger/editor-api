package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type DrawFlowServer struct {
	server *http.Server
}

func Server(mux *chi.Mux) *DrawFlowServer {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &DrawFlowServer{s}
}

func (start *DrawFlowServer) Run() {
	fmt.Print("Server running on port " + start.server.Addr)
	log.Fatal(start.server.ListenAndServe())
}
