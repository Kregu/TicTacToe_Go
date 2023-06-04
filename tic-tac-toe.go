package main

import (
	"fmt"
	"time"
)

const X byte = 'X'
const O byte = 'O'
const EMPTY byte = ' '
const TIE byte = 'T'
const NoOne byte = 'N'

func main() {
	for {
		fmt.Println()
		fmt.Println("Welcome to Tic-Tac-Toe")
		fmt.Println()
		const NumSquares int = 9
		board := [NumSquares]byte{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}

		human := humanPiece()
		computer := opponent(human)
		game(board, human, computer)

		repeat := askYesNo("Do you want repeat Game?")
		if repeat == "y" {
			continue
		} else {
			break
		}
	}
	return
}

func game(board [9]byte, human, computer byte) {
	var move int
	turn := X
	displayBoard(&board)

	for winner(board) == NoOne {

		if turn == human {
			move = humanMove(board)
			board[move] = human
		} else {
			move = computerMove(board, computer)
			board[move] = computer
		}
		displayBoard(&board)
		if turn == human {
			duration := time.Second * 2
			time.Sleep(duration)
		}
		turn = opponent(turn)

	}
	announceWinner(winner(board), computer, human)
	return
}

func humanPiece() byte {
	goFirst := askYesNo("Do you require the first move?")
	if goFirst == "y" {
		fmt.Println("\nThen take the first move. You will need it.")
		return X
	} else {
		fmt.Println("\nYou bravery will be you undoing... I will go first.")
		return O
	}
}

func askYesNo(question string) string {
	var response string

	for {
		fmt.Printf("%s (y/n): ", question)
		fmt.Scanf("%s", &response)

		if response == "y" || response == "n" {
			break
		}
	}
	return response
}

func opponent(human byte) byte {
	if human == X {
		return O
	} else {
		return X
	}
}

func displayBoard(board *[9]byte) {
	instruction()
	fmt.Printf("\n\t%c | %c | %c", board[0], board[1], board[2])
	fmt.Printf("\n\t---------")
	fmt.Printf("\n\t%c | %c | %c", board[3], board[4], board[5])
	fmt.Printf("\n\t---------")
	fmt.Printf("\n\t%c | %c | %c", board[6], board[7], board[8])
	fmt.Println()
	fmt.Println()
}

func winner(board [9]byte) byte {
	WinnerRows := [8][3]int{{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6}}

	const TotalRows = 8

	for row := 0; row < TotalRows; row++ {
		if (board[WinnerRows[row][0]] != EMPTY) &&
			board[WinnerRows[row][0]] == board[WinnerRows[row][1]] &&
			board[WinnerRows[row][1]] == board[WinnerRows[row][2]] {
			return board[WinnerRows[row][0]]
		}
	}
	count := 0
	for _, value := range board {
		if value == EMPTY {
			count++
		}
	}
	if count == 0 {
		return TIE
	}
	return NoOne
}

func instruction() {
	fmt.Println()
	fmt.Println("Make your move know by entering a number. 0 - 8")
	fmt.Println()
	fmt.Println("\t0 | 1 | 2")
	fmt.Println("\t---------")
	fmt.Println("\t3 | 4 | 5")
	fmt.Println("\t---------")
	fmt.Println("\t6 | 7 | 8")
	fmt.Println()
}

func announceWinner(winner, computer, human byte) {
	if winner == computer {
		fmt.Printf("%c's won!\n", winner)
		fmt.Printf("You're pathetic, Human.\n")
		fmt.Printf("I'm better than you in everything!\n\n")
	} else if winner == human {
		fmt.Printf("%c's won!\n", winner)
	} else
	{
		fmt.Printf("It's a tie.\n")
	}
}

func humanMove(board [9]byte) int {
	move := askNumber("Where will you move?", len(board)-1, 0)
	for !isLegal(move, board) {
		fmt.Printf("\nThat square is already occupied.\n")
		move = askNumber("Where will you move?", len(board)-1, 0)
	}
	fmt.Printf("Fine... %d.", move)
	return move
}

func askNumber(question string, high, low int) int {
	var number int
	for {
		fmt.Printf("%s  (%d - %d): ", question, low, high)
		fmt.Scan(&number)
		if number >= low && number <= high {
			break
		}
	}
	return number
}

func isLegal(move int, board [9]byte) bool {
	return board[move] == EMPTY
}

func computerMove(board [9]byte, computer byte) int {
	move := 0
	found := false

	for !found && move < len(board) {
		if isLegal(move, board) {
			board[move] = computer
			found = winner(board) == computer
			board[move] = EMPTY
		}
		if !found {
			move++
		}
	}

	if !found {
		move = 0
		human := opponent(computer)
		for !found && move < len(board) {
			if isLegal(move, board) {
				board[move] = human
				found = winner(board) == human
				board[move] = EMPTY
			}
			if !found {
				move++
			}
		}
	}

	if !found {
		move = 0
		i := 0

		BestMoves := []int{4, 0, 2, 6, 8, 1, 3, 5, 7}
		for !found && i < len(board) {
			move = BestMoves[i]

			if isLegal(move, board) {
				found = true
				break
			}
			i++
		}
	}
	fmt.Printf("I shall take square number %d.", move)
	return move
}
