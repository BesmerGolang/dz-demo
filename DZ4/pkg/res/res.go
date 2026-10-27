package res

import (
	"encoding/json"
	"net/http"
)

type ResOk struct {
	Data any `json:"data"`
}
type ResErr struct {
	Message string `json:"message"`
}

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
