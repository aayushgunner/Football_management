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
// 	absPath, err := filepath.Abs("cmd/cleaned_players.csv")
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
// 		stmt, err := db.Prepare("INSERT INTO stats( first_name, second_name, goals, assists, mins_played,red_cards,yellow_cards) VALUES ($1, $2, $3, $4, $5, $6, $7)")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		// Execute the INSERT statement with row data
// 		fmt.Println(row[0], row[1], row[2], row[3], row[4], row[5], row[6])
// 		_, err = stmt.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6])
		
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
// 		stmt, err := db.Prepare("INSERT INTO moreStats( Player, Nation, Pos, Squad,Comp) VALUES ($1, $2, $3, $4, $5)")
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
// 	absPath, err := filepath.Abs("cmd/understat_player.csv")
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
// 		stmt, err := db.Prepare("INSERT INTO player( player_link_id, player_name, team_name) VALUES ($1, $2, $3)")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		// Execute the INSERT statement with row data
// 		fmt.Println(row[0], row[1], row[2])
// 		_, err = stmt.Exec(row[0], row[1], row[2])
		
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
// 	absPath, err := filepath.Abs("cmd/player_stats.csv")
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
// 		stmt, err := db.Prepare("INSERT INTO stats(games, Mins_played , goal,  assists, shots,key_passes,yellow_cards,red_cards,position,player_id) VALUES ($1, $2, $3, $4,$5, $6, $7,$8, $9,$10)")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
	
// 		// Execute the INSERT statement with row data
// 		fmt.Println(row[0], row[1], row[2],row[3],row[4],row[5],row[6],row[7],row[8],row[9])
// 		_, err = stmt.Exec(row[0], row[1], row[2],row[3],row[4],row[5],row[6],row[7],row[8],row[9])
		
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
// 		stmt, err := db.Prepare("INSERT INTO teams( Team_name, Abbreviation , hex_code, logo, Stadium) VALUES ($1, $2, $3, $4, $5)")
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
	PlayerId int `json:"player_id"`
	Player_name string `json:"player_name"`
    Player_team string `json:"team_name"`	

};
type Team struct {
	TeamName string `json:"team_name"`
	TeamId int `json:"team_id"`
	Stadium string `json:"Stadium"`
}

type Stats struct {
	Games        int    `json:"games"`
	MinsPlayed   int    `json:"mins_played"`
	Goal         int    `json:"goal"`
	Assists      int    `json:"assists"`
	Shots        int    `json:"shots"`
	KeyPasses    int    `json:"key_passes"`
	YellowCards  int    `json:"yellow_cards"`
	RedCards     int    `json:"red_cards"`
	Position     string `json:"position"`
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
	http.HandleFunc("/players/", retrievePlayers)	
	http.HandleFunc("/stats/", retrieveRecordById)	
	
	

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
	query := fmt.Sprintf("SELECT team_id, team_name, Stadium from teams;")
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
func retrievePlayers(w http.ResponseWriter, r *http.Request) {
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
	query := fmt.Sprintf("SELECT player_id, player_name, team_name FROM player ORDER by team_id ASC;")
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
		err := rows.Scan(&player.PlayerId,&player.Player_name, &player.Player_team)
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

func retrieveRecordById(w http.ResponseWriter, r *http.Request) {
	// Extract the player ID from the request URL
	id := r.URL.Path[len("/stats/"):]
	fmt.Println(`id := `, id)

	// Set the necessary HTTP headers
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
	query := fmt.Sprintf("SELECT games, Mins_played, goal, assists, shots, key_passes, yellow_cards, red_cards, position FROM stats WHERE player_id = %s LIMIT 1", id)
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice to store the retrieved records
	players := []Stats{}

	// Iterate over the rows and populate the slice
	for rows.Next() {
		var player Stats
		err := rows.Scan(&player.Games, &player.MinsPlayed, &player.Goal, &player.Assists, &player.Shots, &player.KeyPasses, &player.YellowCards, &player.RedCards, &player.Position)
		if err != nil {
			log.Println(err)
			continue
		}
		players = append(players, player)
		fmt.Println(player)
	}

	// Check for any errors encountered during row iteration
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(`id := `, players)

	// Check if any players were retrieved
	if len(players) == 0 {
		http.NotFound(w, r)
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


// func retrieveStatsbyId(w http.ResponseWriter, r *http.Request, id string) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	// Database connection string
// 	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	// Open the database connection
// 	db, err := sql.Open("postgres", psqlconn)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer db.Close()

// 	// Execute the SELECT query
// 	query := fmt.Sprintf("SELECT games, Mins_played, goal, assists, shots,key_passes,yellow_cards,red_cards,position FROM stats WHERE player_id = %s",id)
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	// Create a slice to store the retrieved records
// 	players := []Stats{}

// 	// Iterate over the rows and populate the slice
// 	for rows.Next() {
// 		var player Stats
// 		err := rows.Scan(&player.Games, &player.MinsPlayed, &player.Goal, &player.Assists, &player.Shots, &player.KeyPasses, &player.YellowCards, &player.RedCards, &player.Position)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}
// 		players = append(players, player)
// 	}

// 	// Check for any errors encountered during row iteration
// 	if err := rows.Err(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Println(`id := `, players)

// 	// Convert the data to JSON
// 	jsonData, err := json.Marshal(players)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the Content-Type header to application/json
// 	w.Header().Set("Content-Type", "application/json")


// 	// Write the JSON response
// 	w.Write(jsonData)
// }


