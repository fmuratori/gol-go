package core

import (
	"fmt"
	"strings"
)

func CreateMatrix(width int, height int) [][]int {
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	return matrix
}

func MatrixToString(board [][]int) string {
	width := len(board[0])
	height := len(board)

	stringBoard := ""

	lineDelimiter := strings.Repeat("+-", width) + "+\n"
	stringBoard += lineDelimiter

	for i := 0; i < height; i++ {

		lineValues := ""
		for j := 0; j < width; j++ {
			lineValues += fmt.Sprintf("|%d", board[i][j])
		}
		lineValues += "|\n"

		stringBoard += lineValues
		stringBoard += lineDelimiter
	}

	return stringBoard
}

func MatrixDeepCopy(src [][]int) [][]int {
	dest := make([][]int, len(src))
	for i := range src {
		// Create a new slice for each row
		dest[i] = make([]int, len(src[i]))
		// Copy the content of the original row to the new row
		copy(dest[i], src[i])
	}
	return dest
}
