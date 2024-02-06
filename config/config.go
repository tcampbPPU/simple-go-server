package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Init(f string) {
	err := godotenv.Load(f)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}
