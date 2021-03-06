package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const STEP = false // display steps or not
// states
const MOVE int = 0
const PICKCAN int = 10 // reward for picking up a can
const NOCAN int = -1   // punish for picking up when no can in cell
const WALL int = -5    // punish for running into a wall
const NOMOVE int = 0

const WIDTH = 10            // width of the board
const LENGTH = 10           // length of the board
const SIZE = WIDTH * LENGTH // number of cells on a board
const CANS int = 10         // number of soda cans
const CR float64 = 1.0      // crossover rate
const MUTATION int = 5      // mutation rate multiplied by 10000

const POPULATION int = 200  // number of population in a generation
const GENERATIONS int = 500 // number of generations
const ACTIONS int = 200     // number of actions per cleaning session
const SITUATIONS int = 243  // max number of possible SITUATIONS
const SESSIONS = 100        // Number of cleaning sessions a robby is evaluated on

const (
	MoveNorth  = iota // move north
	MoveSouth  = iota // move south
	MoveEast   = iota // move east
	MoveWest   = iota // move west
	SatyPut    = iota // Stay put
	PickUpCan  = iota // Pick Up Can
	MoveRandom = iota // Move Random
)

var actions = [7]string{"move north", "move south", "move east", "move west", "pick up", "move random", "stay put"}

var CellInfo = [5]string{"wall north", "wall south", "wall east", "wall west", "can"}

type Board [][]int

type Gen struct {
	boardArray [][]int
}

// setup board for robby. place soda cans around the map
func (gen *Gen) SetupBoard() [][]int {
	gen.boardArray = make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		gen.boardArray[i] = make([]int, 5)
	}

	for i := 0; i < SIZE; i++ {
		// walls around a cell
		if i < WIDTH { // north site is wall
			gen.boardArray[i][0] = 1
		}
		if i >= SIZE-WIDTH { // south site is wall
			gen.boardArray[i][1] = 1
		}
		if (i+1)%WIDTH == 0 { // east site is wall
			gen.boardArray[i][2] = 1
		}
		if i%WIDTH == 0 { // west site is wall
			gen.boardArray[i][3] = 1
		}

	}
	gen.ResetBoard()
	return gen.boardArray
}

func (gen *Gen) ResetBoard() {
	for i := 0; i < SIZE; i++ {
		r := rand.Intn(10000)
		gen.boardArray[i][4] = 0 // default has no can
		if r <= 5000 {
			gen.boardArray[i][4] = 2 // lay down a can
		}
		// cans around a cell
		if i >= WIDTH && gen.boardArray[i-WIDTH][4] == 2 { // north site has can
			gen.boardArray[i][0] = 2
		}
		if i < SIZE-WIDTH && gen.boardArray[i+WIDTH][4] == 2 { // south site has can
			gen.boardArray[i][1] = 2
		}
		if i+1 < SIZE && gen.boardArray[i+1][4] == 2 { // east site has can
			gen.boardArray[i][2] = 2
		}
		if i-1 >= 0 && gen.boardArray[i-1][4] == 2 { // west site has can
			gen.boardArray[i][3] = 2
		}
	}
}

func GenerateS() []int {
	result := make([]int, SITUATIONS)
	for i := 0; i < SITUATIONS; i++ {
		result[i] = rand.Intn(7)
	}
	return result
}
func GenerateStrategies() [][]int {
	result := make([][]int, POPULATION)
	for i := 0; i < POPULATION; i++ {
		result[i] = GenerateS()
	}
	return result
}

// returns the action according to the given stratigy and board situation
// TODO: return errors if value is invalid
func (gen Gen) GenerateAction(s []int, CurrCell []int) int {
	var StratIndex int
	lenCurrCell := len(CurrCell)

	for i := 0; i < lenCurrCell; i++ {
		StratIndex += int(math.Pow(3, float64(lenCurrCell-i-1))) * CurrCell[i]
	}
	return s[StratIndex]
}

func (gen *Gen) ActionOutcome(pos int, act int) (NewPos int, Outcome int) {
	if act == 5 {
		act = rand.Intn(6)
	}
	if act == 0 { // move north
		if gen.boardArray[pos][act] != 1 {
			return pos - WIDTH, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 1 { // move south
		if gen.boardArray[pos][act] != 1 {
			return pos + WIDTH, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 2 { // move east
		if gen.boardArray[pos][act] != 1 {
			return pos + 1, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 3 { // move west
		if gen.boardArray[pos][act] != 1 {
			return pos - 1, MOVE
		} else {
			return pos, WALL
		}
	} else if act == 4 { // pick up can
		if gen.boardArray[pos][act] == 2 {
			gen.boardArray[pos][act] = 0
			return pos, PICKCAN
		} else {
			return pos, NOCAN
		}
	}
	return pos, NOMOVE
}

func (gen Gen) EvalStrategy(s []int) int {
	var NewPos int
	var result int // outcome of an action
	for i := 0; i < ACTIONS; i++ {
		oc := 0
		CurrentPos := NewPos
		action := gen.GenerateAction(s, gen.boardArray[CurrentPos])
		NewPos, oc = gen.ActionOutcome(CurrentPos, action)
		if STEP {
			fmt.Printf("current cell: %d\n", CurrentPos)
			fmt.Printf("cell state: %d\n", gen.boardArray[CurrentPos])
			fmt.Printf("action: %s\n", actions[action])
			fmt.Printf("New position: %d\n", NewPos)
			fmt.Printf("Outcome: %d\n\n", oc)
		}
		result += oc
	}
	return result
}

func CalFit(s []int) int {
	var gen Gen
	var SumOc int
	gen.SetupBoard()
	//var wg sync.WaitGroup
	//wg.Add(SESSIONS)
	for i := 0; i < SESSIONS; i++ {
		//go func(i int) {
		//	defer wg.Done()
		gen.ResetBoard()
		SumOc += gen.EvalStrategy(s)
		//}(i)
	}
	//wg.Wait()
	return SumOc / SESSIONS
}

func CalculateFits(p [][]int) []int {
	fit := make([]int, POPULATION)
	for i := 0; i < POPULATION; i++ {
		fit[i] = CalFit(p[i])
		fmt.Printf("individual: %d ", i)
		fmt.Printf("fitness: %d\n", fit[i])
	}
	return fit
}

func RankSelection(s []int) int {
	sorted := make([]int, POPULATION)
	copy(sorted, s)
	sort.Ints(sorted) // sort s in ascending order
	rank := make([]int, len(s))
	var sum int
	for i := 0; i < len(s); i++ {
		sum += i
		rank[i] = sum
	}
	var SelectedIndx int
	RandomPosition := rand.Intn(sum - len(s))
	for i := 0; i < len(s); i++ {
		if (rank[i] >= RandomPosition && i == 0) || rank[i] >= RandomPosition && RandomPosition > rank[i-1] {
			SelectedIndx = i
			break
		}
	}
	for i := 0; i < len(rank); i++ {
		if s[i] == sorted[SelectedIndx] {
			return i
		}
	}

	return -1
}

func PickParents(f []int) (int, int) {
	a := RankSelection(f)
	b := RankSelection(f)
	return a, b
}

func GenChild(p1 []int, p2 []int) ([]int, []int) {
	split := rand.Intn(SITUATIONS)
	child1 := make([]int, SITUATIONS)
	child2 := make([]int, SITUATIONS)
	copy(child1[:split], p1[:split])
	copy(child1[split:], p2[split:])
	copy(child2[:split], p2[:split])
	copy(child2[split:], p1[split:])
	// random mutation
	for i := 0; i < len(child1); i++ {
		if rand.Intn(10000) <= MUTATION {
			child1[i] = rand.Intn(6)
			child2[i] = rand.Intn(6)
		}
	}
	return child1, child2
}

func NewGen(OldGen [][]int) [][]int {
	NewGeneration := make([][]int, POPULATION)
	f := CalculateFits(OldGen)
	for i := 0; i < POPULATION; i += 2 {
		p1, p2 := PickParents(f)
		NewGeneration[i], NewGeneration[i+1] = GenChild(OldGen[p1], OldGen[p2])
	}
	return NewGeneration
}

func PrintBoard(board [][]int) {
	fmt.Printf("===============\n")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%d\n", board[i])
	}
}

func PrintPopulation(population [][]int) {
	for i := 0; i < len(population); i++ {
		fmt.Printf("Individual: %d\n", i)
		fmt.Printf("%d\n", population[i])
	}
}

func PrintFitness(f []int) {
	for i := 0; i < POPULATION; i++ {
		fmt.Printf("individual: %d\n", i)
		fmt.Printf("fitness: %d\n", f[i])
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	p := GenerateStrategies() // Initial population
	fmt.Printf("=======Population=%d=======\n", 1)
	//PrintPopulation(p)
	//s := CalculateFits(p)
	//a,b := PickParents2(s)
	//c, d := PickParents(s)
	//fmt.Printf("%d %d\n",a,b)
	//fmt.Printf("%d %d\n", c, d)

	for i := 0; i < 1000; i++ {
		fmt.Printf("=======Population=%d=======\n", i+2)
		newp := NewGen(p)
		p = newp
	}
	PrintPopulation(p)

	//s := GenerateS()
	//fmt.Printf("%d", CalFit(s))
	//var gen Gen
	//gen.SetupBoard()
	//gen.EvalStrategy(s)
	//fmt.Print(s)
	//fmt.Printf("scores: %d", gen.EvalStrategy(s))
}
