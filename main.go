package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Represents our board
type ticTacBoard [3][3]rune

func main() {
	// Generate a rand.Seed for random numbers
	rand.Seed(time.Now().UnixNano())

	var playerMove bool // true if player's move
	var whoWon string   // String for who won the game
	var win bool        // true if won. False if tie

	var board ticTacBoard // Create a board instance

	fmt.Println("Starting Game: Board Empty")
	board.displayBoard() // Display the empty board

	if rand.Intn(2) == 0 { // randomly determine who goes first
		playerMove = true // player starts
	} else {
		playerMove = false // computer starts
	}

	for i := 0; i < 9; i++ { // Only can be 9 turns in the game
		if playerMove { // if it is the player's move
			fmt.Println("Player Move: ", i+1)
			board.Player() // handle the player's turn
			playerMove = false
		} else { // if it is the computer's move
			fmt.Println("Computer Move: ", i+1)
			time.Sleep(time.Second) // sleep for 1 second to simulate thinking
			board.Computer()        // handle the computer's turn
			playerMove = true       // now it is the player's turn

		}
		if whoWon, win = board.Check(); win { // Check who won. If someone won, break the loop. If not, continue
			break
		}
		board.displayBoard() // display the board for the next turn
	}
	fmt.Printf("*****%v won*****\nFinal Board View:\n", whoWon) // prompt about who won
	board.displayBoard()                                        // display the winning board

}

// displayBoard Displays the board in the console
func (t *ticTacBoard) displayBoard() {
	fmt.Print("------------") // top line
	for i := 0; i < 3; i++ {
		fmt.Print("\n|") // Side wall
		for j := 0; j < 3; j++ {
			fmt.Printf(" %c |", t[i][j]) // value in the specified slot, and wall
		}
		fmt.Print("\n------------") // bottom wall
	}
	fmt.Print("\n") // add a new line below
}

// Player handles the player's turn
func (t *ticTacBoard) Player() {
	var x, y int // x and y coords for turn
	fmt.Println("Enter the Row(1-3) and the column(1-3) positions: ")
	if _, err := fmt.Scan(&x, &y); err == nil {
		x--                                                             // minus one for 2d array vals
		y--                                                             // minus one for 2d array vals
		if (x >= 0 && x <= 2) && (y >= 0 && y <= 2) && (t[x][y] == 0) { // make sure it is a valid move
			t[x][y] = 'X' // if valid make the move happen
		} else {
			fmt.Println("Invalid input or position not empty. Try again") // if not valid redo requesting player's turn after prompting
			t.Player()
		}
	} else {
		fmt.Println("Invalid input or position not empty. Try again") // if not valid redo requesting player's turn after prompting
		t.Player()
	}
}

// Computer handles the computer's turn
func (t *ticTacBoard) Computer() {
	var x, y int // x and y coords for turn in 2d array
	for {
		x = rand.Intn(2) // computer x
		y = rand.Intn(2) // computer y
		if t[x][y] == 0 {
			t[x][y] = 'O'
			break
		}
	}
}

// Check checks who won the game
func (t *ticTacBoard) Check() (string, bool) {
	for i := 0; i < 3; i++ {
		if (rune(t[i][0]) == 'X') && (rune(t[i][1]) == 'X') && (rune(t[i][2]) == 'X') { // horizontal Player
			return "Player", true
		} else if (rune(t[i][0]) == 'O') && (rune(t[i][1]) == 'O') && (rune(t[i][2]) == 'O') { // horizontal Computer
			return "Computer", true
		} else if (rune(t[0][i]) == 'X') && (rune(t[1][i]) == 'X') && (rune(t[2][i]) == 'X') { // Vertical player
			return "Player", true
		} else if (rune(t[0][i]) == 'O') && (rune(t[1][i]) == 'O') && (rune(t[2][i]) == 'O') { // Vertical Computer
			return "Computer", true
		}
	}

	if ((rune(t[0][0]) == 'X') && (rune(t[1][1]) == 'X') && (rune(t[2][2]) == 'X')) || ((rune(t[0][2]) == 'X') && (rune(t[1][1]) == 'X') && (rune(t[2][0]) == 'X')) {
		// Diagonal Player
		return "Player", true
	} else if ((rune(t[0][0]) == 'O') && (rune(t[1][1]) == 'O') && (rune(t[2][2]) == 'O')) || ((rune(t[0][2]) == 'O') && (rune(t[1][1]) == 'O') && (rune(t[2][0]) == 'O')) {
		// Diagonal computer
		return "Computer", true
	}
	// Tie
	return "Nobody Wins", false
}
