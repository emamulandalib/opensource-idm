package v1

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Success: true,
		Message: "Open Source IDM API.",
	}
	JsonResponse(w, response, http.StatusOK)
}
