package server

import (
	"encoding/json"
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/healthz", healthHandler)
	mux.HandleFunc("/ping", pingHandler)
	return mux
}

type response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	writeJSON(w, http.StatusOK, response{
		Message: "urban-invention is running",
		Status:  "ok",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, response{
		Message: "healthy",
		Status:  "ok",
	})
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, response{
		Message: "pong",
		Status:  "ok",
	})
}

func writeJSON(w http.ResponseWriter, statusCode int, payload response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
