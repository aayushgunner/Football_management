
package main

import (
	"log"
	//"net/http"
	"encoding/csv"
	//"github.com/gorilla/mux"
	//"github.com/gorilla/handlers"
	//"football_server/pkg/routes"
	"database/sql"
    _ "github.com/lib/pq"
	"fmt"
	"os"
	"io"
	"path/filepath"
	"football_server/pkg/controllers"
)
var db *sql.DB
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "golyanglyang"
	dbname   = "try"
)
// func setupDatabase() (*sql.DB, error) {
//     // Replace the connection string with your actual PostgreSQL database details
//     db, err := sql.Open("postgres", "postgres://user:password@host:port/database?sslmode=disable")
//     if err != nil {
//         return nil, err
//     }
//     return db, nil
// }
func main() {
	// Create the connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Open the database connection
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	absPath, err := filepath.Abs("database/standings/standings.csv")
if err != nil {
	log.Fatal(err)
}
file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	controllers.SetDB(db)
	for {
		// Read each row from the CSV file
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	
		fmt.Println("helloworld")
		// Prepare the SQL INSERT statement
		stmt, err := db.Prepare("INSERT INTO standings(team_rank, points,goalDiff, team_id,team_name,matches_played,wins,draws,losses,goals_for,goals_against,home_games,home_wins,home_draws,home_losses,home_goals_for,home_goals_against,away_games,away_wins,away_draws,away_losses,away_goals_for,awaY_goals_against ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23)")
		if err != nil {
			log.Fatal(err)
		}
	
		// Execute the INSERT statement with row data
		fmt.Println(row[0], row[1], row[2], row[3], row[4],row[5],row[6],row[7],row[8],row[9],row[10],row[11],row[12],row[13],row[14],row[15],row[16],row[17],row[18],row[19],row[20],row[21],row[22])
		_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4],row[5],row[6],row[7],row[8],row[9],row[10],row[11],row[12],row[13],row[14],row[15],row[16],row[17],row[18],row[19],row[20],row[21],row[22])
		
		if err != nil {
			log.Fatal(err)
		}
	}
}