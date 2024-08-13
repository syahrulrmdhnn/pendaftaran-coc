package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/syahrulrmdhnn/pendaftaran-coc/backend/config"
	"github.com/syahrulrmdhnn/pendaftaran-coc/backend/routes"
)

func main() {
	config.InitDB()
	routes.Routes()

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	port := os.Getenv("APP_PORT")

	fmt.Println("Server running on port", port)
	server(":"+port)
}

func server(addr string) {
	http.ListenAndServe(addr, nil)
}