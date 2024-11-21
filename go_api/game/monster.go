package game

import (
	"encoding/csv"
	"goapi/constants"
	"goapi/models"
	"log"
	"math/rand"
	"os"
)

// This function receives an SheetBaseStats Obj with vitality meaning its multiplier
func getBaseStats(multipliers models.SheetBaseStats) models.SheetBaseStats {
	return models.SheetBaseStats{
		Vitality:  RollDice(10, multipliers.Vitality),
		Strength:  multipliers.Strength,
		Willpower: multipliers.Willpower,
		Dextery:   multipliers.Dextery,
	}
}

func GenerateRandomHuman() models.HumanMonster {
	file, err := os.Open(constants.HumanMonsterCsvPath)
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
	var baseStats = getBaseStats(constants.BaseStats.Human[rand.Intn(5)])
	var goal string = choose(constants.Goals.Human)
	var disposition string = choose(constants.GeneralDispositions)

	// return value
	var responseMonster = models.HumanMonster{
		BaseStats:   baseStats,
		Goals:       goal,
		Disposition: disposition,
	}

	// Define a mapping for dynamic handling
	fieldMapping := []struct {
		hand     string
		relation string
		target   *string
	}{
		{monsterHand.Thumb, constants.HandRelation.Human.Thumb, &responseMonster.Armor},
		{monsterHand.Pointer, constants.HandRelation.Human.Pointer, &responseMonster.Weapon},
		{monsterHand.Middle, constants.HandRelation.Human.Middle, &responseMonster.Trait},
		{monsterHand.Ringy, constants.HandRelation.Human.Ringy, &responseMonster.Ring},
		{monsterHand.Little, constants.HandRelation.Human.Little, &responseMonster.Trinket},
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
		{monsterHand.Thumb, constants.HandRelation.Human.Thumb, &responseMonster.Armor},
		{monsterHand.Pointer, constants.HandRelation.Human.Pointer, &responseMonster.Attack},
		{monsterHand.Middle, constants.HandRelation.Human.Middle, &responseMonster.Trait},
		{monsterHand.Ringy, constants.HandRelation.Human.Ringy, &responseMonster.Trait2},
		{monsterHand.Little, constants.HandRelation.Human.Little, &responseMonster.Ability},
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
