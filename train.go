package main

import (
	"fmt"
	"math/rand"
	"time"
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

func main() {
	var c_x []int
	var c_y []int
	Setup_board(10, 15, 5, c_x, c_y)
}
