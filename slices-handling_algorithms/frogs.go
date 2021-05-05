package main

import (
	"fmt"
	"math"
)

func main() {
	blocks := []int{1, 2, 5, 5, 6, 3}
	fmt.Println(Solution(blocks))
}

func Solution(blocks []int) int {
	var currentHeight int = blocks[0]
	var max float64
	var i int
	goingUp := false

	if len(blocks) == 2 {
		return len(blocks)
	}

	for i < len(blocks)-1 {
		if (!goingUp && currentHeight >= blocks[i+1]) || (goingUp && currentHeight <= blocks[i+1]) {
			currentHeight = blocks[i+1]
			i++
		} else if !goingUp {
			goingUp = true
		} else if goingUp {
			var distance int = i
			max = math.Max(max, float64(distance))

			goingUp = false
		}
	}

	var d = i
	max = math.Max(max, float64(d))

	return int(max)
}
