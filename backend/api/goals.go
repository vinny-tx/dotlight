package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vinny-tx/dotlight/internal/goals"
)

func CreateGoalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Id          string `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		DueDate     string `json:"dueDate"`
		Status      string `json:"status"` // Add status here
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	dueDate, err := time.Parse("2006-01-02", input.DueDate)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	goal := goals.CreateGoal(input.Title, input.Description, input.Status, dueDate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(goal)
}

func GetGoalsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	allGoals := goals.GetGoals()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allGoals)
}
