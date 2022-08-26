package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/michaelbunch/notifications-minisys/db"
)

func main() {
	http.HandleFunc("/api/notifications", apiGetNotifications)

	fmt.Println("Running API Server...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func apiGetNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}
	notifications := db.GetNotifications(db.Connection(), 1)

	payload, err := json.Marshal(notifications)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
