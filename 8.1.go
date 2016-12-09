package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"regexp"
	"strconv"
)

// const WIDTH = 7
const WIDTH = 50
const HEIGHT = 6
// const HEIGHT = 3
type Row [WIDTH]bool
type Grid [HEIGHT]Row

func printGrid (grid Grid) {
	for i := 0; i < HEIGHT; i++ {
		fmt.Println("")
		for j := 0; j < WIDTH; j++ {
			if grid[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print("-")
			}
		}
	}
	fmt.Println("")
}

func main() {

	if file, err := os.Open("input.8.txt"); err == nil {
		scanner := bufio.NewScanner(file)
		defer file.Close()

		grid := Grid{}

		for scanner.Scan() {
			instruction := scanner.Text()
			fmt.Println(instruction)

			if strings.Contains(instruction, "rect") {
				re := regexp.MustCompile("(\\d+)x(\\d+)")
				vars := re.FindStringSubmatch(scanner.Text())
				width, _ := strconv.Atoi(vars[1])
				height, _ := strconv.Atoi(vars[2])

				for i := 0; i < height; i++ {
					for j := 0; j < width; j++ {
						grid[i][j] = true
					}
				}
				printGrid(grid)
			} else if strings.Contains(instruction, "rotate row") {
				re := regexp.MustCompile("(\\d+)\\sby\\s(\\d+)")
				vars := re.FindStringSubmatch(scanner.Text())
				row, _ := strconv.Atoi(vars[1])
				offset, _ := strconv.Atoi(vars[2])
				tmp := make([]bool, len(grid[row]))
				copy(tmp, grid[row][:])

				for i := 0; i < WIDTH; i++ {
					trueOffset := (i + offset) % WIDTH
					grid[row][trueOffset] = tmp[i]
				}
					printGrid(grid)
			} else if strings.Contains(instruction, "rotate column") {
				re := regexp.MustCompile("(\\d+)\\sby\\s(\\d+)")
				vars := re.FindStringSubmatch(scanner.Text())
				col, _ := strconv.Atoi(vars[1])
				offset, _ := strconv.Atoi(vars[2])
				tmp := make([]bool, HEIGHT)

				for i := 0; i < HEIGHT; i++ {
						tmp[i] = grid[i][col]
				}

				for i := 0; i < HEIGHT; i++ {
					trueOffset := (i + offset) % HEIGHT
					grid[trueOffset][col] = tmp[i]
				}
					printGrid(grid)
			} else {

			}

		}
		count := 0
		for _, row := range grid {
			for _, cell := range row {
				if cell {
					count++;
				}
			}
		}
		fmt.Println(count)
	}
}
