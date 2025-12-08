package main

import (
	"fmt"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	input := helper_functions_go.GetInputAsStringArray("./day_07_input.txt")
	// input := helper_functions_go.GetInputAsStringArray("./day_07_test_input.txt")

	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))
}

func part1(input []string) int {
	result := 0 // Count only the number of splits (not the initial segment)

	for i, line := range input {
		fmt.Printf("Line %d: %s\n", i, line)
	}

	fmt.Println("")

	// Create a deep copy of the grid (convert strings to []rune for modification)
	shadowGrid := make([][]rune, len(input))
	for i := range input {
		shadowGrid[i] = []rune(input[i])
	}

	for i, line := range input {
		for j, col := range line {
			switch col {
			case 'S':
				shadowGrid[i+1][j] = '|'
			case '.':
				if i != 0 && i != len(input)-1 {
					if shadowGrid[i-1][j] == '|' {
						shadowGrid[i][j] = '|'
					}
				}
			case '^':
				// Check if water is flowing into this position from above
				if i > 0 && shadowGrid[i-1][j] == '|' {
					// Each ^ that has water flowing through it creates ONE additional segment
					// (splits 1 segment into 2, so adds 1 to the total count)
					result++
					fmt.Printf("Counting ^ at row %d, col %d (result now: %d)\n", i, j, result)
					shadowGrid[i][j-1] = '|'
					shadowGrid[i][j+1] = '|'
				} else {
					fmt.Printf("NOT counting ^ at row %d, col %d (no water above)\n", i, j)
				}
			}
		}
	}

	for idx, line := range shadowGrid {
		fmt.Printf("Row %2d: %s\n", idx, string(line))
	}

	// 1752 too high
	// 1751 too high

	return result
}

func part2(input []string) int {
	result := 0
	// Implementation
	return result
}
