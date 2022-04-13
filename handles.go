package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgraph-io/dgo/v210/protos/api"
)

type Project struct {
	Uid  string `json:"uid"`
	Data string `json:"data"`
}

func findAllProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dgClient := newClient()
	txn := dgClient.NewTxn()
	res, err := txn.Query(context.Background(), getAllProjects)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(res.Json)
}

func newProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req Project
	_ = json.NewDecoder(r.Body).Decode(&req)

	p := Project{Data: req.Data}
	pb, err := json.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	dgClient := newClient()
	txn := dgClient.NewTxn()

	mu := &api.Mutation{
		CommitNow: true,
		SetJson:   pb,
	}

	res, err := txn.Mutate(context.Background(), mu)

	if err != nil {
		log.Fatal(err)
	}
	w.Write(res.Json)
}
