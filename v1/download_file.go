package v1

import (
	"fmt"
	"net/http"
)

// DownloadFile ...
func DownloadFile(w http.ResponseWriter, req *http.Request) {
	url := "https://download.jetbrains.com/idea/ideaIU-2020.2.tar.gz"
	generalMsg := "There is something wrong during communication with the server. [ErrCode: 001]"

	resp, err := http.Head(url)

	if err != nil {
		response := Response{
			Success: false,
			Message: generalMsg,
			Error:   []string{err.Error()},
		}
		JSONResponse(w, response, http.StatusBadRequest)
		return
	}

	if resp.StatusCode != http.StatusOK {
		response := Response{
			Success: false,
			Message: generalMsg,
			Error:   []string{err.Error()},
		}
		JSONResponse(w, response, resp.StatusCode)
		return
	}

	contentLength := resp.ContentLength
	acceptRanges := resp.Header["Accept-Ranges"]

	if len(acceptRanges) == 0 {
		response := Response{
			Success: false,
			Message: generalMsg,
			Error:   []string{generalMsg},
		}
		JSONResponse(w, response, http.StatusBadRequest)
		return
	}

	if acceptRanges[0] != "bytes" {
		response := Response{
			Success: false,
			Message: "We are working on it currently. [ErrCode: 002]",
			Error:   []string{generalMsg},
		}
		JSONResponse(w, response, http.StatusBadRequest)
		return
	}

	fmt.Println(contentLength)
	JSONResponse(w, Response{
		Success: true,
		Message: "File is in the queue for download.",
	}, http.StatusOK)
}
