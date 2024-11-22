package game

import (
	"encoding/csv"
	"goapi/constants"
	"goapi/models"
	"log"
	"math/rand"
	"os"
)

func GenerateRandomBeast() models.BeastMonster {
	file, err := os.Open(constants.BeastMonsterCsvPath)
	if err != nil {
		log.Fatal(err)
		file.Close()
	}

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var monsterHand = getRandomHand()
	var baseStats = getBaseStats(constants.BaseStats.Beast[rand.Intn(5)])
	var goal string = choose(constants.Goals.Beast)
	var disposition string = choose(constants.GeneralDispositions)

	var responseMonster = models.BeastMonster{
		BaseStats:   baseStats,
		Desires:     goal,
		Disposition: disposition,
	}

	// Define a mapping for dynamic handling
	fieldMapping := []struct {
		hand     string
		relation string
		target   *string
	}{
		{monsterHand.Thumb, constants.HandRelation.Beast.Thumb, &responseMonster.Armor},
		{monsterHand.Pointer, constants.HandRelation.Beast.Pointer, &responseMonster.Attack},
		{monsterHand.Middle, constants.HandRelation.Beast.Middle, &responseMonster.Trait},
		{monsterHand.Ringy, constants.HandRelation.Beast.Ringy, &responseMonster.Trait2},
		{monsterHand.Little, constants.HandRelation.Beast.Little, &responseMonster.Ability},
	}

	// Iterate over the records
	for _, item := range records {
		for _, field := range fieldMapping {
			if (field.hand + field.relation) == (item[constants.Column.Type] + item[constants.Column.Category]) {
				*field.target = item[constants.Column.Description]
			}
		}
	}

	return responseMonster
}
