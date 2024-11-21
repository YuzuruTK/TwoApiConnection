// handler.go
package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// func getHuman(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(generateRandomHuman())
// }
// func getBeast(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	json.NewEncoder(w).Encode(generateRandomBeast())
// }

func generateRandomHuman() HumanMonster {
	csv_file := "human.csv"
	file, err := os.Open("csv/" + csv_file)

	if err != nil {
		log.Fatal(err)
		file.Close()
	}

	dispositions := []string{"Friendly", "Curious", "Indifferent", "Aggressive", "Hostile"}
	goals := []string{"Help the most noble Traveler.", "Recruit forces to hunt down a fearsome local legend.", "Find their way back to their lost Safe Haven.", "Understand the nature of a set of ruins.", "Defeat the most infamous Traveler."}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	disposition := dispositions[rand.Intn(len(dispositions))]
	goal := goals[rand.Intn(len(goals))]

	status := []int{}
	switch rand.Intn(5) {
	case 0:
		status = []int{1, 2, 0, 1}
	case 1:
		status = []int{rollDice(10, 1), 4, 0, 1}
	case 2:
		status = []int{rollDice(10, 2), 6, 1, 2}
	case 3:
		status = []int{rollDice(10, 3), 8, 1, 2}
	case 4:
		status = []int{rollDice(10, 4), 10, 2, 4}
	}

	typesOfFinger := []string{"LL", "LA", "LW", "AL", "AA", "AW", "WL", "WA", "WW"}

	thumb := typesOfFinger[rollDice(9, 1)-1]
	pointer := typesOfFinger[rollDice(9, 1)-1]
	middle := typesOfFinger[rollDice(9, 1)-1]
	ringy := typesOfFinger[rollDice(9, 1)-1]
	little := typesOfFinger[rollDice(9, 1)-1]

	Armor := ""
	Weapon := ""
	Trait := ""
	Ring := ""
	Trinket := ""

	for _, item := range records {
		if item[0] == thumb && item[1] == "Armor" {
			Armor = item[2]
		}
		if item[0] == pointer && item[1] == "Weapon" {
			Weapon = item[2]
		}
		if item[0] == middle && item[1] == "Trait" {
			Trait = item[2]
		}
		if item[0] == ringy && item[1] == "Ring" {
			Ring = item[2]
		}
		if item[0] == little && item[1] == "Trinket" {
			Trinket = item[2]
		}
	}

	return HumanMonster{
		Disposition: disposition,
		Vitality:    status[0],
		Strength:    status[1],
		Willpower:   status[2],
		Dextery:     status[3],
		Goals:       goal,
		Armor:       Armor,
		Weapon:      Weapon,
		Trait:       Trait,
		Ring:        Ring,
		Trinket:     Trinket,
	}
}

func generateRandomBeast() BeastMonster {
	csv_file := "beast.csv"
	file, err := os.Open("csv/" + csv_file)

	if err != nil {
		log.Fatal(err)
		file.Close()
	}
	dispositions := []string{"Friendly", "Curious", "Indifferent", "Aggressive", "Hostile"}
	goals := []string{"Mark their territory", "Defend their young from hunters", "Hunt for food.", "Flee from impending ecological disaster", "Find their missing brood"}
	status := []int{}

	randomNumber := rand.Intn(5)
	switch {
	case randomNumber == 0:
		status = []int{1, 2, 0, 1}
	case randomNumber == 1:
		status = []int{rollDice(10, 1), 4, 0, 1}
	case randomNumber == 2:
		status = []int{rollDice(10, 2), 6, 1, 2}
	case randomNumber == 3:
		status = []int{rollDice(10, 3), 8, 1, 2}
	case randomNumber == 4:
		status = []int{rollDice(10, 4), 10, 2, 4}
	}

	typesOfFinger := []string{"LL", "LA", "LW", "AL", "AA", "AW", "WL", "WA", "WW"}

	thumb := typesOfFinger[rollDice(9, 1)-1]
	pointer := typesOfFinger[rollDice(9, 1)-1]
	middle := typesOfFinger[rollDice(9, 1)-1]
	ringy := typesOfFinger[rollDice(9, 1)-1]
	little := typesOfFinger[rollDice(9, 1)-1]

	// Roll dice
	Disposition := dispositions[rand.Intn(5)]
	Vitality := status[0]
	Strength := status[1]
	Willpower := status[2]
	Dextery := status[3]
	Goal := goals[rand.Intn(4)]
	// Get from CSV
	Armor := ""
	Attack := ""
	Trait := ""
	Trait2 := ""
	Ability := ""

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range records {
		if item[0] == thumb && item[1] == "Armor" {
			Armor = item[2]
		}
		if item[0] == pointer && item[1] == "Attack" {
			Attack = item[2]
		}
		if item[0] == middle && item[1] == "Trait" {
			Trait = item[2]
		}
		if item[0] == ringy && item[1] == "Trait2" {
			Trait2 = item[2]
		}
		if item[0] == little && item[1] == "Ability" {
			Ability = item[2]
		}
	}
	return BeastMonster{
		Disposition: Disposition,
		Vitality:    Vitality,
		Strength:    Strength,
		Willpower:   Willpower,
		Dextery:     Dextery,
		Desires:     Goal,
		Armor:       Armor,
		Attack:      Attack,
		Trait:       Trait,
		Trait2:      Trait2,
		Ability:     Ability,
	}
}

func rollDice(d int, s int) int {
	rand.Seed(time.Now().UnixNano())
	dice := 0
	for i := 0; i < s; i++ {
		dice += rand.Intn(d) + 1
	}
	return dice
}

func doPostRequest(url string, payload []byte) []byte {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Uploadthing-Api-Key", API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return body
}

func doPutRequest(url string, payload []byte) []byte {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")

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
	req.Header.Add("X-Uploadthing-Api-Key", API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	println(string(body))
	return body
}
