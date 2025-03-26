package api

import (
	"encoding/json"
	"net/http"

	"github.com/vinny-tx/dotlight/internal/ai"
)

type AskRequest struct {
	Question string `json:"question"`
}

type AskResponse struct {
	Answer string `json:"answer"`
}

func AskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}

	var req AskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	answer, err := ai.GetAIResponse(req.Question)
	if err != nil {
		http.Error(w, "AI error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	ai.AddToHistory(req.Question, answer)

	json.NewEncoder(w).Encode(AskResponse{Answer: answer})
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	history := ai.GetHistory()
	json.NewEncoder(w).Encode(history)
}
