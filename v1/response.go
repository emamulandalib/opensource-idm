package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response ...
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   []string    `json:"error,omitempty"`
}

// JSONResponse ...
func JSONResponse(w http.ResponseWriter, payload Response, statusCode int) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}
