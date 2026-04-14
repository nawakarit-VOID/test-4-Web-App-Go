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

func setLang(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	json.NewDecoder(r.Body).Decode(&data)

	lang := data["lang"]

	http.SetCookie(w, &http.Cookie{
		Name:  "lang",
		Value: lang,
		Path:  "/",
	})

	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	// API
	http.HandleFunc("/api/hello", hello)
	http.HandleFunc("/api/post", postHandler)
	http.HandleFunc("/api/set-lang", setLang)
	// serve frontend
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
