package httpserver

import (
	"encoding/json"
	"net/http"
)

type errorJSONResponse struct {
	Message string `json:"error"`
}

type resultJSONResponse struct {
	Result interface{} `json:"result"`
}

func resultResponse(w http.ResponseWriter, statusCode int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	message := resultJSONResponse{Result: msg}
	json.NewEncoder(w).Encode(message)
}

func errorResponse(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	message := errorJSONResponse{Message: msg}
	json.NewEncoder(w).Encode(message)
}
