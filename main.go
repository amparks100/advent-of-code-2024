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

	day2(day2input)
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

func day2(input string) {
	reports := strings.Split(input, "\n")
	var totalSafe = 0
	for _, report := range reports {

		report = strings.TrimSpace(report)
		if report == "" {
			continue
		}
		fmt.Println("Report:", report)
		reportSplit := strings.Fields(report)
		var levels []int
		for _, level := range reportSplit {
			num, _ := strconv.Atoi(level)
			levels = append(levels, num)
		}

		countIncreasing := 0
		countDecreasing := 0
		safe := true
		importantNumber := len(levels) - 1
		for index, _ := range levels {

			if index == importantNumber {
				continue
			}
			distance := levels[index] - levels[index+1]
			if distance > 0 {
				countDecreasing += 1
			} else {
				countIncreasing += 1
			}
		}
		for index, _ := range levels {

			if index == importantNumber {
				continue
			}
			distance := levels[index] - levels[index+1]
			fmt.Println("Distance:", distance)
			distance = AbsInt(distance)
			if distance >= 1 && distance <= 3 {
				safe = true
				fmt.Println("safe distance")
			} else {
				fmt.Println("Report not safe! - too distant")
				safe = false
				break
			}
		}
		if !safe {
			continue
		}
		fmt.Println("Count decreasing:", countDecreasing)
		fmt.Println("Count increasing:", countIncreasing)
		if (countDecreasing == importantNumber) || (countIncreasing == importantNumber) {
			fmt.Println("Report safe! - all same")
			safe = true
		} else {
			safe = false
		}
		if safe {
			fmt.Println("Safe report!")
			totalSafe += 1
		}
	}
	fmt.Println("Total Safe:", totalSafe)
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
