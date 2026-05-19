package main

import (
	"fmt"
	"net/http"
)

type StoragePlayer interface {
	GetScoreByPlayer(name string) int
	RecordsVictory(name string)
}

type PlayerServer struct {
	storage StoragePlayer
}

func (s *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.recordVictory(w, r)
	case http.MethodGet:
		s.showScore(w, r)
	}
}

func (s *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := s.getPlayerName(r)

	scores := s.storage.GetScoreByPlayer(player)

	if scores == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, scores)
}

func (s *PlayerServer) recordVictory(w http.ResponseWriter, r *http.Request) {
	player := s.getPlayerName(r)
	s.storage.RecordsVictory(player)
	w.WriteHeader(http.StatusAccepted)
}

func (s *PlayerServer) getPlayerName(r *http.Request) string {
	return r.URL.Path[len("/players/"):]
}
