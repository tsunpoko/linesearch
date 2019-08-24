package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "PUT, GET, POST, DELETE, OPTIONS")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
