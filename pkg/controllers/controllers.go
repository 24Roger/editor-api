package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/24Roger/editor-api/pkg/database"
	"github.com/24Roger/editor-api/pkg/query"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type Project struct {
	Uid  string `json:"uid"`
	Data string `json:"data"`
}

func FindAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dgClient := database.NewClient()
	txn := dgClient.NewTxn()
	res, err := txn.Query(context.Background(), query.GetAllProjects)

	if err != nil {
		log.Println(err)
	}
	w.Write(res.Json)
}

func NewProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Project
	_ = json.NewDecoder(r.Body).Decode(&req)

	p := Project{Data: req.Data}
	pb, err := json.Marshal(p)

	if err != nil {
		log.Println(err)
	}

	dgClient := database.NewClient()
	txn := dgClient.NewTxn()

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	res, err := txn.Mutate(context.Background(), mu)

	if err != nil {
		log.Println(err)
	}
	w.Write(res.Json)
}
