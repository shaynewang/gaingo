package main

import (
	"fmt"
	"math/rand"
	"time"
)

const STEP = false
// states
const MOVE int = 0
const PICKCAN int = 10 // reward for picking up a can
const NOCAN int = -1 // punish for picking up when no can in cell
const WALL int = -5 // punish for running into a wall
const NOMOVE int = 0

const WIDTH = 10 // width of the board
const LENGTH = 10 // length of the board
const SIZE = WIDTH * LENGTH
const CANS int = 10 // number of soda cans
const POPULATION int = 200 // number of population in a generation
const CR float64 = 1.0       // crossover rate
const MUTATION float64 = 0.5 // mutation rate
const GENERATIONS int = 500 // number of generations
const ACTIONS int = 200    // number of actions per cleaning session
const SITUATIONS int = 243 // max number of possible SITUATIONS
const SESSIONS = 100

const (
	MoveNorth  = iota // move north
	MoveSouth  = iota // move south
	MoveEast   = iota // move east
	MoveWest   = iota // move west
	SatyPut    = iota // Stay put
	PickUpCan  = iota // Pick Up Can
	MoveRandom = iota // Move Random
)
var actions = [7]string {"move north","move south","move east","move west","pick up","move random","stay put"}

var CellInfo = [5]string {"wall north","wall south","wall east","wall west","can"}

type Board [][]int

type Gen struct {
	boardArray [][] int
}

// setup board for robby. place soda cans around the map
func (gen *Gen) SetupBoard(NumberOfCans int) [][] int {
	gen.boardArray = make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		gen.boardArray[i] = make([]int, 5)
		if i < WIDTH { // north side has wall
			gen.boardArray[i][0] = 1
		}
		if i >= SIZE - WIDTH { // south side has wall
			gen.boardArray[i][1] = 1
		}
		if (i+1)%WIDTH == 0{ // east side has wall
			gen.boardArray[i][2] = 1
		}
		if i%WIDTH == 0{ // west side has wall
			gen.boardArray[i][3] = 1
		}
		rand.Seed(time.Now().UnixNano())	
		r := rand.Intn(100000)
		gen.boardArray[i][4] = 0 // default has no can
		if r <= 50000{
			gen.boardArray[i][4] = 1 // default has no can
		}
	}
	return gen.boardArray
}

func GenerateS()[SITUATIONS] int{
	var result [SITUATIONS] int
	for i := 0; i < SITUATIONS; i++ {
		rand.Seed(time.Now().UnixNano())
		result[i] = rand.Intn(7)
	}
	return result
}
func GenerateStrategies() [POPULATION][SITUATIONS]int {
	var result [POPULATION][SITUATIONS]int
	for i := 0; i < POPULATION; i++ {
		result[i] = GenerateS()
	}
	return result
}

func (gen Gen) ActionOutcome(pos int, act int) (NewPos int , Outcome int) {
	if act == 5 {
		rand.Seed(time.Now().UnixNano())
		act = rand.Intn(6)
	}
	if act == 0 { // move north
		if gen.boardArray[pos][act] == 0 {
			return pos - WIDTH, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 1 { // move south
		if gen.boardArray[pos][act] == 0 {
			return pos + WIDTH, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 2 { // move east
		if gen.boardArray[pos][act] == 0 {
			return pos + 1, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 3 { // move west
		if gen.boardArray[pos][act] == 0 {
			return pos - 1, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 4 { // pick up can
		if gen.boardArray[pos][act] == 1 {
			gen.boardArray[pos][act] = 0
			return pos, PICKCAN
		} else {
			return pos, NOCAN
		}
	}
	return pos, NOMOVE
}

func (gen Gen)EvalStrategy(s [SITUATIONS]int) int{
	var NewPos int
	result := 0 // outcome of an action
	for i:=0; i< len(s); i++{
		oc := 0
		CurrentPos := NewPos
		NewPos, oc = gen.ActionOutcome(CurrentPos, s[i])
		if STEP{
			fmt.Printf("current cell: %d\n",CurrentPos)
			fmt.Printf("cell state: %d\n", gen.boardArray[CurrentPos])
			fmt.Printf("action: %s\n",actions[s[i]])
			fmt.Printf("New position: %d\n", NewPos)
			fmt.Printf("Outcome: %d\n\n",oc)
		}
		result += oc
	}
	return result
}

func CalFit(s [SITUATIONS] int) int{
	var SumOc int
	var gen Gen
	for i:= 0; i < SESSIONS; i++{
		gen.SetupBoard(50)
		SumOc += gen.EvalStrategy(s)
	}
	return SumOc/SITUATIONS
}

func CalculateFits(p [POPULATION][SITUATIONS]int) [] int{
	fit := make([]int, POPULATION)

	for i:=0; i < POPULATION; i++{
		fit[i] = CalFit(p[i])
		fmt.Printf("individual: %d ", i)
		fmt.Printf("fitness: %d\n", fit[i])
	}
	return fit
}

func main() {
	//var c_x []int
	//var c_y []int
	//Setup_board(SIZE, SIZE, CANS, c_x, c_y)
	//s := GenerateS()
	//for i := 0; i < POPULATION; i++ {
	//	fmt.Printf("Individual %d:\n",i)
	//	for j := 0; j < SITUATIONS; j++ {		
	//		fmt.Printf("%d", s[i][j])
	//	}
	//	fmt.Printf("\n")
	//}
	//	for i:=0; i < SIZE; i++{
	//	fmt.Printf("\nCell %d:\n", i+1)
	//	fmt.Printf("%d", gen.boardArray[i])
	//	fmt.Print("  ")
	//}

	p := GenerateStrategies() // Initial population
	CalculateFits(p)

}
