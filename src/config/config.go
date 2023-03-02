package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DatabaseConnectString = ""
	Port                  = ""
	SecretKey             []byte
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port = fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	DatabaseConnectString = fmt.Sprintf("postgres://%s:%s@%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}
