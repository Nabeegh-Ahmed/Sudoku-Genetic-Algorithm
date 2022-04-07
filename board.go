package main

import "fmt"

type Pair struct {
	first  int
	second int
}

type InitialState struct {
	coords []Pair
	values []int
}

func setInitialState(board *[9][9]int, locs []Pair, values []int, initialState *InitialState) {
	initialState.coords = locs
	initialState.values = values
	for i := 0; i < len(locs); i++ {
		board[locs[i].first][locs[i].second] = values[i]
	}
}

func isFromInitialState(row int, col int, initialState *InitialState) (bool, int) {
	for i := range initialState.coords {
		if row == initialState.coords[i].first && col == initialState.coords[i].second {
			return true, initialState.values[i]
		}
	}
	return false, -1
}

func placeMove(board *[9][9]int, row int, col int, move int, initialState *InitialState) {
	status, _ := isFromInitialState(row, col, initialState)
	if !status {
		board[row][col] = move
	}
}

func checkGrid(board *[9][9]int, gridX int, gridY int) (bool, int) {
	mistakes := 0
	var valuesMap [10]int
	for i := gridX; i < gridX+3; i++ {
		for j := gridY; j < gridY+3; j++ {
			if board[i][j] == -1 {
				mistakes++
			} else {
				valuesMap[board[i][j]]++
				if valuesMap[board[i][j]] > 1 {
					mistakes++
				}
			}
		}
	}
	return mistakes == 0, mistakes
}

func checkRow(board *[9][9]int, row int) (bool, int) {
	mistakes := 0
	var valuesMap [10]int
	for i := 0; i < 9; i++ {
		if board[row][i] == -1 {
			mistakes++
		} else {
			valuesMap[board[row][i]]++
			if valuesMap[i] > 1 {
				mistakes++
			}
		}
	}
	return mistakes == 0, mistakes
}

func checkCol(board *[9][9]int, col int) (bool, int) {
	mistakes := 0
	var valuesMap [10]int
	for i := 0; i < 9; i++ {
		if board[i][col] == -1 {
			mistakes++
		} else {
			valuesMap[board[i][col]]++
			if valuesMap[i] > 1 {
				mistakes++
			}
		}
	}
	return mistakes == 0, mistakes
}

func initBoard(board *[9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			board[i][j] = -1
		}
	}
}

func checkBoard(board *[9][9]int) (bool, int) {
	mistakes := 0
	for i := 0; i < 9; i++ {
		rowMistakes := 0
		_, rowMistakes = checkRow(board, i)
		mistakes += rowMistakes

		colMistakes := 0
		_, colMistakes = checkCol(board, i)
		mistakes += colMistakes
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			gridMistakes := 0
			_, gridMistakes = checkGrid(board, i*3, j*3)
			mistakes += gridMistakes
		}
	}
	return mistakes == 0, mistakes
}

func printBoard(board [9][9]int) {
	for i := 0; i < 12; i++ {
		fmt.Printf("+--")
	}
	for i := 0; i < 9; i++ {
		fmt.Println("+")
		for j := 0; j < 9; j++ {
			if board[i][j] == -1 {
				fmt.Printf("|   ")
			} else {
				fmt.Printf("| %d ", board[i][j])
			}
		}
		fmt.Println("|")
		for i := 0; i < 12; i++ {
			fmt.Printf("+--")
		}
	}
	fmt.Println("+")
}
