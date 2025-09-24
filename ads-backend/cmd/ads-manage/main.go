package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/db"
)

// Manage service is responsible for data management.
// Handles CRUD for data stored within the DB.
func main() {
	fmt.Println("ads-manage service")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := db.NewPGConnecter("db", db.DefaultCredentials())
	db := db.NewDB(conn)

	maxWeights, err := db.GetMaximumWeights()
	if err != nil {
		log.Fatalf("Error fetching maximum weights: %v", err)
	}

	for _, mw := range maxWeights {
		fmt.Printf("%+v\n", mw)
	}
}
