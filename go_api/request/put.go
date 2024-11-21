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
	apiKey := env.GetUploadthingApiToken()

	uploadRequestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(uploadRequestBody)

	// Create the multipart file part for the JSON file
	part, err := writer.CreateFormFile("file", "sample.json")
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON data to the multipart form file part
	_, err = part.Write(payload)
	if err != nil {
		log.Fatal(err)
	}

	// Close the writer to finish the multipart body
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("PUT", url, uploadRequestBody)

	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("X-Uploadthing-Api-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	println(string(body))
	return body
}
