package request

import (
	"encoding/json"
	"goapi/constants"
	"goapi/models"
)

func PrepareUpload(payload models.PrepareUploadPayload) (models.PrepareUploadResponse, error) {
	var response models.PrepareUploadResponse

	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return models.PrepareUploadResponse{}, err
	}

	postResponseBody := PostRequest(constants.URL.PrepareUpload, jsonBody)
	json.Unmarshal(postResponseBody, &response)

	return response, nil
}
