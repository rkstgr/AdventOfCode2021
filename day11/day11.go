package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type Board struct {
	energy  [10][10]int
	flashed [10][10]bool
	flashes int
}

func parseInputs(input []string) Board {
	board := Board{}
	for i, row := range input {
		for j, z := range row {
			board.energy[i][j], _ = strconv.Atoi(fmt.Sprintf("%c", z))
		}
	}
	return board
}

func (board *Board) Increase() {
	for i := range board.energy {
		for j := range board.energy[i] {
			board.energy[i][j]++
		}
	}
}

func (board *Board) IncreaseNeighbours(i, j int) {
	offsets := [3]int{-1, 0, 1}
	for _, xOffset := range offsets {
		for _, yOffset := range offsets {
			x := i + xOffset
			y := j + yOffset
			if x != i || y != j { //don't increase center point
				if x >= 0 && x < 10 { // check boundaries
					if y >= 0 && y < 10 {
						board.energy[x][y]++
					}
				}
			}
		}
	}
}

func (board *Board) Flash() {
	for i := range board.energy {
		for j := range board.energy[i] {
			// if energy > 9 & it has not flashed already
			if board.energy[i][j] > 9 && !board.flashed[i][j] {
				// i j flashes
				board.flashed[i][j] = true
				board.flashes++
				// increase adjacent
				board.IncreaseNeighbours(i, j)
				board.Flash()
			}
		}
	}
}

func (board *Board) Reset() {
	for i := range board.energy {
		for j := range board.energy[i] {
			// if energy > 9 & it has not flashed already
			if board.energy[i][j] > 9 {
				if !board.flashed[i][j] {
					panic("cannot reset, there is a point which high energy which has not flashed yes")
				}
				board.energy[i][j] = 0
				board.flashed[i][j] = false
			}
		}
	}
}

func stepBoard(board Board) Board {
	board.Increase()
	board.Flash()
	board.Reset()
	return board
}

func (board Board) String() string {
	output := ""
	for i, row := range board.energy {
		output += fmt.Sprintf("%v", row)
		output += " | "
		output += fmt.Sprintf("%v\n", board.flashed[i])
	}
	return output
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
	board := parseInputs(input)
	for i := 0; i < 100; i++ {
		board = stepBoard(board)
	}
	return board.flashes
}

func part2(input []string) int {
	board := parseInputs(input)

	flashedDelta := 0
	flashes := 0
	i :=0
	for flashedDelta != 100 {
		board = stepBoard(board)
		flashedDelta = board.flashes - flashes
		flashes = board.flashes
		i++
	}

	return i
}
