package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/michaelbunch/notifications-minisys/db"
)

func main() {
	fs := http.FileServer(http.Dir("./frontend/dist/"))
	http.Handle("/", fs)
	http.HandleFunc("/api/notifications", apiGetNotifications)

	fmt.Println("Running API Server...")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

type notificationRequest struct {
	UserId int `json:"user_id"`
}

func apiGetNotifications(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
	}

	qsUserId := r.URL.Query().Get("user_id")
	userId, _ := strconv.Atoi(qsUserId)

	notifications := db.GetNotifications(db.Connection(), userId)

	payload, err := json.Marshal(notifications)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}
