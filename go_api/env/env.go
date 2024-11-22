package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const dotenv_path = "../.env"

func loadEnv() {
	log.Println("Loading any env file existent... ")
	godotenv.Load(dotenv_path)
}

func GetUploadthingApiToken() string {
	loadEnv()
	return os.Getenv("UPLOADTHING_TOKEN")
}
