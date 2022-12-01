package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	fmt.Println("Server starting up")
	http.HandleFunc("/current", getCurrent)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/decrement", decrementCounter)
	http.HandleFunc("/reset", reset)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}

func getCurrent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}
	current, err := ReadDB()
	if err != nil {
		http.Error(w, "Error reading from DB", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]int)
	resp["CurrentCount"] = current
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error parsing data", http.StatusInternalServerError)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Fatal(err)
	}
	return

}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}
	current, err := ReadDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	current++
	err = WriteDB(current)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func decrementCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}
	current, err := ReadDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	current--
	err = WriteDB(current)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func reset(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusBadRequest)
	}
	err := ResetDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
