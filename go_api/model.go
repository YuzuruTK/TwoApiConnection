// model.go
package main

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type HumanMonster struct {
	Disposition string `json:"disposition"`
	Vitality    int    `json:"vitality"`
	Strength    int    `json:"strength"`
	Willpower   int    `json:"willpower"`
	Dextery     int    `json:"dextery"`
	Goals       string `json:"goals"`
	Armor       string `json:"armor"`
	Weapon      string `json:"weapon"`
	Trait       string `json:"trait"`
	Ring        string `json:"ring"`
	Trinket     string `json:"trinket"`
}

type BeastMonster struct {
	Disposition string `json:"disposition"`
	Vitality    int    `json:"vitality"`
	Strength    int    `json:"strength"`
	Willpower   int    `json:"willpower"`
	Dextery     int    `json:"dextery"`
	Desires     string `json:"desires"`
	Armor       string `json:"armor"`
	Attack      string `json:"attack"`
	Trait       string `json:"trait"`
	Trait2      string `json:"trait2"`
	Ability     string `json:"ability"`
}

type preloadPayload struct {
	FileName string `json:"fileName"`
	FileSize int    `json:"fileSize"`
	FileType string `json:"fileType"`
}

type preloadResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	URL     string `json:"url"`
}

type File struct {
	Name string      `json:"name"`
	Size int         `json:"size"`
	Data interface{} `json:"data"`
}

type UploadRequest struct {
	Files              []File `json:"files"`
	ACL                string `json:"acl"`
	ContentDisposition string `json:"contentDisposition"`
}
