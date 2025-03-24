package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type Visit struct {
	Time      time.Time `json:"time"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"userAgent"`
}

var db *sql.DB

func initDB() {
	db = clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"clickhouse-service:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "default",
		},
	})
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`
    CREATE TABLE IF NOT EXISTS visits (
        time DateTime,
        ip String,
        userAgent String
    ) ENGINE = MergeTree()
    ORDER BY time;`)
	log.Println("Table created")
}

func logVisitHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/visit\n")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if db == nil {
		http.Error(w, "Database not ready", http.StatusServiceUnavailable)
		return
	}
	var visit Visit
	json.NewDecoder(r.Body).Decode(&visit)
	visit.Time = time.Now()

	_, err := db.Exec("INSERT INTO visits (time, ip, userAgent) VALUES (?, ?, ?)", visit.Time, visit.IP, visit.UserAgent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("INSERTED: %+v\n", visit)

	w.WriteHeader(http.StatusOK)
}

func getVisitsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/visits\n")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if db == nil {
		http.Error(w, "Database not ready", http.StatusServiceUnavailable)
		return
	}
	rows, err := db.Query("SELECT time, ip, userAgent FROM visits ORDER BY time ASC")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var visits []Visit
	for rows.Next() {
		var visit Visit
		rows.Scan(&visit.Time, &visit.IP, &visit.UserAgent)
		visits = append(visits, visit)
	}

	json.NewEncoder(w).Encode(visits)
	log.Printf("RETRIEVED %d visits\n", len(visits))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/health\n")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/ready\n")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if db == nil {
		http.Error(w, "Database not ready", http.StatusServiceUnavailable)
		return
	}
	err := db.Ping()
	if err != nil {
		http.Error(w, "Database not ready", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	initDB()
	http.HandleFunc("/visit", logVisitHandler)
	http.HandleFunc("/visits", getVisitsHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ready", readyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
