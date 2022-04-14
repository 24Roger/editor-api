package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type DrawFlowServer struct {
	server *http.Server
}

func Server(mux *chi.Mux) *DrawFlowServer {
	s := &http.Server{
		Addr:           ":4000",
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
