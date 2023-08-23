package internal

import (
	"fmt"
	"log"
	"path"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println(path.Dir("../"))

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading env file : %v", err)
	}
}
