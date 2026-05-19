package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PlayerStorageOutline struct {
	scores        map[string]int
	recordVictory []string
}

func (p *PlayerStorageOutline) GetScoreByPlayer(name string) int {
	scores := p.scores[name]
	return scores
}

func (p *PlayerStorageOutline) RecordsVictory(name string) {
	p.recordVictory = append(p.recordVictory, name)
}

func TestGetPlayer(t *testing.T) {
	// Dados compartilhados de leitura são aceitáveis aqui porque nenhum teste altera o mapa
	storagePlayer := PlayerStorageOutline{
		map[string]int{
			"Maria": 20,
			"Pedro": 10,
		},
		nil,
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

	t.Run("retorga 404 para jogador não encontrado", func(t *testing.T) {
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

func TestRecordsVictory(t *testing.T) {
	t.Run("retorna status 'aceito' para chamadas ao método POST", func(t *testing.T) {
		storage := PlayerStorageOutline{
			map[string]int{},
			nil,
		}
		server := &PlayerServer{&storage}

		request := newRequestRegisterVictoryPost("Maria")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		checkResponseCodeStatus(t, response.Code, http.StatusAccepted)

		if len(storage.recordVictory) != 1 {
			t.Errorf("verifiquei %d chamadas a RegistrarVitoria, esperava %d", len(storage.recordVictory), 1)
		}
	})

	t.Run("registra vitorias na chamada ao método HTTP POST", func(t *testing.T) {
		storage := PlayerStorageOutline{
			map[string]int{},
			nil,
		}
		server := &PlayerServer{&storage}
		player := "Maria"

		request := newRequestRegisterVictoryPost(player)
		resposta := httptest.NewRecorder()

		server.ServeHTTP(resposta, request)

		checkResponseCodeStatus(t, resposta.Code, http.StatusAccepted)

		if len(storage.recordVictory) != 1 {
			t.Errorf("verifiquei %d chamadas a RegistrarVitoria, esperava %d", len(storage.recordVictory), 1)
		}

		if storage.recordVictory[0] != player {
			t.Errorf("não registrou o vencedor corretamente, recebi '%s', esperava '%s'", storage.recordVictory[0], player)
		}
	})
}

func newRequestRegisterVictoryPost(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
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
