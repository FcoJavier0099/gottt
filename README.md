# gottt
package main

import (
	"tictactoe"
)

func main() {
	gameMap := tictactoe.BuildGameMap()
	humanMark, computerMark := tictactoe.MarkerSelection()

	for {
		computerMove := tictactoe.ComputerMove()
		tictactoe.PlayMove(gameMap, computerMove, computerMark)
		tictactoe.CheckWinner(gameMap)
		humanMove := tictactoe.AskUser()
		tictactoe.PlayMove(gameMap, humanMove, humanMark)
		tictactoe.CheckWinner(gameMap)
	}
}
