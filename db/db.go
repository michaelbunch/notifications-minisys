package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Connection creates a database connection
func Connection() *sql.DB {
	c := mysql.NewConfig()
	c.User = os.Getenv("DB_User")
	c.Passwd = os.Getenv("DB_Password")
	c.DBName = os.Getenv("DB_Name")
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Fatalf("Error creating database connection. %v", err)
	}
	return db
}

// Notification is a notification object
type Notification struct {
	ID        int    `json:"id"`
	Subject   string `json:"subject"`
	Domain    string `json:"domain"`
	ActorId   int    `json:"actor_id"`
	ActorName string `json:"actor_name"`
	IsRead    bool   `json:"isRead"`
}

// GetNotifications fetches a user specific list of notifications from the database
func GetNotifications(conn *sql.DB, userId int) []Notification {
	rows, err := conn.Query("SELECT n.id, n.subject, n.domain, n.actor_id, ua.name `actor_name`, n.is_read FROM notifications n INNER JOIN users ua ON n.actor_id = ua.id WHERE notifier_id = ?", userId)
	if err != nil {
		log.Fatal(err)
	}

	var notifications []Notification
	for rows.Next() {
		var id int
		var subject string
		var domain string
		var actorId int
		var actorName string
		var isRead int
		if err := rows.Scan(&id, &subject, &domain, &actorId, &actorName, &isRead); err != nil {
			log.Fatal(err)
		}
		notification := Notification{
			ID:        id,
			Subject:   subject,
			Domain:    domain,
			ActorId:   actorId,
			ActorName: actorName,
			IsRead:    convertReadStatus(isRead),
		}
		notifications = append(notifications, notification)
	}
	return notifications
}

func convertReadStatus(s int) bool {
	if s == 1 {
		return true
	} else {
		return false
	}
}
