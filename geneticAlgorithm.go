package main

import (
	"math/rand"
)

var PopulationSize = 10

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
				newIndividual.gene[i][j] = randBetween(0, 9)
			} else {
				newIndividual.gene[i][j] = value
			}
		}
	}
	return newIndividual
}

func sortPopulation(individuals []Individual) {
	for i := 1; i < PopulationSize; i++ {
		key := individuals[i]
		j := i - 1
		for j >= 0 && key.fitness < individuals[j].fitness {
			individuals[j+1] = individuals[j]
			j--
		}
		individuals[j+1] = key
	}
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
