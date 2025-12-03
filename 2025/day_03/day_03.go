package main

import (
	"fmt"
	"strconv"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	input := helper_functions_go.GetInputAsIntGrid("./day_03_test_input.txt")
	// input := helper_functions_go.GetInputAsIntGrid("./day_03_input.txt")

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
	targetLength := 12
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
	
	return result
}
