package main

import (
	"fmt"
	"log"
	"os"

	"example.com/practice-clean-architecture/driver"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env not found")
	}

	log.Println("Server running...")
	// サーバー起動
	driver.Serve(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
