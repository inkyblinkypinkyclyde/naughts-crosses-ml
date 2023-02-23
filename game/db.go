package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connStr = "postgresql://localhost/naughts_and_crosses?user=postgres&password=postgres&sslmode=disable&port=5434" // Our connection string

func generateTable() {
	db, err := sql.Open("postgres", connStr)
	if err != nil { // If there is an error
		fmt.Println(err)
	}
	fmt.Println("Connected to db")
	_, err = db.Exec("CREATE TABLE games (game_id SERIAL PRIMARY KEY,turn1 VARCHAR(30),turn2 VARCHAR(30),turn3 VARCHAR(30),turn4 VARCHAR(30),turn5 VARCHAR(30),turn6 VARCHAR(30),turn7 VARCHAR(30),turn8 VARCHAR(30),turn9 VARCHAR(30),winner VARCHAR(2));")
	if err != nil {
		fmt.Println(err)
	}

}

// func dumpGameTurns(turns []string) {
// 	db, err := sql.Open("postgres", connStr) // Open a database connection
// 	if err != nil {                          // If there is an error
// 		fmt.Println(err)
// 	}
// 	_, err = db.Exec() // add some sql here
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, v := range gameTurns {
// 		fmt.Println(v)
// 	}
// }
