package handler

import "time"

func GetFileName(monsterType string) string {
	return monsterType + "_monster_" + time.Now().Format("02-Jan-2006_150405") + ".json"
}
