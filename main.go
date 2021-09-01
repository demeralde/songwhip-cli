package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.design/x/clipboard"
)

func getApiUrl() string {
	argLength := len(os.Args[1:])

	if argLength <= 0 {
		log.Fatal("Song URL is required.")
	}

	return os.Args[1]
}

func getRequestBody() []byte {
	reqBody, err := json.Marshal(map[string]string{
		"url": getApiUrl(),
	})

	if err != nil {
		log.Fatal(err)
	}

	return reqBody
}

func getResponse() *http.Response {
	var apiUrl = "https://songwhip.com/"
	var reqBody = getRequestBody()
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getResponseBody() map[string]interface{} {
	var resp = getResponse()

	defer resp.Body.Close()
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)

	return body
}

func getSongwhipTrackUrl() string {
	var responseBody = getResponseBody()
	var trackUrl = responseBody["url"].(string)
	return trackUrl
}

func main() {
	var trackUrl = getSongwhipTrackUrl()
	clipboard.Write(clipboard.FmtText, []byte(trackUrl))
	fmt.Printf("Copied %v to clipboard", trackUrl)
}
