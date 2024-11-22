package request

import (
	"bytes"
	"goapi/env"
	"io"
	"log"
	"mime/multipart"
	"net/http"
)

func PutRequest(url string, payload []byte) []byte {
	log.Println("Making PUT request to URL:", url)
	apiKey := env.GetUploadthingApiToken()

	uploadRequestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(uploadRequestBody)

	// Create the multipart file part for the JSON file
	part, err := writer.CreateFormFile("file", "sample.json")
	if err != nil {
		log.Fatalf("Error creating form file: %v", err)
	}

	// Write the JSON data to the multipart form file part
	_, err = part.Write(payload)
	if err != nil {
		log.Fatalf("Error writing payload to form file: %v", err)
	}

	// Close the writer to finish the multipart body
	err = writer.Close()
	if err != nil {
		log.Fatalf("Error closing writer: %v", err)
	}

	req, _ := http.NewRequest("PUT", url, uploadRequestBody)

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("X-Uploadthing-Api-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending PUT request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	log.Println("PUT request completed with status:", res.Status)
	log.Println("Response body:", string(body))
	return body
}
