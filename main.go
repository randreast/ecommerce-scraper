package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	limit := os.Getenv("LIMIT")
	fmt.Println(limit)
}
