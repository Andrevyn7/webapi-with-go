package main

import (
	"github.com/Andrevyn7/webapi-with-go.git/database"
	"github.com/Andrevyn7/webapi-with-go.git/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
