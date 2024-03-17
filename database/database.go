// database/database.go

package database

import (
    "database/sql"
    "github.com/spf13/viper"
)

// InitializeDB initializes the database connection and creates tables if they don't exist
func InitializeDB() (*sql.DB, error) {
    db, err := sql.Open(viper.GetString("database.driver"), viper.GetString("database.connection"))
    if err != nil {
        return nil, err
    }

    // Create the users table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            Name TEXT,
            Email TEXT,
            Phone TEXT
        )
    `)
    if err != nil {
        return nil, err
    }

    // Create the feedback table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS feedback (
            Date TEXT,
            Time TEXT,
            Feedback TEXT
        )
    `)
    if err != nil {
        return nil, err
    }

    return db, nil
}
