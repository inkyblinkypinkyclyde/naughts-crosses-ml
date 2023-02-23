package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	// _ "github.com/lib/pq"
)

var grid = [10]string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

type GameTurn struct {
	turn     int
	position int
	player   string
}

var turn int = 0

var gameTurns = []GameTurn{}

func main() {
	generateTable()
	for {
		showGrid()
		playerTurn("x")
		// computerTurn("x")
		if winCheck("x") == "x" {
			fmt.Println("You win!")
			break
		}
		if staleMateCHeck() {
			fmt.Println("Stalemate!")
			break
		}
		computerTurn("o")
		if winCheck("o") == "o" {
			fmt.Println("You lose!")
			break
		}
		if staleMateCHeck() {
			fmt.Println("Stalemate!")
			break
		}
	}
	fmt.Println(gameTurns)
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
	turn++
	fmt.Println("Computer turn number ", turn)
	for {
		c := rng()
		if grid[c] != "x" && grid[c] != "o" {
			grid[c] = ox
			// fmt.Println("Computer chose: ", c)
			gameTurns = append(gameTurns, GameTurn{turn, c, ox})
			break
		}
	}
}

func playerTurn(ox string) {
	turn++
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
			gameTurns = append(gameTurns, GameTurn{turn, i, ox})
			break
		}
		fmt.Println("Invalid selection")
	}
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

func staleMateCHeck() bool {
	if turn == 9 {
		return true
	}
	return false
}
