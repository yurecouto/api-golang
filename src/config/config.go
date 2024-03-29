package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ConnectString = ""
	ApiPort       = ""
	AccessKey     []byte
	RefreshKey    []byte
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	ApiPort = fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	ConnectString = fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	AccessKey = []byte(os.Getenv("SECRET_KEY_ACCESS"))
	RefreshKey = []byte(os.Getenv("SECRET_KEY_REFRESH"))
}
