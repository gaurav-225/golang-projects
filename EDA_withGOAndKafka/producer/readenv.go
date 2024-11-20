package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main1() {
	godotenv.Load(".env")

	/*
	envs, err := godotenv.Read(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	name := envs["NAME"]
	*/

	val := os.Getenv("KAFKA_ADDRESS")
	fmt.Println(val)
	calling()
}

func calling() {
	val := os.Getenv("KAFKA_ADDRESS")
	
	log.Println(val)
}
