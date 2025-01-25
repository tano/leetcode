package array

import (
	"fmt"
	"strconv"
)

func isValidSudoku(board [][]byte) bool {
	var toCheck [][]byte
	for i := 0; i < 9; i++ {
		toCheck = append(toCheck, board[i])
	}
	for i := 0; i < 9; i++ {
		var columnToCheck []byte
		for j := 0; j < 9; j++ {
			columnToCheck = append(columnToCheck, board[j][i])
		}
		toCheck = append(toCheck, columnToCheck)
	}

	var upperLeft []byte
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			upperLeft = append(upperLeft, board[i][j])
		}
	}
	toCheck = append(toCheck, upperLeft)
	var upperMiddle []byte
	for i := 3; i < 6; i++ {
		for j := 0; j < 3; j++ {
			upperMiddle = append(upperMiddle, board[i][j])
		}
	}
	toCheck = append(toCheck, upperMiddle)
	var upperRight []byte
	for i := 6; i < 9; i++ {
		for j := 0; j < 3; j++ {
			upperRight = append(upperRight, board[i][j])
		}
	}
	toCheck = append(toCheck, upperRight)

	var middleLeft []byte
	for i := 0; i < 3; i++ {
		for j := 3; j < 6; j++ {
			middleLeft = append(middleLeft, board[i][j])
		}
	}
	toCheck = append(toCheck, middleLeft)
	var middle []byte
	for i := 3; i < 6; i++ {
		for j := 3; j < 6; j++ {
			middle = append(middle, board[i][j])
		}
	}
	toCheck = append(toCheck, middle)
	var middleRight []byte
	for i := 6; i < 9; i++ {
		for j := 3; j < 6; j++ {
			middleRight = append(middleRight, board[i][j])
		}
	}
	toCheck = append(toCheck, middleRight)

	var bottomLeft []byte
	for i := 0; i < 3; i++ {
		for j := 6; j < 9; j++ {
			bottomLeft = append(bottomLeft, board[i][j])
		}
	}
	toCheck = append(toCheck, bottomLeft)
	var bottomMiddle []byte
	for i := 3; i < 6; i++ {
		for j := 6; j < 9; j++ {
			bottomMiddle = append(bottomMiddle, board[i][j])
		}
	}
	toCheck = append(toCheck, bottomMiddle)
	var bottomRight []byte
	for i := 6; i < 9; i++ {
		for j := 6; j < 9; j++ {
			bottomRight = append(bottomRight, board[i][j])
		}
	}
	toCheck = append(toCheck, bottomRight)

	for _, dots := range toCheck {
		if !check(dots) {
			return false
		}
	}
	return true
}

func check(board []byte) bool {
	stats := map[int]bool{}
	for i := 0; i < len(board); i++ {
		num := board[i]
		if rune(num) == '.' {
			continue
		}
		atoi, err := strconv.Atoi(string(num))
		if err != nil {
			panic(err)
		}
		if atoi > 9 || atoi < 1 {
			return false
		}
		if stats[atoi] {
			return false
		}
		stats[atoi] = true
	}
	return true
}

func convertToNumber(sym byte) (int, error) {
	if rune(sym) == '.' {
		return 0, fmt.Errorf("invalid symbol")
	} else {
		atoi, err := strconv.Atoi(string(sym))
		if err != nil {
			return 0, err
		}
		return atoi, nil
	}
}
