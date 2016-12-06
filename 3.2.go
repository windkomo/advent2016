package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Shape struct {
	length1 int
	length2 int
	length3 int
}

func (shape Shape) isTriangle() bool {
	return shape.length1+shape.length2 > shape.length3 && shape.length1+shape.length3 > shape.length2 && shape.length2+shape.length3 > shape.length1
}

func main() {

	if file, err := os.Open("input.3.txt"); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0
		re := regexp.MustCompile("[0-9]+")
		shape1 := Shape{0, 0, 0}
		shape2 := Shape{0, 0, 0}
		shape3 := Shape{0, 0, 0}
		line := 1

		for scanner.Scan() {
			lengths := re.FindAllString(scanner.Text(), -1)
			switch line % 3 {
			case 1:
				shape1.length1, _ = strconv.Atoi(lengths[0])
				shape2.length1, _ = strconv.Atoi(lengths[1])
				shape3.length1, _ = strconv.Atoi(lengths[2])
			case 2:
				shape1.length2, _ = strconv.Atoi(lengths[0])
				shape2.length2, _ = strconv.Atoi(lengths[1])
				shape3.length2, _ = strconv.Atoi(lengths[2])
			case 0:
				shape1.length3, _ = strconv.Atoi(lengths[0])
				shape2.length3, _ = strconv.Atoi(lengths[1])
				shape3.length3, _ = strconv.Atoi(lengths[2])

				if shape1.isTriangle() {
					count++
				}
				if shape2.isTriangle() {
					count++
				}
				if shape3.isTriangle() {
					count++
				}
			}
			line++
		}

		fmt.Println(count)
	}
}
