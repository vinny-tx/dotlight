package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vinny-tx/dotlight/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", api.HomeHandler)
	http.HandleFunc("/api/ping", api.PingHandler)
	http.HandleFunc("/api/ask", api.AskHandler)

	fmt.Println("🚀 Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("❌ Server error:", err)
	}
}
