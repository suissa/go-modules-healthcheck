package tools

import (
	"encoding/json"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Cria uma resposta JSON indicando que o servidor está ativo
	response := map[string]string{
		"status":  "OK",
		"message": "O servidor está rodando!",
	}

	// Define o cabeçalho e retorna a resposta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
