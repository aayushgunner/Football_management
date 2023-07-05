package main
 
import (
    "database/sql"
    // "fmt"
	"encoding/json"
    _ "github.com/lib/pq"
	"net/http"
	"log"
)

type players struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
 
const (
    host     = "localhost"
    port     =  5432
    user     = "postgres"
    password = "golyanglyang"
    dbname   = "try"
)

func main() {
// //         // connection string
// //     psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
// //         // open database
// //     db, err := sql.Open("postgres", psqlconn)
// //     CheckError(err)
     
// //         // close database
// //     defer db.Close()x
 
// //         // check db
// //     err = db.Ping()
// //     CheckError(err)
 
// //     fmt.Println("Connected!")
// // 	rows, err := db.Query(`SELECT "first_name", "last_name" FROM "players"`)
// //     CheckError(err)
 
// //    defer rows.Close()
// //    for rows.Next() {
// //     var first_name string
// //     var last_name string
 
// //     err = rows.Scan(&first_name, &last_name)
// //     CheckError(err)
 
// //     fmt.Println(first_name, last_name)

	 
// }
 
// CheckError(err)
http.HandleFunc("/records", retrieveRecord)
err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
	
}

func retrieveRecord(w http.ResponseWriter, r *http.Request) {
	// Database connection string
	psqlconn := "host=localhost port=5432 user=postgres password=golyanglyang dbname=try sslmode=disable"
	
	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SELECT query
	rows, err := db.Query("SELECT first_name, last_name FROM players")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	students := []players{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var student players
		err := rows.Scan( &student.FirstName, &student.LastName)
		if err != nil {
			log.Println(err)
			continue
		}
		students = append(students, student)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
	
	}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}