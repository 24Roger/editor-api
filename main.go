package main

import (
	"github.com/24Roger/editor-api/pkg/routes"
	"github.com/24Roger/editor-api/pkg/server"
)

func main() {
	mux := routes.Route()
	server := server.Server(mux)
	server.Run()
}
