package game

import (
	"goapi/constants"
	"goapi/models"
	"math/rand"
	"time"
)

func RollDice(dice int, qtt int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	result := 0
	for i := 0; i < qtt; i++ {
		result += rand.Intn(dice) + 1
	}
	return result
}

func choose(items []string) string {
	if len(items) < 1 {
		return ""
	}

	r := rand.New(rand.NewSource(time.Now().Unix())) // Set random seed
	return items[r.Intn(len(items))]
}

func getRandomHand() models.Hand {
	return models.Hand{
		Thumb:   constants.TypesOfFinger[RollDice(9, 1)-1],
		Pointer: constants.TypesOfFinger[RollDice(9, 1)-1],
		Middle:  constants.TypesOfFinger[RollDice(9, 1)-1],
		Ringy:   constants.TypesOfFinger[RollDice(9, 1)-1],
		Little:  constants.TypesOfFinger[RollDice(9, 1)-1],
	}
}
