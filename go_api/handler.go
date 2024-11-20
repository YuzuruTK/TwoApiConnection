// handler.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func getHuman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(generateRandomHuman())
}
func getBeast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(generateRandomBeast())
}

/*************  ✨ Codeium Command ⭐  *************/
// generateRandomHuman returns a randomly generated HumanMonster based on
// the values in the "human.csv" file. The generated HumanMonster will have
// a randomly selected disposition, vitality, strength, willpower, dextery,
// and goals. The armor, weapon, trait, ring, and trinket will also be
// randomly selected based on the values in the "human.csv" file.
/******  bff20cd9-0ba9-4cf9-8dfb-2ee0021890ba  *******/
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
