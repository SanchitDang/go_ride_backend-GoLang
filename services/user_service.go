// services/user_service.go
package services

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang-API-save-data-sql/models"
)

// UserStoreData stores User data in the SQLite database.
func UserStoreData(db *sql.DB, data models.User) error {
	// Prepare the SQL statement for inserting data
	stmt, err := db.Prepare(`
		INSERT INTO users (Name, Email, Phone)
		VALUES (?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement with data values
	_, err = stmt.Exec(data.Name, data.Email, data.Phone)
	if err != nil {
		return err
	}

	return nil
}

// UserGetData retrieves User data from the SQLite database.
func UserGetData(db *sql.DB) ([]models.User, error) {
	// Retrieve all data from the User table
	rows, err := db.Query(`
		SELECT Name, Email, Phone
		FROM users
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dataList []models.User

	// Iterate over the rows and scan the data into a slice of User
	for rows.Next() {
		var data models.User
		err := rows.Scan(
			&data.Name,
			&data.Email,
			&data.Phone,
		)
		if err != nil {
			return nil, err
		}
		dataList = append(dataList, data)
	}

	return dataList, nil
}

// PostUserData handles the HTTP POST request for User data.
func PostUserData(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Decode the JSON payload
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if data.Name == "" || data.Email == "" || data.Phone == "" {
		http.Error(w, "Required fields cannot be empty", http.StatusBadRequest)
		return
	}

	// Store the data in the SQLite database
	err = UserStoreData(db, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response
	response := map[string]string{"message": "Data received and stored successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUserData handles the HTTP GET request for User data.
func GetUserData(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Retrieve data from the SQLite database
	dataList, err := UserGetData(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the retrieved data as a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataList)
}
