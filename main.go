package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
	"net/http"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "golyanglyang"
    dbname   = "try"
)


 
func main() {
        // connection string
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
        // open database
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
     
        // close database
    defer db.Close()
 
        // check db
    err = db.Ping()
    CheckError(err)
 
    fmt.Println("Connected!")
	rows, err := db.Query(`SELECT "first_name", "last_name" FROM "players"`)
CheckError(err)
 
defer rows.Close()
for rows.Next() {
    var first_name string
    var last_name string
 
    err = rows.Scan(&first_name, &last_name)
    CheckError(err)
 
    fmt.Println(first_name, last_name)
}
 
CheckError(err)
	
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}