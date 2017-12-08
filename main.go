package main

import (
	"encoding/json"
	"log"
	"os"
	"web-status/api"
	"web-status/db"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	connectionString := "user=postgres dbname=web-status sslmode=disable"
	handler, dbErr := db.MakeDatabaseHandler(connectionString)
	if dbErr != nil {
		log.Fatal(err)
	}
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Println("Starting web server on address ", config.ServerAddress)
	api.Run(config.ServerAddress, handler)
}
