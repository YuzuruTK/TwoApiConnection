package handler

import (
	"encoding/json"
	"goapi/game"
	"goapi/models"
	"goapi/request"
	"io"
	"log"
	"net/http"
)

func RandomBeastMonsterHandler(res http.ResponseWriter, req *http.Request) {
	resEncoder := json.NewEncoder(res)

	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case http.MethodGet:
		log.Println("Handling GET request")
		res.WriteHeader(http.StatusOK)
		resEncoder.Encode(game.GenerateRandomBeast())

	case http.MethodPost:
		log.Println("Handling POST request")
		reqBody, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			resEncoder.Encode(models.SimpleResponse{Message: "Failed to Read Request Body"})
			return
		}
		defer req.Body.Close()

		var reqJson models.UploadBody
		err = json.Unmarshal(reqBody, &reqJson)
		if err != nil {
			log.Printf("Error unmarshalling request body: %v", err)
			res.WriteHeader(http.StatusBadRequest)
			resEncoder.Encode(models.SimpleResponse{Message: "Invalid Payload"})
			return
		}

		var randomMonster = game.GenerateRandomBeast()

		if !reqJson.SaveFile {
			log.Println("SaveFile is false, returning local response")
			res.WriteHeader(http.StatusOK)
			resEncoder.Encode(models.UploadBeastResponse{FileStatus: models.Local, FileContent: randomMonster})
			return
		}

		monsterJson, err := json.Marshal(randomMonster)
		if err != nil {
			log.Printf("Error marshalling monster JSON: %v", err)
			res.WriteHeader(http.StatusInternalServerError)
			resEncoder.Encode(models.SimpleResponse{Message: "Server could not generate Json Data"})
			return
		}

		prepareResponse, err := request.PrepareUpload(models.PrepareUploadPayload{
			FileName: GetFileName("beast"),
			FileSize: len(monsterJson),
			FileType: "json",
		})
		if err != nil {
			log.Printf("Error preparing upload: %v", err)
			res.WriteHeader(http.StatusBadRequest)
			resEncoder.Encode(models.SimpleResponse{Message: err.Error()})
			return
		}

		request.PutRequest(prepareResponse.URL, monsterJson)
		log.Println("File uploaded successfully")
		res.WriteHeader(http.StatusCreated)
		resEncoder.Encode(models.UploadBeastResponse{FileStatus: models.Upload, FileContent: randomMonster})

	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
		resEncoder.Encode(models.SimpleResponse{Message: "Method Not Allowed"})
		return
	}
}
