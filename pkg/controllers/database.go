package controllers

import "database/sql"

var DB *sql.DB

// SetDB assigns the database connection to the global variable
func SetDB(database *sql.DB) {
    DB = database
}
