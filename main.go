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

func postHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	json.NewDecoder(r.Body).Decode(&data)

	name := data["name"]

	json.NewEncoder(w).Encode(map[string]string{
		"reply": "สวัสดี " + name,
	})
}

func main() {
	// API
	//http.HandleFunc("/api/hello", hello)
	http.HandleFunc("/api/post", postHandler)
	// serve frontend
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
