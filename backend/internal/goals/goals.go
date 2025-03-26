package goals

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

// Goal represents the structure of a goal
type Goal struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     string    `json:"dueDate"` // Change this to a string
	Status      string    `json:"status"`  // Add this field to store the goal status
	CreatedAt   time.Time `json:"createdAt"`
}

var (
	goals = []Goal{}
	mu    sync.Mutex
)

// CreateGoal creates a new goal and appends it to the goals list
func CreateGoal(title, description, status string, due time.Time) Goal {
	mu.Lock()
	defer mu.Unlock()

	goal := Goal{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		DueDate:     due.Format("2006-01-02"), // Format the date to YYYY-MM-DD
		Status:      status,                   // Add the status here
		CreatedAt:   time.Now(),
	}

	goals = append(goals, goal)
	return goal
}

// GetGoals retrieves all the goals
func GetGoals() []Goal {
	mu.Lock()
	defer mu.Unlock()
	return goals
}
