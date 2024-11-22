package request

import (
	"bytes"
	"goapi/env"
	"io"
	"log"
	"net/http"
)

func PostRequest(url string, payload []byte) []byte {
	log.Println("Making POST request to URL:", url)
	apiKey := env.GetUploadthingApiToken()

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Uploadthing-Api-Key", apiKey)

	log.Println("Sending POST request")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending POST request: %v", err)
		return nil
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return nil
	}

	log.Println("POST request completed with status:", res.Status)
	log.Println("Response body:", string(body))
	return body
}
