package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Board struct {
	numbers [5][5]int
	marked  [5][5]bool
}

var multSpaces, _ = regexp.Compile("\\s+")

func parseBingoNumbersInput(input string) []int {
	numbersStr := strings.Split(input, ",")
	var numbers []int
	for _, nStr := range numbersStr {
		n, _ := strconv.Atoi(nStr)
		numbers = append(numbers, n)
	}
	return numbers
}

func parseBoards(input []string) []Board {
	var boards []Board
	for _, boardInput := range input {
		board := parseBoardInput(strings.Split(boardInput, "\n"))
		boards = append(boards, board)
	}
	return boards
}

func parseBoardInput(input []string) Board {
	var nums [5][5]int
	var n int
	for i, line := range input {
		line = strings.TrimSpace(line)
		numbers := multSpaces.Split(line, -1)
		for j, n_str := range numbers {
			n, _ = strconv.Atoi(n_str)
			nums[i][j] = n
		}
	}
	return Board{
		numbers: nums,
		marked:  [5][5]bool{},
	}
}

func (b Board) markedCol(j int) [5]bool {
	var col [5]bool
	for i := range b.marked {
		col[i] = b.marked[i][j]
	}
	return col
}

func All(ba []bool) bool {
	for _, v := range ba {
		if !v {
			return false
		}
	}
	return true
}

func hasBoardWon(board Board) bool {
	// check all rows
	for _, row := range board.marked {
		if All(row[:]) {
			return true
		}
	}
	// check all columns
	for colId := range board.marked[0] {
		col := board.markedCol(colId)
		if All(col[:]) {
			return true
		}
	}
	return false
}

func markBoards(boards []Board, num int) []Board {
	for i, board := range boards {
		newBoard := markBoard(board, num)
		boards[i] = newBoard
	}
	return boards
}

func markBoard(board Board, num int) Board {
	// iterate over numbers
	for i, row := range board.numbers {
		for j, n := range row {
			if n == num {
				board.marked[i][j] = true
			}
		}
	}
	return board
}

func boardSum(board Board) int {
	x := 0
	for i, row := range board.marked {
		for j, mark := range row {
			if !mark {
				x += board.numbers[i][j]
			}
		}
	}
	return x
}

func main() {
	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic("could not read")
	}

	input := strings.Split(string(bs), "\n\n")
	bingoNumbers := parseBingoNumbersInput(input[0])
	boards := parseBoards(input[1:])

	sol1 := part1(bingoNumbers, boards)
	sol2 := part2(bingoNumbers, boards)

	end := time.Since(start)

	fmt.Println("Part 1: ", sol1)
	fmt.Println("Part 2: ", sol2)
	fmt.Println("Time: ", end)

}

func part1(bingoNumbers []int, boards []Board) int {
	for _, bingoNum := range bingoNumbers {
		boards = markBoards(boards, bingoNum)
		for _, board := range boards {
			if hasBoardWon(board) {
				return boardSum(board) * bingoNum
			}
		}
	}
	return -1
}

func part2(bingoNumbers []int, boards []Board) int {
	var boardsFinished []bool
	for range boards {
		boardsFinished = append(boardsFinished, false)
	}
	for _, bingoNum := range bingoNumbers {
		boards = markBoards(boards, bingoNum)
		for bId, board := range boards {
			if hasBoardWon(board) {
				boardsFinished[bId] = true
			}
			if All(boardsFinished) {
				return boardSum(board) * bingoNum
			}
		}
	}
	return -1
}
