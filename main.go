package main

import "github.com/jeziel-almeida/webapi-with-go/server"

func main() {
	server := server.NewServer()

	server.Run()
}
