package main

import (
	"log"
	"net/http"
)

// Enquanto o cenário do tratamento ao método HTTP POST nos deixa mais perto do

type StoragePlayerInMemory struct {}

func (s *StoragePlayerInMemory) GetScoreByPlayer(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&StoragePlayerInMemory{}}
	
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("não foi possível escutar na porta 5000 %v", err)
	}
}