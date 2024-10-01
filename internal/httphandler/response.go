package httphandler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(w http.ResponseWriter, logg *slog.Logger, statusCode int, message string) {
	logg.Error("Response status: " + strconv.Itoa(statusCode) + " message: " + message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := errorResponse{Message: message}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logg.Error("Error encoding response: " + err.Error())
	}
}
