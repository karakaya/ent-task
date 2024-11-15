package utils

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Message struct {
	Message string `json:"message"`
}

func WriteJSONError(logger zerolog.Logger, w http.ResponseWriter, statusCode int, err error) {
	logger.Error().Err(err).Msg(err.Error())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

func WriteJSONMessage(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(Message{Message: message})
}
