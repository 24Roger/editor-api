package database

import (
	"log"
	"os"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/joho/godotenv"
)

func NewClient() *dgo.Dgraph {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	API_KEY := os.Getenv("API_KEY")

	conn, err := dgo.DialCloud(DATABASE_URL, API_KEY)

	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(api.NewDgraphClient(conn))
}
