package main

import "fmt"

func main() {
	gameBoard := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	isGameOver := false
	playerTurn := 1
	s := `*****************************************
* 		                        *
* Welcome to Go Tic Tac Toe CLI Game.   *
*                                       *
* Please enter a number under 10.       *
*                                       *
* Enter 9 to Quit the Game              *
*                                       *
*****************************************

	`
	fmt.Print(s)
	for !isGameOver {
		drawBoard(gameBoard)

		currentPlayer := playerTurn%2 + 1
		if currentPlayer == 1 {
			fmt.Println("Player 1 turn ")
		} else {
			fmt.Println("Player 2 turn ")
		}

		move := askForMove()
		if move == 9 {
			fmt.Println("Quitting Game. Bye")
			return
		}

		gameBoard = executePlayerMove(move, currentPlayer, gameBoard)
		w := checkForWin(gameBoard)
		if w > 0 {
			fmt.Printf("\n Player %d wins!\n\n", w)
			isGameOver = true
		} else if playerTurn == 9 { // 9 moves but no winner
			fmt.Printf("We have a tie !!!\n\n")
			isGameOver = true
		} else {
			playerTurn++
		}
	}
}

func drawBoard(b [9]int) {
	fmt.Printf("\n")
	for i, v := range b {
		if v == 0 {
			fmt.Printf("%d", i)
		} else if v == 1 {
			fmt.Printf("X")
		} else if v == 10 {
			fmt.Printf("O")
		}
		// draw 3x3 grid
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("  |  ")
		}
	}
	fmt.Printf("\n")
}

func askForMove() int {
	fmt.Print("What's your next move? ")
	var moveValue int
	fmt.Scan(&moveValue)
	return moveValue
}

func executePlayerMove(move int, currentPlayer int, gameBoard [9]int) [9]int {
	if move > 9 {
		fmt.Println("Please enter a number under 10.")
		move = askForMove()
	}
	if gameBoard[move] != 0 {
		fmt.Println("Please select am empty space.")
		move = askForMove()
		gameBoard = executePlayerMove(move, currentPlayer, gameBoard)
	} else {
		if currentPlayer == 1 {
			gameBoard[move] = 1
		} else if currentPlayer == 2 {
			gameBoard[move] = 10
		}
	}
	return gameBoard
}

func checkForWin(b [9]int) int {
	winningCombination := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0} //there are 8 possible winning combinations
	winningCombination[0] = b[0] + b[1] + b[2]
	winningCombination[1] = b[3] + b[4] + b[5]
	winningCombination[2] = b[6] + b[7] + b[8]
	winningCombination[3] = b[0] + b[3] + b[6]
	winningCombination[4] = b[1] + b[4] + b[7]
	winningCombination[5] = b[2] + b[5] + b[8]
	winningCombination[6] = b[0] + b[4] + b[8]
	winningCombination[7] = b[2] + b[4] + b[6]

	for _, v := range winningCombination {
		if v == 3 { //see line 80
			return 1
		} else if v == 30 {
			return 2
		}
	}
	return 0
}
