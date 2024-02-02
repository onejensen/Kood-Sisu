package sprint

import (
	"strconv"
	"strings"
)

const boardSize = 8

func EightQueensSolver() string {
	var solutions []string
	board := make([]int, boardSize)
	solveQueens(0, board, &solutions)
	return strings.Join(solutions, "\n")
}

func solveQueens(row int, board []int, solutions *[]string) {
	if row == boardSize {
		*solutions = append(*solutions, formatSolution(board))
		return
	}

	for col := 0; col < boardSize; col++ {
		if isSafe(row, col, board) {
			board[row] = col
			solveQueens(row+1, board, solutions)
		}
	}
}

func isSafe(row, col int, board []int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || abs(board[i]-col) == abs(i-row) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func formatSolution(board []int) string {
	var solutionBuilder strings.Builder
	for _, col := range board {
		solutionBuilder.WriteString(strconv.Itoa(col + 1))
	}
	return solutionBuilder.String()
}
