package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {

	url := "https://api.uploadthing.com/v7/prepareUpload"

	fileName := "monster_" + time.Now().Format("2006-01-02 15:04:05") + ".json"

	jsonData, err := json.Marshal(generateRandomHuman())

	log.Print("Monster generated")
	if err != nil {
		fmt.Println(err)
	}

	preloadPayload := preloadPayload{
		FileName: fileName,
		FileSize: len(jsonData),
		FileType: "application/json",
	}
	var preloadResponse preloadResponse
	preloadPayloadJSON, err := json.Marshal(preloadPayload)

	if err != nil {
		fmt.Println(err)
	}

	log.Print("Preload payload created, doing POST request")

	body := doPostRequest(url, preloadPayloadJSON)

	log.Print("POST request done")

	err = json.Unmarshal(body, &preloadResponse)

	if err != nil {
		fmt.Println(err)
	}

	url = preloadResponse.URL
	log.Printf("Preload response created, doing PUT request on : %s", url)

	body = doPutRequest(url, jsonData)

	log.Print("PUT request done")
}
