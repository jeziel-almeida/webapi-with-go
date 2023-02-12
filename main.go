package main

import (
	"github.com/jeziel-almeida/webapi-with-go/database"
	"github.com/jeziel-almeida/webapi-with-go/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
