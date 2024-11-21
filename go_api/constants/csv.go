package constants

import (
	"os"
	"path/filepath"
)

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

var BeastMonsterCsvPath = getCurrentDir() + "/csv/beast.csv"

var HumanMonsterCsvPath = getCurrentDir() + "/csv/human.csv"

// Recognize csv fields
var Column = struct {
	Type        int
	Category    int
	Description int
}{
	Type:        0,
	Category:    1,
	Description: 2,
}
