package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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

	ConnectString = fmt.Sprintf("%s:%s@babar.db.elephantsql.com/xqnhqtbh",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}
