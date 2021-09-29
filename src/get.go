package songwhip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/atotto/clipboard"
)

func getRequestBody(trackUrl string) []byte {
	reqBody, err := json.Marshal(map[string]string{
		"url": trackUrl,
	})

	if err != nil {
		log.Fatal(err)
	}

	return reqBody
}

func getResponse(trackUrl string) *http.Response {
	const apiUrl = "https://songwhip.com/"
	var reqBody = getRequestBody(trackUrl)
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func getResponseBody(trackUrl string) map[string]interface{} {
	var resp = getResponse(trackUrl)

	defer resp.Body.Close()
	var body map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&body)

	return body
}

func getSongwhipUrl(trackUrl string) string {
	var responseBody = getResponseBody(trackUrl)
	var songwhipUrl = responseBody["url"]

	if songwhipUrl == nil {
		log.Fatal("Songwhip link not found")
	}

	return songwhipUrl.(string)
}

func copy(trackUrl string) {
	clipboard.WriteAll(trackUrl)
}

func Get(trackUrl string) {
	var songwhipUrl = getSongwhipUrl(trackUrl)
	fmt.Println(songwhipUrl)
	copy(songwhipUrl)
}
