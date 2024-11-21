package request

import (
	"bytes"
	"goapi/env"
	"io"
	"net/http"
)

func PostRequest(url string, payload []byte) []byte {
	apiKey := env.GetUploadthingApiToken()

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Uploadthing-Api-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil
	}

	return body
}
