package v1

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"sync"
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
	defer resp.Body.Close()

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
	chunkSize := int64(math.Ceil(float64(contentLength / 10)))
	fmt.Println(chunkSize)

	start := int64(0)
	end := chunkSize
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendChunkRequest(&wg, start, end)
		start = end + 1
		end = start + chunkSize
	}

	JSONResponse(w, Response{
		Success: true,
		Message: "File is in the queue for download.",
	}, http.StatusOK)
	wg.Wait()
}

func sendChunkRequest(wg *sync.WaitGroup, start int64, end int64) {
	fmt.Println(start, end)

	client := &http.Client{}
	url := "https://download.jetbrains.com/idea/ideaIU-2020.2.tar.gz"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err.Error())
	}

	rangeParam := fmt.Sprintf("bytes=%d-%d", start, end)
	req.Header.Add("Range", rangeParam)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	out, err := os.Create("downloads/" + rangeParam)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	fmt.Println(resp)
	wg.Done()
}
