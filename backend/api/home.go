package api

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("👋 Welcome to Dotlight – AI-powered strategy assistant."))
}
