package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/vinny-tx/dotlight/api"
	"github.com/vinny-tx/dotlight/internal/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()

	// Wrap your handlers with CORS middleware
	mux.Handle("/", middleware.EnableCORS(http.HandlerFunc(api.HomeHandler)))
	mux.Handle("/api/ping", middleware.EnableCORS(http.HandlerFunc(api.PingHandler)))
	mux.Handle("/api/ask", middleware.EnableCORS(http.HandlerFunc(api.AskHandler)))
	mux.Handle("/api/history", middleware.EnableCORS(http.HandlerFunc(api.HistoryHandler)))
	mux.Handle("/api/goals", middleware.EnableCORS(http.HandlerFunc(api.GetGoalsHandler)))
	mux.Handle("/api/goals/create", middleware.EnableCORS(http.HandlerFunc(api.CreateGoalHandler)))

	fmt.Println("🚀 Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("❌ Server error:", err)
	}
}
