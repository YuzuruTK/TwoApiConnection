package constants

import (
	"goapi/models"
)

var TypesOfFinger = []string{"LL", "LA", "LW", "AL", "AA", "AW", "WL", "WA", "WW"}

var GeneralDispositions = []string{
	"Friendly",
	"Curious",
	"Indifferent",
	"Aggressive",
	"Hostile",
}

var HandRelation = struct {
	Human models.Hand
	Beast models.Hand
}{
	Human: models.Hand{
		Thumb:   "Armor",
		Pointer: "Weapon",
		Middle:  "Trait",
		Ringy:   "Ring",
		Little:  "Trinket",
	},
	Beast: models.Hand{
		Thumb:   "Armor",
		Pointer: "Attack",
		Middle:  "Trait",
		Ringy:   "Trait2",
		Little:  "Ability",
	},
}

var Goals = struct {
	Human []string
	Beast []string
}{
	Human: []string{
		"Help the most noble Traveler.",
		"Recruit forces to hunt down a fearsome local legend.",
		"Find their way back to their lost Safe Haven.",
		"Understand the nature of a set of ruins.",
		"Defeat the most infamous Traveler.",
	},
	Beast: []string{
		"Mark their territory",
		"Defend their young from hunters",
		"Hunt for food.",
		"Flee from impending ecological disaster",
		"Find their missing brood",
	},
}

var BaseStats = struct {
	Human []models.SheetBaseStats
	Beast []models.SheetBaseStats
}{
	Human: []models.SheetBaseStats{
		{Vitality: 0, Strength: 0, Willpower: 0, Dextery: 0},
		{Vitality: 1, Strength: 2, Willpower: 1, Dextery: 1},
		{Vitality: 2, Strength: 4, Willpower: 2, Dextery: 2},
		{Vitality: 3, Strength: 8, Willpower: 4, Dextery: 4},
		{Vitality: 4, Strength: 10, Willpower: 8, Dextery: 8},
	},
	Beast: []models.SheetBaseStats{
		{Vitality: 0, Strength: 2, Willpower: 0, Dextery: 1},
		{Vitality: 1, Strength: 4, Willpower: 0, Dextery: 1},
		{Vitality: 2, Strength: 6, Willpower: 1, Dextery: 2},
		{Vitality: 3, Strength: 8, Willpower: 1, Dextery: 2},
		{Vitality: 4, Strength: 10, Willpower: 2, Dextery: 4},
	},
}
