package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var connStr = "postgresql://localhost/naughts_and_crosses?user=postgres&password=postgres&sslmode=disable&port=5434" // Our connection string

func generateTable(tableName string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("CREATE TABLE " + tableName + " (game_id SERIAL PRIMARY KEY,turn0 VARCHAR(30),turn1 VARCHAR(30),turn2 VARCHAR(30),turn3 VARCHAR(30),turn4 VARCHAR(30),turn5 VARCHAR(30),turn6 VARCHAR(30),turn7 VARCHAR(30),turn8 VARCHAR(30),winner VARCHAR(2));")
	if err != nil {
		fmt.Println(err)
	}
	return db

}

func dumpGameTurns(tableName string, turns []string, db *sql.DB) {
	sqlString := "INSERT INTO " + tableName + " (turn0, turn1, turn2, turn3, turn4, turn5, turn6, turn7, turn8, winner) VALUES ("
	for turn, turnString := range turns {
		sqlString = sqlString + "'" + turnString + "'"
		if turn == 9 {
			sqlString = sqlString + ")"
		} else {
			sqlString = sqlString + ", "
		}
		// fmt.Println(sqlString)
	}
	_, err := db.Exec(sqlString)
	if err != nil {
		fmt.Println(err)
	}

}
