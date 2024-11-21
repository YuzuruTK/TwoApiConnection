package env

import (
	"os"

	"github.com/joho/godotenv"
)

const dotenv_path = "../.env"

func loadEnv() {
	godotenv.Load(dotenv_path)
}

func GetUploadthingApiToken() string {
	loadEnv()
	return os.Getenv("UPLOADTHING_TOKEN")
}
