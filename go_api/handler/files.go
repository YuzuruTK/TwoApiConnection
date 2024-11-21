package handler

import "time"

func GetFileName(monsterType string) string {
	return monsterType + "_monster_" + time.Now().Format("2006-01-02 15:04:05") + ".json"
}
