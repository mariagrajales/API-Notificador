package main

import (
	"api/notification/src/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    HOST := os.Getenv("HOST_SERVER")
    PORT := os.Getenv("PORT_SERVER")

    if HOST == "" || PORT == "" {
        log.Fatal("HOST_SERVER or PORT_SERVER is not set")
    }

    srv := server.NewServer(HOST, PORT)
    srv.Run()
}
