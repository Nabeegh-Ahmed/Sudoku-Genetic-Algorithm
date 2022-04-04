package main

func fitnessFunction(board *[9][9]int) int {
	mistakes := 0
	_, mistakes = checkBoard(board)
	return mistakes
}
