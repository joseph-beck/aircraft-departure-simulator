package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/joho/godotenv"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/db"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/service"
)

// Manage service is responsible for data management.
// Handles CRUD for data stored within the DB.
func main() {
	fmt.Println("ads-manage service")

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	conn := db.NewPGConnecter("db", db.DefaultCredentials())
	db := db.NewDB(conn)

	svc := service.NewManageService(db)

	err = rpc.Register(svc)
	if err != nil {
		panic(err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	http.Serve(listener, nil)
}
