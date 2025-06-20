package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/supabase-community/supabase-go"
)

func Todos(w http.ResponseWriter, r *http.Request) {
	client, _ := supabase.NewClient(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_KEY"),
		&supabase.ClientOptions{},
	)

	var todos []map[string]interface{}
	err := client.DB.From("todos").Select("*").Order("created_at", &supabase.OrderOpts{Ascending: false}).Execute(&todos)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
