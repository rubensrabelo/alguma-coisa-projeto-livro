package main

import (
	"log"
	"net/http"
)

// JSON, roteamento e aninhamento

type InMemoryPlayerStorage struct {
	scores        map[string]int
	recordVictory []string
}

func (i *InMemoryPlayerStorage) GetScoreByPlayer(name string) int {
	return i.scores[name]
}

func (i *InMemoryPlayerStorage) RecordsVictory(name string) {
	i.scores[name]++
}

func main() {
	storage := &InMemoryPlayerStorage{
		scores: map[string]int{},
	}
	server := &PlayerServer{storage: storage}

	log.Println("Servidor rodando na porta :5000...")
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
