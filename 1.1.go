package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

type Walker struct {
	direction string
	position  Position
}

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func (walker *Walker) turn(direction string) {
	directions := []string{"N", "E", "S", "W"}
	directionIndex := SliceIndex(len(directions), func(i int) bool { return directions[i] == walker.direction })

	switch direction {
	case "R":
		if directionIndex+1 >= len(directions) {
			directionIndex = 0
		} else {
			directionIndex++
		}
	case "L":
		if directionIndex == 0 {
			directionIndex = len(directions) - 1
		} else {
			directionIndex--
		}
	}

	walker.direction = directions[directionIndex]
}

func (walker *Walker) walk(steps int) {
	switch walker.direction {
	case "N":
		walker.position.y += steps
	case "E":
		walker.position.x += steps
	case "S":
		walker.position.y -= steps
	case "W":
		walker.position.x -= steps
	}
}

func main() {
	walker := Walker{
		position:  Position{0, 0},
		direction: "N",
	}
	instructions := strings.Split("L5, R1, L5, L1, R5, R1, R1, L4, L1, L3, R2, R4, L4, L1, L1, R2, R4, R3, L1, R4, L4, L5, L4, R4, L5, R1, R5, L2, R1, R3, L2, L4, L4, R1, L192, R5, R1, R4, L5, L4, R5, L1, L1, R48, R5, R5, L2, R4, R4, R1, R3, L1, L4, L5, R1, L4, L2, L5, R5, L2, R74, R4, L1, R188, R5, L4, L2, R5, R2, L4, R4, R3, R3, R2, R1, L3, L2, L5, L5, L2, L1, R1, R5, R4, L3, R5, L1, L3, R4, L1, L3, L2, R1, R3, R2, R5, L3, L1, L1, R5, L4, L5, R5, R2, L5, R2, L1, L5, L3, L5, L5, L1, R1, L4, L3, L1, R2, R5, L1, L3, R4, R5, L4, L1, R5, L1, R5, R5, R5, R2, R1, R2, L5, L5, L5, R4, L5, L4, L4, R5, L2, R1, R5, L1, L5, R4, L3, R4, L2, R3, R3, R3, L2, L2, L2, L1, L4, R3, L4, L2, R2, R5, L1, R2", ", ")

	for _, instruction := range instructions {
		direction := instruction[0:1]
		steps, _ := strconv.Atoi(instruction[1:])

		walker.turn(direction)
		walker.walk(steps)
	}

	fmt.Printf("Distance: %.0f blocks", math.Abs(float64(walker.position.x))+math.Abs(float64(walker.position.y)))
}
