package main

import (
	"math/rand"
	"sync"
	"time"
)

func geneticConcurrency(board [9][9]int, gameInitialState []Pair, values []int) {
	rand.Seed(time.Now().UnixNano())
	var initialState InitialState

	setInitialState(&board, gameInitialState, values, &initialState)
	geneticAlgorithm(&initialState)
}

func main() {
	var board [9][9]int

	var gameInitialState []Pair
	gameInitialState = append(gameInitialState,
		Pair{0, 1}, Pair{0, 7})
	//Pair{1, 1}, Pair{1, 7},
	//Pair{2, 0}, Pair{2, 8},
	//Pair{3, 2}, Pair{3, 3}, Pair{3, 5}, Pair{3, 6},
	//Pair{4, 3}, Pair{4, 4}, Pair{4, 6}, Pair{4, 7},
	//Pair{5, 3}, Pair{5, 4}, Pair{5, 6}, Pair{5, 7},
	//Pair{6, 4},
	//Pair{7, 0}, Pair{7, 1}, Pair{7, 7}, Pair{7, 8},
	//Pair{8, 1}, Pair{8, 3}, Pair{8, 5}, Pair{8, 7})
	var values []int
	// values = append(values, 1, 6, 2, 3, 9, 1, 2, 8, 6, 4, 6, 9, 2, 3, 4, 3, 7, 1, 8, 7, 3, 5, 8, 8, 2, 1, 9)
	values = append(values, 1, 6)
	var threads sync.WaitGroup
	threads.Add(5)
	for i := 0; i < 5; i++ {
		go geneticConcurrency(board, gameInitialState, values)
	}
	threads.Wait()
}
