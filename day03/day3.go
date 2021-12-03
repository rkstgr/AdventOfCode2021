package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	// Input reading
	bs, err := ioutil.ReadFile("input.txt")

	if err != nil {
		panic("could not read")
	}

	input := strings.Split(string(bs), "\n")
	input = input[:len(input)-1]

	sol1 := part1(input)
	sol2 := part2(input)


	end := time.Since(start)

	fmt.Println("The solution to part 1 is: ", sol1)
	fmt.Println("The solution to part 2 is: ", sol2)
	fmt.Println("Time: ", end)

}


func part1(input []string) int64 {
	var x string
	// i := current column
	for i := 0; i < len(input[0]); i++ {
		count0 := 0
		count1 := 0
		for n := 0; n < len(input); n++ {
			if input[n][i] == '0' {
				count0++
			} else {
				count1++
			}
		}
		if count0 > count1 {
			x += "0"
		} else {
			x += "1"
		}
	}
	gammaRate, _ := strconv.ParseInt(x, 2, 16)
	epsilonRate := int64(math.Pow(2, float64(len(x)))) - gammaRate - 1
	return gammaRate * epsilonRate

}


func part2(input []string) int64 {
	col := 0
	result := input

	// oxygen
	for len(result)>1 {
		result = keepCommon(result, col, true)
		col++
	}
	oxygenRating, _ := strconv.ParseInt(result[0], 2, 16)


	col = 0
	result = input
	// co2
	for len(result)>1 {
		result = keepCommon(result, col, false)
		col++
	}
	co2Rating, _ := strconv.ParseInt(result[0], 2, 16)

	return oxygenRating * co2Rating

}


func keepCommon(input []string, col int, mostCommon bool) []string {
	var output [1000]string
	
	common := 0

	count0 := 0
	count1 := 0

	// find the common bit
	for n := 0; n<len(input); n++ {
		if input[n][col] == '0' {
			count0++
		} else {
			count1++
		}
	}
	
	if count0 > count1 {
		common = 0
	} else {
		common = 1
	}

	if !mostCommon {
		common = (common+1) % 2
	}
	commonString := strconv.Itoa(common)[0]

	// iterate over all columns and keep those that correspond to common
	j := 0
	for n := 0; n<len(input); n++ {
		if input[n][col] == commonString {
			output[j] = input[n]
			j++
		}
	}

	return output[:j]

}