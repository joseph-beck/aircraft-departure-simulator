package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/db"
)

func main() {
	fmt.Println("ads-seed")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := db.NewPGConnecter("db", db.DefaultCredentials())
	db.NewDB(conn)
}
