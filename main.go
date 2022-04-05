package main

import "fmt"

func main() {
	var board [9][9]int
	initBoard(&board)
	var gameInitialState []Pair
	gameInitialState = append(gameInitialState, Pair{0, 1})
	var values []int
	values = append(values, 9)
	setInitialState(&board, gameInitialState, values)
	printBoard(board)

	population := generatePopulation()
	sortPopulation(population)

	for i := 0; i < PopulationSize; i++ {
		fmt.Println(population[i].fitness)
	}
}
