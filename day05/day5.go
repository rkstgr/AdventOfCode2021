package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Line struct {
	start, end Point
}

func (line Line) isAxial() bool {
	if line.start.x == line.end.x || line.start.y == line.end.y {
		return true
	} else {
		return false
	}
}

func (line Line) Distance() int {
	xDistance := int(math.Abs(float64(line.end.x - line.start.x)))
	yDistance := int(math.Abs(float64(line.end.y - line.start.y)))

	if xDistance > yDistance {
		return xDistance
	} else {
		return yDistance
	}
}

func (line Line) Interpolate() []Point {
	xDir, yDir := 1, 1
	if line.end.x < line.start.x {
		xDir = -1
	} else if line.end.x == line.start.x {
		xDir = 0
	}

	if line.end.y < line.start.y {
		yDir = -1
	} else if line.end.y == line.start.y {
		yDir = 0
	}
	distance := line.Distance()
	points := make([]Point, distance+1)

	var x, y int

	for i := 0; i <= distance; i++ {
		x = line.start.x + i*xDir
		y = line.start.y + i*yDir
		point := Point{x, y}
		points[i] = point
	}
	return points
}

func parseLineInput(lineInput string) Line {
	inputs := strings.Split(lineInput, " -> ")
	startInputs := strings.Split(inputs[0], ",")
	endInputs := strings.Split(inputs[1], ",")

	startX, _ := strconv.Atoi(startInputs[0])
	startY, _ := strconv.Atoi(startInputs[1])
	endX, _ := strconv.Atoi(endInputs[0])
	endY, _ := strconv.Atoi(endInputs[1])

	start := Point{
		x: startX,
		y: startY,
	}

	end := Point{
		x: endX,
		y: endY,
	}

	return Line{
		start,
		end,
	}
}

func main() {
	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(bs), "\n")

	sol1 := part1(input)
	sol2 := part2(input)

	end := time.Since(start)

	fmt.Println("Part 1: ", sol1)
	fmt.Println("Part 2: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(input []string) int {
	var ground [1000][1000]int

	for _, lineInput := range input {
		line := parseLineInput(lineInput)
		if line.isAxial() {
			points := line.Interpolate()

			for _, point := range points {

				ground[point.x][point.y] += 1
			}
		}
	}

	var s int
	for _, row := range ground {
		for _, v := range row {
			if v >= 2 {
				s++
			}
		}
	}
	return s
}

func part2(input []string) int {
	var ground [1000][1000]int

	for _, lineInput := range input {
		line := parseLineInput(lineInput)
		points := line.Interpolate()

		for _, point := range points {

			ground[point.x][point.y] += 1
		}

	}

	var s int
	for _, row := range ground {
		for _, v := range row {
			if v >= 2 {
				s++
			}
		}
	}
	return s
}
