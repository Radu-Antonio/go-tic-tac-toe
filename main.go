package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var board [3][3]string
var buttons [3][3]*widget.Button
var currentPlayer = "X"
var aiEnabled = false
var myWindow fyne.Window

func main() {
	myApp := app.New()
	myWindow = myApp.NewWindow("Tic Tac Toe")

	showMenu()

	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()
}

func showMenu() {
	btn2Players := widget.NewButton("2 Players", func() {
		aiEnabled = false
		startGame()
	})

	btnVsAI := widget.NewButton("Vs AI", func() {
		aiEnabled = true
		startGame()
	})

	menu := container.NewVBox(
		widget.NewLabel("Choose game mode:"),
		btn2Players,
		btnVsAI,
	)

	myWindow.SetContent(menu)
}

func startGame() {
	currentPlayer = "X"
	resetBoard()

	grid := container.NewGridWithRows(3)

	for i := 0; i < 3; i++ {
		row := container.NewGridWithColumns(3)
		for j := 0; j < 3; j++ {
			btn := widget.NewButton("", nil)

			btn.OnTapped = func() {
				if board[i][j] == "" && (!aiEnabled || currentPlayer == "X") {
					board[i][j] = currentPlayer
					btn.SetText(currentPlayer)
					if checkWin(currentPlayer) {
						dialog.ShowInformation("Game over", currentPlayer+" WON", myWindow)
						resetBoard()
						updateButtons()
						return
					}
					if isBoardFull() {
						dialog.ShowInformation("Game over", "DRAW", myWindow)
						resetBoard()
						updateButtons()
						return
					}
					if currentPlayer == "X" {
						currentPlayer = "O"
					} else {
						currentPlayer = "X"
					}

					if aiEnabled && currentPlayer == "O" {
						aiMove()
					}
				}
			}

			buttons[i][j] = btn
			row.Add(btn)
		}
		grid.Add(row)
	}

	myWindow.SetContent(grid)
}

func aiMove() {
	bestScore := -1000
	var move [2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				board[i][j] = "O"
				score := minimax(0, false)
				board[i][j] = ""
				if score > bestScore {
					bestScore = score
					move = [2]int{i, j}
				}
			}
		}
	}

	i, j := move[0], move[1]
	board[i][j] = "O"
	buttons[i][j].SetText("O")

	if checkWin("O") {
		dialog.ShowInformation("Joc terminat", "O (AI) a câștigat!", myWindow)
		resetBoard()
		updateButtons()
		return
	}

	if isBoardFull() {
		dialog.ShowInformation("Remiză", "Nu există câștigător!", myWindow)
		resetBoard()
		updateButtons()
		return
	}

	currentPlayer = "X"
}

func minimax(depth int, isMaximizing bool) int {
	if checkWin("O") {
		return 10 - depth
	}
	if checkWin("X") {
		return depth - 10
	}
	if isBoardFull() {
		return 0
	}

	if isMaximizing {
		bestScore := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == "" {
					board[i][j] = "O"
					score := minimax(depth+1, false)
					board[i][j] = ""
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == "" {
					board[i][j] = "X"
					score := minimax(depth+1, true)
					board[i][j] = ""
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
}

func resetBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = ""
		}
	}
	currentPlayer = "X"
}

func updateButtons() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			buttons[i][j].SetText(board[i][j])
		}
	}
}

func checkWin(player string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func isBoardFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				return false
			}
		}
	}
	return true
}
