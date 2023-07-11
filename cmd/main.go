// package main

// import (
// 	"log"
// 	"net/http"
// 	"encoding/csv"
// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/handlers"
// 	"football_server/pkg/routes"
// 	"database/sql"
//     _ "github.com/lib/pq"
// 	"fmt"
// 	"os"
// 	"io"
// 	"path/filepath"
// 	"football_server/pkg/controllers"
// )
// var db *sql.DB
// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "postgres"
// 	password = "golyanglyang"
// 	dbname   = "try"
// )
// // func setupDatabase() (*sql.DB, error) {
// //     // Replace the connection string with your actual PostgreSQL database details
// //     db, err := sql.Open("postgres", "postgres://user:password@host:port/database?sslmode=disable")
// //     if err != nil {
// //         return nil, err
// //     }
// //     return db, nil
// // }
// func main() {
// 	// Create the connection string
// 	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	// Open the database connection
// 	var err error
// 	db, err = sql.Open("postgres", psqlconn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	absPath, err := filepath.Abs("cmd/PlayerStats.csv")
// if err != nil {
// 	log.Fatal(err)
// }
// file, err := os.Open(absPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// Create a new CSV reader
// 	reader := csv.NewReader(file)
// 	controllers.SetDB(db)
// 	for {
// 		// Read each row from the CSV file
// 		row, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		fmt.Println("helloworld")
// 		// Prepare the SQL INSERT statement
// 		stmt, err := db.Prepare("INSERT INTO stats( Player, Nation, Pos, Squad, Comp) VALUES ($1, $2, $3, $4, $5)")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		// Execute the INSERT statement with row data
// 		fmt.Println(row[0], row[1], row[2], row[3], row[4])
// 		_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4])
		
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	router := mux.NewRouter()
// 	routes.RegisterRoutes(router)

// 	// Apply CORS middleware
// 	corsHandler := handlers.CORS(
// 		handlers.AllowedOrigins([]string{"*"}),
// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
// 		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
// 	)

// 	// Use the CORS middleware with the router
// 	router.Use(corsHandler)
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }






// package main

// import (
// 	"log"
// 	"net/http"
// 	"encoding/csv"
// 	"github.com/gorilla/mux"
// 	"github.com/gorilla/handlers"
// 	"football_server/pkg/routes"
// 	"database/sql"
//     _ "github.com/lib/pq"
// 	"fmt"
// 	"os"
// 	"io"
// 	"path/filepath"
// 	"football_server/pkg/controllers"
// )
// var db *sql.DB
// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "postgres"
// 	password = "golyanglyang"
// 	dbname   = "try"
// )
// // func setupDatabase() (*sql.DB, error) {
// //     // Replace the connection string with your actual PostgreSQL database details
// //     db, err := sql.Open("postgres", "postgres://user:password@host:port/database?sslmode=disable")
// //     if err != nil {
// //         return nil, err
// //     }
// //     return db, nil
// // }
// func main() {
// 	// Create the connection string
// 	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	// Open the database connection
// 	var err error
// 	db, err = sql.Open("postgres", psqlconn)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	absPath, err := filepath.Abs("cmd/clubInfo.csv")
// if err != nil {
// 	log.Fatal(err)
// }
// file, err := os.Open(absPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// Create a new CSV reader
// 	reader := csv.NewReader(file)
// 	controllers.SetDB(db)
// 	for {
// 		// Read each row from the CSV file
// 		row, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		} else if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		fmt.Println("helloworld")
// 		// Prepare the SQL INSERT statement
// 		stmt, err := db.Prepare("INSERT INTO team( team_name, abbreviation , hex_code, team_logo,Stadium) VALUES ($1, $2, $3, $4, $5)")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		// Execute the INSERT statement with row data
// 		fmt.Println(row[0], row[1], row[2], row[3], row[4])
// 		_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4])
		
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	router := mux.NewRouter()
// 	routes.RegisterRoutes(router)

// 	// Apply CORS middleware
// 	corsHandler := handlers.CORS(
// 		handlers.AllowedOrigins([]string{"*"}),
// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
// 		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
// 	)

// 	// Use the CORS middleware with the router
// 	router.Use(corsHandler)
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }


package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Player struct {
	Player string `json:"Player"`
	Nation string `json:"Nation"`
	Pos string `json:"Pos"`
	Squad string `json:"Squad"`
	Comp string `json:"Comp"`

};
type Team struct {
	TeamName string `json:"team_name"`
	TeamId int `json:"team_id"`
	Stadium string `json:"Stadium"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "golyanglyang"
	dbname   = "try"
)

func main() {
	http.HandleFunc("/teams/", retrieveTeams)	
	http.HandleFunc("/records/", retrieveRecordByteam)
	http.HandleFunc("/playerStats", retrievePlayerStats)
	

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func retrieveTeams (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Database connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SELECT query
	query := fmt.Sprintf("SELECT team_id, team_name, Stadium from team;")
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Team{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Team
		err := rows.Scan(&player.TeamId, &player.TeamName, &player.Stadium)
		if err != nil {
			log.Println(err)
			continue
		}
		players = append(players, player)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
func retrievePlayerStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Database connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SELECT query
	query := fmt.Sprintf("SELECT Player, Nation, Pos, Squad, Comp FROM stats WHERE Comp = 'Premier League';")
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Player{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.Player, &player.Nation, &player.Pos, &player.Squad, &player.Comp)
		if err != nil {
			log.Println(err)
			continue
		}
		players = append(players, player)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}

func retrieveRecordByteam (w http.ResponseWriter, r *http.Request) {
    // vars := mux.Vars(r)
    id := r.URL.Path[len("/records/"):]
    // if !ok {
    //     fmt.Println("id is missing in parameters")
    // }
    fmt.Println(`id := `, id)
	retrieveRecord(w,r, id)
    //call http://localhost:8080/provisions/someId in your browser
    //Output : id := someId
}
// func retrieveRecordByMikel(w http.ResponseWriter, r *http.Request) {
// 	retrieveRecord(w, r, "Mikel")
// }

// func retrieveRecordByPep(w http.ResponseWriter, r *http.Request) {
// 	retrieveRecord(w, r, "Pep")
// }

func retrieveRecordByArsenal(w http.ResponseWriter , r *http.Request) {
	retrieveRecordteam(w, r, "Arsenal")
}

func retrieveRecordByLiverpool(w http.ResponseWriter , r * http.Request) {
	retrieveRecordteam(w , r , "Liverpool")
}
func retrieveRecordteam(w http.ResponseWriter , r *http.Request, teamName string){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Database connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SELECT query
	query := fmt.Sprintf("SELECT first_name , last_name FROM players where team_id IN (SELECT team_id FROM team  WHERE team_name='%s') ORDER BY first_name ASC;", teamName)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Player{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Player
		// err := rows.Scan(&player.FirstName, &player.LastName)
		if err != nil {
			log.Println(err)
			continue
		}
		players = append(players, player)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)


}


func retrieveRecord(w http.ResponseWriter, r *http.Request, id string) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Database connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open the database connection
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Execute the SELECT query
	query := fmt.Sprintf("SELECT first_name, last_name FROM players WHERE team_id = '%s'", id)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Player{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Player
		// err := rows.Scan(&player.FirstName, &player.LastName)
		if err != nil {
			log.Println(err)
			continue
		}
		players = append(players, player)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the data to JSON
	jsonData, err := json.Marshal(players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}