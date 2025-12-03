package main

import (
	"fmt"
	"strconv"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	// input := helper_functions_go.GetInputAsIntGrid("./day_03_test_input.txt")
	input := helper_functions_go.GetInputAsIntGrid("./day_03_input.txt")

	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))
}

func part1(input [][]int) int {
	result := 0
	// fmt.Printf("%v\n", input)
	for _, row := range input {
		greatest := 0
		for i, column := range row {
			first := column

			for j, second := range row {
				if j <= i {
					continue
				}

				test, err := strconv.Atoi(fmt.Sprintf("%d%d", first, second))
				helper_functions_go.ErrorCheck(err)
				if test > greatest {
					greatest = test
				}
			}
		}
		// fmt.Printf("Greatest in row: %d\n", greatest)
		result += greatest
	}
	return result
}

func part2(input [][]int) int {
	result := 0
	// fmt.Printf("%v\n", input)

	for _, row := range input {
		// fmt.Printf("Row: %v\n", row)

		// Building it recursively
		largestJoltage := findLargestJoltage(row, 0, 12)
		
		// fmt.Println("Largest joltage found in row ", row, ": ", largestJoltage)
		largestJoltageInt, err := strconv.Atoi(largestJoltage)
		helper_functions_go.ErrorCheck(err)
		result += largestJoltageInt 
	}

	return result
}

// findLargestJoltage finds recursively the greatest number possible
func findLargestJoltage(row []int, startIdx, digitsNeeded int) string {
	// Base case, exit
	if digitsNeeded == 0 {
		return ""
	}

	// Calculate how far we can still search
	// If we need still 10 numbers, look as far as the maximum length - 9
	canLookUntil := len(row) - (digitsNeeded -1)

	// Find the largest digit in the valid range established
	maxDigit := 0
	maxIdx := startIdx
	for i := startIdx; i < canLookUntil; i++ {
		if row[i] > maxDigit {
			maxDigit = row[i]
			maxIdx = i
		}
	}

	// Recursively build the rest of the number
	return fmt.Sprintf("%d", maxDigit) + findLargestJoltage(row, maxIdx+1, digitsNeeded-1)
}