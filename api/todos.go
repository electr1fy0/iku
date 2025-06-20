package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/supabase-community/supabase-go"
)

type TodoRequest struct {
	Task string `json:"task"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var req TodoRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Task == "" {
		http.Error(w, "Missing task", http.StatusBadRequest)
		return
	}

	client, _ := supabase.NewClient(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_KEY"),
		&supabase.ClientOptions{},
	)

	todo := map[string]interface{}{
		"task": req.Task,
	}

	err := client.DB.From("todos").Insert(todo).Execute(&todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Added task: " + req.Task))
}
