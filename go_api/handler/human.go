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

func RandomHumanMonsterHandler(res http.ResponseWriter, req *http.Request) {
	resEncoder := json.NewEncoder(res)

	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case http.MethodGet:
		log.Println("Handling GET request")
		res.WriteHeader(http.StatusOK)
		resEncoder.Encode(game.GenerateRandomHuman())

	case http.MethodPost:
		log.Println("Handling POST request")
		reqBody, err := io.ReadAll(req.Body)
		if err != nil {
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

		var randomMonster = game.GenerateRandomHuman()

		if !reqJson.SaveFile {
			log.Println("SaveFile is false, returning local response")
			res.WriteHeader(http.StatusOK)
			resEncoder.Encode(models.UploadHumanResponse{FileStatus: models.Local, FileContent: randomMonster})
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
			FileName: GetFileName("human"),
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
		resEncoder.Encode(models.UploadHumanResponse{FileStatus: models.Upload, FileContent: randomMonster})

	default:
		log.Printf("Method not allowed: %s", req.Method)
		res.WriteHeader(http.StatusMethodNotAllowed)
		resEncoder.Encode(models.SimpleResponse{Message: "Method Not Allowed"})
		return
	}
}
