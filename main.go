package main

import (
	"encoding/json"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "สวัสดีจาก Go 🚀",
	})
}

func main() {
	// API
	http.HandleFunc("/api/hello", hello)

	// serve frontend
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":8080", nil)
}
