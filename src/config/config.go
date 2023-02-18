package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

var (
	ConnectString = ""
	Port          = 0
	SecretKey     []byte
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 7777
	}

	pgUrl, erro := pq.ParseURL(os.Getenv("ELEPHANT_SQL"))
	if erro != nil {
		log.Fatal("erro", erro)
	}

	ConnectString = pgUrl

	SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}
