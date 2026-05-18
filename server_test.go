package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PlayerStorageOutline struct {
	scores map[string]int
}

func (p *PlayerStorageOutline) GetScoreByPlayer(name string) int {
	scores := p.scores[name]
	return scores
}

func TestGetPlayer(t *testing.T) {
	storagePlayer := PlayerStorageOutline{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		},
	}

	server := &PlayerServer{&storagePlayer}

	t.Run("retornar resultado de Maria", func(t *testing.T) {
		request := newRequestGetScore("Maria")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseCodeStatus(t, response.Code, http.StatusOK)
		checkRequestBody(t, response.Body.String(), "20")
	})

	t.Run("retornar resultado de Pedro", func(t *testing.T) {
		request := newRequestGetScore("Pedro")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseCodeStatus(t, response.Code, http.StatusOK)
		checkRequestBody(t, response.Body.String(), "10")
	})

	t.Run("retorna 404 para jogador não encontrado", func(t *testing.T) {
		request := newRequestGetScore("Jorge")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		received := response.Code
		expected := http.StatusNotFound

		if received != expected {
			t.Errorf("received status %d expected %d", received, expected)
		}
	})
}

func newRequestGetScore(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func checkRequestBody(t *testing.T, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("corpo da requisição é inválido, obtive '%s' esperava '%s'", received, expected)
	}
}


func checkResponseCodeStatus(t *testing.T, received, expected int) {
	t.Helper()
    if received != expected {
        t.Errorf("não recebeu código de status HTTP esperado, received %d, esperado %d", received, expected)
    }
}