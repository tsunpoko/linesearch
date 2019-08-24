package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", " Origin, X-Requested-With, Content-Type, Accept")

	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
