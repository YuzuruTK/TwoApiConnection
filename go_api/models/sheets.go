package models

type Hand struct {
	Thumb   string
	Pointer string
	Middle  string
	Ringy   string
	Little  string
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type SheetBaseStats struct {
	Vitality  int `json:"vitality"`
	Strength  int `json:"strength"`
	Willpower int `json:"willpower"`
	Dextery   int `json:"dextery"`
}

type HumanMonster struct {
	BaseStats   SheetBaseStats `json:"stats"`
	Disposition string         `json:"disposition"`
	Goals       string         `json:"goals"`
	Armor       string         `json:"armor"`
	Weapon      string         `json:"weapon"`
	Trait       string         `json:"trait"`
	Ring        string         `json:"ring"`
	Trinket     string         `json:"trinket"`
}

type BeastMonster struct {
	BaseStats   SheetBaseStats `json:"stats"`
	Disposition string         `json:"disposition"`
	Desires     string         `json:"desires"`
	Armor       string         `json:"armor"`
	Attack      string         `json:"attack"`
	Trait       string         `json:"trait"`
	Trait2      string         `json:"trait2"`
	Ability     string         `json:"ability"`
}
