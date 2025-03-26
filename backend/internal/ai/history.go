package ai

type Message struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

var history []Message

func AddToHistory(q, a string) {
	history = append(history, Message{
		Question: q,
		Answer:   a,
	})
}

func GetHistory() []Message {
	return history
}
