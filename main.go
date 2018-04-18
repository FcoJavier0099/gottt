package main

import (
	"tictactoe"
)

func main() {
	gameMap := tictactoe.BuildGameMap()
	humanMark, computerMark := tictactoe.MarkerSelection()

	if tictactoe.WhoStarts() == "C" {
		for {
			move := tictactoe.ComputerMove()
			mark := computerMark
			executeGameLogic(gameMap, move, mark)
			move = tictactoe.AskUser()
			mark = humanMark
			executeGameLogic(gameMap, move, mark)
		}
	} else {
		for {
			tictactoe.DrawBoard(gameMap)
			move := tictactoe.AskUser()
			mark := humanMark
			executeGameLogic(gameMap, move, mark)
			move = tictactoe.ComputerMove()
			mark = computerMark
			executeGameLogic(gameMap, move, mark)
		}
	}

}

func executeGameLogic(gameMap map[int]string, move int, mark string) {
	tictactoe.PlayMove(gameMap, move, mark)
	tictactoe.CheckWinner(gameMap)
}
