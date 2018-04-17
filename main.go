package main

import (
	"tictactoe"
)

func main() {
	gameMap := tictactoe.BuildGameMap()

	var mark = [2]string{"X", "O"}

	for {
		computerMove := tictactoe.ComputerMove()
		tictactoe.PlayMove(gameMap, computerMove, mark[0])
		// tictactoe.DisplayGameHistory()
		tictactoe.CheckWinner(gameMap)
		playerMove := tictactoe.AskUser()
		tictactoe.PlayMove(gameMap, playerMove, mark[1])
		// tictactoe.DisplayGameHistory()
		tictactoe.CheckWinner(gameMap)
	}
}
