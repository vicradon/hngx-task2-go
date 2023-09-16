package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapis/personcrudapi/database"
)

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	db, err := database.NewDatabase("database.db")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodGet {
			rows, err := db.Query("SELECT * FROM persons")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Convert the rows to JSON
			jsonData, err := json.Marshal(rows)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(jsonData)

		} else if r.Method == http.MethodPost {

			insertQuery := "INSERT INTO persons (id, first_name, last_name, email) VALUES (?, ?, ?, ?)"
			if err := db.Exec(insertQuery, "3", "Fucker", "Dude", "fucker.dude@gmail.com"); err != nil {
				fmt.Println("Error inserting data:", err)
			}

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	})
	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}
