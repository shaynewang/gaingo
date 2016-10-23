package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE int = 10
const CANS int = 10
const POPULATION int = 200
const CR float64 = 1.0       // crossover rate
const MUTATION float64 = 0.5 // mutation rate
const GENERATIONS int = 500
const ACTIONS int = 200    // number of actions per cleaning session
const SITUATIONS int = 243 // max number of possible situations
const (
	MoveNorth  = iota // move north
	MoveSouth  = iota // move south
	MoveEast   = iota // move east
	MoveWest   = iota // move west
	SatyPut    = iota // Stay put
	PickUpCan  = iota // Pick Up Can
	MoveRandom = iota // Move Random
)

// setup board for robby. place pop cans around the map
func Setup_board(x, y, num_cans int, cans_x []int, cans_y []int) int {
	for i := 0; i < num_cans; i++ {
		rand.Seed(time.Now().UnixNano())
		cans_x = append(cans_x, rand.Intn(x))
		cans_y = append(cans_y, rand.Intn(y))
		fmt.Printf("x = %d, y = %d\n", cans_x[i], cans_y[i])
	}
	return 1
}

func Generate_Strategies(n int) [POPULATION][SITUATIONS]int {
	var result [POPULATION][SITUATIONS]int
	for i := 0; i < POPULATION; i++ {
		for j := 0; j < SITUATIONS; j++ {
			rand.Seed(time.Now().UnixNano())
			result[i][j] = rand.Intn(7)
		}
	}
	return result
}

func main() {
	//var c_x []int
	//var c_y []int
	//Setup_board(SIZE, SIZE, CANS, c_x, c_y)
	s := Generate_Strategies(200)
	for i := 0; i < POPULATION; i++ {
		for j := 0; j < SITUATIONS; j++ {
			fmt.Printf("%d", s[i][j])
		}
		fmt.Printf("\n")
	}
}
