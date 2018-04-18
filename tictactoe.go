package tictactoe

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var gameHistory []int

//BuildGameMap generates key-value pairs to represent positions in the tic-tac-toe board
func BuildGameMap() map[int]string {
	gameMap := make(map[int]string)
	for i := 1; i < 10; i++ {
		gameMap[i] = strconv.Itoa(i)
	}
	return gameMap
}

//MarkerSelection allow the user to choose whether to play using "X" or "O"
func MarkerSelection() (string, string) {

	var selectedMark string

	fmt.Print("\nWelcome, please hit X or O to choose your marker: ")
	for {
		fmt.Scan(&selectedMark)
		switch strings.ToUpper(selectedMark) {
		case "X":
			humanMark := "X"
			computerMark := "O"
			return humanMark, computerMark
		case "O":
			humanMark := "O"
			computerMark := "X"
			return humanMark, computerMark
		default:
			fmt.Print("\nInvalid option, you must choose between X or O. Please try again: ")
		}
	}
}

//WhoStarts allow user to select who makes the first move by choosing "C" for computer or "H" for human
func WhoStarts() string {
	var firstPlayer string
	for {
		fmt.Print("\nWho should play first? Choose (H)uman or (C)omputer: ")
		fmt.Scan(&firstPlayer)
		firstPlayer = strings.ToUpper(firstPlayer)
		if firstPlayer == "H" {
			return "H"
		} else if firstPlayer == "C" {
			return "C"
		} else {
			fmt.Println("\nINVALID ENTRY.  You must choose H or C.\nPlease try again...")
		}
	}
}

//AskUser gets input from user
func AskUser() int {
	for {
		var userEntry int
		fmt.Print("Please enter an available position: ")
		fmt.Scan(&userEntry)
		if validateEntry(userEntry) {
			return userEntry
		}
		fmt.Println("\nINVALID ENTRY, you must choose a valid position.")
	}
}

//ComputerMove generates a random number as the computer move
func ComputerMove() int {
	for {
		computerMove := rand.Intn(9) + 1
		if validateEntry(computerMove) {
			return computerMove
		}
	}
}

func validateEntry(userEntry int) bool {
	for {
		if userEntry > 0 && userEntry < 10 && !contains(gameHistory, userEntry) {
			gameHistory = append(gameHistory, userEntry)
			return true
		}
		return false
	}
}

func contains(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}

//PlayMove actually puts userEntry in the position array
func PlayMove(gameMap map[int]string, turnMove int, mark string) {
	for k := range gameMap {
		if k == turnMove {
			gameMap[k] = mark
			DrawBoard(gameMap)
		}
	}
}

// DrawBoard displays tic-tac-toe game board on the screen
func DrawBoard(gameMap map[int]string) {
	fmt.Println("\n " + gameMap[1] + " | " + gameMap[2] + " | " + gameMap[3] + "  ")
	fmt.Println("---+---+---")
	fmt.Println(" " + gameMap[4] + " | " + gameMap[5] + " | " + gameMap[6] + "  ")
	fmt.Println("---+---+---")
	fmt.Println(" " + gameMap[7] + " | " + gameMap[8] + " | " + gameMap[9] + "  ")
	fmt.Println(" ")
	fmt.Println("\nGame history: ", gameHistory)
}

//CheckWinner will check all possible winning positions
func CheckWinner(gameMap map[int]string) {
	if gameMap[1] == "X" && gameMap[2] == "X" && gameMap[3] == "X" ||
		gameMap[4] == "X" && gameMap[5] == "X" && gameMap[6] == "X" ||
		gameMap[7] == "X" && gameMap[8] == "X" && gameMap[9] == "X" ||
		gameMap[1] == "X" && gameMap[4] == "X" && gameMap[7] == "X" ||
		gameMap[2] == "X" && gameMap[5] == "X" && gameMap[8] == "X" ||
		gameMap[3] == "X" && gameMap[6] == "X" && gameMap[9] == "X" ||
		gameMap[1] == "X" && gameMap[5] == "X" && gameMap[9] == "X" ||
		gameMap[3] == "X" && gameMap[5] == "X" && gameMap[7] == "X" ||
		gameMap[1] == "O" && gameMap[2] == "O" && gameMap[3] == "O" ||
		gameMap[4] == "O" && gameMap[5] == "O" && gameMap[6] == "O" ||
		gameMap[7] == "O" && gameMap[8] == "O" && gameMap[9] == "O" ||
		gameMap[1] == "O" && gameMap[4] == "O" && gameMap[7] == "O" ||
		gameMap[2] == "O" && gameMap[5] == "O" && gameMap[8] == "O" ||
		gameMap[3] == "O" && gameMap[6] == "O" && gameMap[9] == "O" ||
		gameMap[1] == "O" && gameMap[5] == "O" && gameMap[9] == "O" ||
		gameMap[3] == "O" && gameMap[5] == "O" && gameMap[7] == "O" {
		fmt.Println("\nWe got a WINNER")
		goodByeMessage()
	} else if len(gameHistory) == 9 {
		fmt.Println("\nIT'S A TIED GAME")
		goodByeMessage()
	}
}

func goodByeMessage() {
	fmt.Println("\nThanks for playing, come back soon...")
	os.Exit(0)
}
