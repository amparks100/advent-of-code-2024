package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Merry Holiday Time!")

	day1(day1input)
}

func day1(input string) {

	var leftValues []int
	var rightValues []int
	var totalDistance = 0.0

	lines := strings.Split(input, "\n")
	for _, line := range lines {

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, "   ", 2)
		if len(parts) == 2 {
			left := strings.TrimSpace(parts[0])
			lnum, _ := strconv.Atoi(left)
			leftValues = append(leftValues, lnum)

			right := strings.TrimSpace(parts[1])
			rnum, _ := strconv.Atoi(right)
			rightValues = append(rightValues, rnum)
			// fmt.Printf("left: %s, right: %s\n", left, right)
		}
	}

	sort.Ints(leftValues)
	sort.Ints(rightValues)

	fmt.Println("Sorted lefts:", leftValues)
	fmt.Println("Sorted rights:", rightValues)

	for index := range leftValues {
		distance := math.Abs(float64(leftValues[index] - rightValues[index]))
		fmt.Println("Adding Distance:", distance)
		totalDistance += distance
	}

	fmt.Println("Total Distance:", totalDistance)

	// part 2 - similarity score
	// using left lift, for each item in left, multiply by number of times that number appears in right list
	var similarityScore = 0.0

	for _, lValue := range leftValues {
		countInRight := 0
		for _, rValue := range rightValues {
			if rValue == lValue {
				countInRight++
			}
		}
		similarityScore += float64(lValue * countInRight)
	}

	fmt.Println("Similarity Score:", similarityScore)

}
