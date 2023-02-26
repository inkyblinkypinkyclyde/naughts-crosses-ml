package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var grid = [10]string{}

type SessionOptions struct {
	mode  string
	loops int
	table string
}

var gameTurnsStrings = make([]string, 10)
var turn int

func main() {
	sessionOptions := getSessionOptions()

	db := generateTable(sessionOptions.table)
	for i := 0; i < sessionOptions.loops; i++ {
		fmt.Println("Starting game #" + strconv.Itoa(i))
		resetGame()
		for {
			if sessionOptions.mode == "h" {
				showGrid()
				playerTurn("x")
			} else {
				computerTurn("x")
			}
			if winCheck("x") == "x" {
				fmt.Println("You win!")
				endGame("x", sessionOptions.table, db)
				break
			}
			if staleMateCHeck() {
				fmt.Println("Stalemate!")
				endGame("s", sessionOptions.table, db)
				break
			}
			computerTurn("o")
			if winCheck("o") == "o" {
				endGame("o", sessionOptions.table, db)
				break
			}
			if staleMateCHeck() {
				endGame("s", sessionOptions.table, db)
				break
			}
		}
	}

}

func showGrid() {
	fmt.Println("=========")
	fmt.Printf("%s | %s | %s\n", grid[1], grid[2], grid[3])
	fmt.Printf("%s | %s | %s\n", grid[4], grid[5], grid[6])
	fmt.Printf("%s | %s | %s\n", grid[7], grid[8], grid[9])
	fmt.Println("=========")
}

func rng() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9) + 1
}

func computerTurn(ox string) {
	// fmt.Println("Computer turn number ", turn)
	for {
		c := rng()
		if grid[c] != "x" && grid[c] != "o" {
			grid[c] = ox
			// fmt.Println("Computer chose: ", c)
			var newReportString string
			if turn == 0 {
				newReportString = gameTurnsStrings[turn] + ox + strconv.Itoa(c)
			} else {
				newReportString = gameTurnsStrings[turn-1] + ox + strconv.Itoa(c)
			}
			gameTurnsStrings[turn] = newReportString
			break
		}
	}
	turn++
}

func playerTurn(ox string) {
	for {
		fmt.Println("Enter a selection")
		var c string
		fmt.Scanln(&c)
		i, e := strconv.Atoi(c)
		if e != nil {
			fmt.Println("Invalid selection")
			break
		}
		if i > 0 && i < 10 {
			grid[i] = ox
			break
		}
		fmt.Println("Invalid selection")
	}
	turn++
}

func winCheck(p string) string {
	if grid[1] == p && grid[2] == p && grid[3] == p {
		return p
	}
	if grid[4] == p && grid[5] == p && grid[6] == p {
		return p
	}
	if grid[7] == p && grid[8] == p && grid[9] == p {
		return p
	}
	if grid[1] == p && grid[4] == p && grid[7] == p {
		return p
	}
	if grid[2] == p && grid[5] == p && grid[8] == p {
		return p
	}
	if grid[3] == p && grid[6] == p && grid[9] == p {
		return p
	}
	if grid[1] == p && grid[5] == p && grid[9] == p {
		return p
	}
	if grid[3] == p && grid[5] == p && grid[7] == p {
		return p
	}
	return ""
}

func staleMateCHeck() bool { //returns true if turn is 9
	return turn == 9
}

func resetGame() {
	// Reset turns
	turn = 0
	// Reset Grid
	for i, _ := range grid {
		if i == 0 {
			grid[i] = ""
		} else {
			grid[i] = strconv.Itoa(i)
		}
	}
	// Reset Game turns strings
	for i, _ := range gameTurnsStrings {
		gameTurnsStrings[i] = ""
	}
}

func endGame(player string, table string, db *sql.DB) {
	gameTurnsStrings[9] = player
	dumpGameTurns(table, gameTurnsStrings, db)
	fmt.Println(gameTurnsStrings)
	resetGame()
}

func getSessionOptions() SessionOptions {
	var mode string
	fmt.Println("Play (h)uman against computer or (c)omputer against computer?")
	fmt.Scanln(&mode)
	numberOfLoops := "1"
	var table string
	if mode == "c" {
		fmt.Println("Enter number of loops:")
		fmt.Scanln(&numberOfLoops)
		table = "random"
	} else {
		table = "human"
	}
	numberOfLoopsInt, _ := strconv.Atoi(numberOfLoops)
	newSessionOptions := SessionOptions{mode, numberOfLoopsInt, table}
	return newSessionOptions
}
