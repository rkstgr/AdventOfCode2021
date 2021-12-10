package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

/*
Opening,Closing
123,125:  {,}
 91, 93:  [,]
 60, 62:  <,>
 40, 41:  (,)
*/

func isPair(i, j int32) bool {
	if i == 40 && j == 41 {
		return true
	} else if j == i+2 {
		return true
	}
	return false
}

func isOpening(i int32) bool {
	switch i {
	case 123:
		return true
	case 91:
		return true
	case 60:
		return true
	case 40:
		return true
	default:
		return false
	}
}

func isClosing(i int32) bool {
	switch i {
	case 125:
		return true
	case 93:
		return true
	case 62:
		return true
	case 41:
		return true
	default:
		return false
	}
}

func parseLine(lineInput string) (string, []string) {
	//return the last valid
	currentPos := -1
	var stack [120]int32

	closingBracket := map[string]string{"{": "}", "[": "]", "<": ">", "(": ")"}

	for _, i := range lineInput {
		if isOpening(i) {
			currentPos++
			stack[currentPos] = i
		} else if isClosing(i) {
			if currentPos < 0 {
				panic("found closing bracket without opening bracket")
			} else if isPair(stack[currentPos], i) {
				currentPos--
			} else {
				// we have unresolved opening brackets, but this closing symbol does not fit
				// -> corrupt line
				return fmt.Sprintf("%c", i), []string{}
			}
		}
	}
	missingBrackets := make([]string, currentPos+1)
	for j := currentPos; j >= 0; j-- {
		char := fmt.Sprintf("%c", stack[j])
		closingSymb := closingBracket[char]
		missingBrackets[currentPos-j] = closingSymb
	}
	return "", missingBrackets

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
	totalScore := 0
	scoreMap := map[string]int{"": 0, ")": 3, "]": 57, "}": 1197, ">": 25137}
	lastChar := ""
	score := 0
	for _, line := range input {
		lastChar, _ = parseLine(line)
		score = scoreMap[lastChar]
		totalScore += score
	}
	return totalScore
}

func missingBracketsScore(input []string) int {
	score := 0
	/*
		): 1 point.
		]: 2 points.
		}: 3 points.
		>: 4 points.
	*/
	bracketScore := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	for _, c := range input {
		score *= 5
		score += bracketScore[c]
	}
	return score
}

func part2(input []string) int {
	var scores []int
	var missingBrackets []string
	for _, line := range input {
		_, missingBrackets = parseLine(line)
		score := missingBracketsScore(missingBrackets)
		if score > 0 {
			scores = append(scores, missingBracketsScore(missingBrackets))
		}
	}
	sort.Ints(scores)
	return scores[22]
}
