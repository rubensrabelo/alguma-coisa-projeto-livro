package main

import (
	"fmt"
	"net/http"
)

type StoragePlayer interface {
	GetScoreByPlayer(name string) int
}

type PlayerServer struct {
	storage StoragePlayer
}


func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	fmt.Fprint(w, s.storage.GetScoreByPlayer(player))
}