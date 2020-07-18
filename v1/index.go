package v1

import "net/http"

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "Open Source IDM API.",
	}
	JSONResponse(w, response, http.StatusOK)
}
