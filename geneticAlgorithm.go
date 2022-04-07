package main

import (
	"fmt"
	"math/rand"
)

var PopulationSize = 7000
var Outliers = 0.1
var CrossOvers = 0.6
var Mutants = 0.4

func randBetween(a int, b int) int {
	span := (b - a) + 1
	return a + (rand.Int() % span)
}

func fitnessFunction(board *[9][9]int) int {
	mistakes := 0
	_, mistakes = checkBoard(board)
	return mistakes
}

type Individual struct {
	gene    [9][9]int
	fitness int
}

func createIndividual() Individual {
	newIndividual := Individual{fitness: 0}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			status, value := isFromInitialState(i, j)
			if !status {
				newIndividual.gene[i][j] = randBetween(1, 9)
			} else {
				newIndividual.gene[i][j] = value
			}
		}
	}
	return newIndividual
}

func sortPopulation(individuals []Individual) []Individual {
	for i := 1; i < PopulationSize; i++ {
		key := individuals[i]
		j := i - 1
		for j >= 0 && key.fitness < individuals[j].fitness {
			individuals[j+1] = individuals[j]
			j--
		}
		individuals[j+1] = key
	}
	return individuals
}

func generatePopulation() []Individual {
	var population []Individual
	for i := 0; i < PopulationSize; i++ {
		individual := createIndividual()
		individual.fitness = fitnessFunction(&individual.gene)
		population = append(population, individual)
	}
	return population
}

func mutate(individual Individual) Individual {
	mutant := Individual{}
	mutant.gene = individual.gene

	_, mistakes := checkBoard(&individual.gene)

	if mistakes < 5 {
		for i := 0; i < randBetween(0, 20); {
			randomX := randBetween(0, 8)
			randomY := randBetween(0, 8)

			status, _ := isFromInitialState(randomX, randomY)
			if !status {
				i++
				mutant.gene[randomX][randomY] = randBetween(1, 9)
			}
		}
		mutant.fitness = fitnessFunction(&mutant.gene)
		return mutant
	}

	process := randBetween(0, 3)
	if process == 0 {
		// mutate a random row
		randomRow := randBetween(0, 8)
		for i := 0; i < 9; i++ {
			status, _ := isFromInitialState(randomRow, i)
			if !status {
				mutant.gene[randomRow][i] = randBetween(1, 9)
			}
		}
	} else if process == 1 {
		// mutate a random column
		randomCol := randBetween(0, 8)
		for i := 0; i < 9; i++ {
			status, _ := isFromInitialState(i, randomCol)
			if !status {
				mutant.gene[i][randomCol] = randBetween(1, 9)
			}
		}
	} else {
		// mutate a random grid
		randomGridX := randBetween(0, 2)
		randomGridY := randBetween(0, 2)

		for i := randomGridX * 3; i < randomGridX*3+3; i++ {
			for j := randomGridY * 3; j < randomGridY*3+3; j++ {
				status, _ := isFromInitialState(i, j)
				if !status {
					mutant.gene[i][j] = randBetween(1, 9)
				}
			}
		}
	}
	mutant.fitness = fitnessFunction(&mutant.gene)
	return mutant
}

func naturalSelection(individuals []Individual) []Individual {
	var buffer []Individual
	outliersRange := int(float64(PopulationSize) * Outliers)
	crossOverRange := int(float64(PopulationSize) * CrossOvers)

	for i := 0; i < outliersRange; i++ {
		buffer = append(buffer, individuals[i])
	}
	for i := outliersRange; i < crossOverRange+outliersRange; i++ {
		parent1 := individuals[randBetween(0, PopulationSize-1)]
		parent2 := individuals[randBetween(0, PopulationSize-1)]

		var offspring1 Individual

		barrier := randBetween(1, 9)
		for i := 0; i < barrier; i++ {
			offspring1.gene[i] = parent2.gene[i]
		}
		for i := barrier; i < 9; i++ {
			offspring1.gene[i] = parent1.gene[i]
		}

		offspring1.fitness = fitnessFunction(&offspring1.gene)

		buffer = append(buffer, offspring1)
	}
	for i := outliersRange + crossOverRange; i < PopulationSize; i++ {
		buffer = append(buffer, mutate(individuals[randBetween(0, PopulationSize-1)]))
	}
	return buffer
}

func geneticAlgorithm() {
	population := generatePopulation()
	for i := 0; i < 100000; i++ {
		population = sortPopulation(population)
		fmt.Println(population[0].fitness, " ", i)
		// printBoard(population[0].gene)
		if population[0].fitness == 0 {
			printBoard(population[0].gene)
			break
		}
		population = naturalSelection(population)
	}
	printBoard(population[0].gene)
}
