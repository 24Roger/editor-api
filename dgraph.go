package main

import (
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/joho/godotenv"
)

func newClient() *dgo.Dgraph {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DATABASE_URL := os.Getenv("DATABASE_URL")
	API_KEY := os.Getenv("API_KEY")

	conn, err := dgo.DialCloud(DATABASE_URL, API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}