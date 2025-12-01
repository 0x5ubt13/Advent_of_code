package main

import (
	"fmt"
	"strconv"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	input := helper_functions_go.GetInputAsStringArray("./day_01_input.txt")

	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))

	
}

func part1(input []string) int {
	result := 0
	position := 50
	for _, line := range input {
		direction := rune(line[0])
		number, err := strconv.Atoi(line[1:])
		helper_functions_go.ErrorCheck(err)

		switch direction {
		case 'L':
			position = position - (number % 100)
			if position < 0 {
				position += 100
			}
		case 'R':
			position += number
			position %= 100
		}

		if position == 0 {
			result++
		}

	}
	return result
}

func part2(input []string) int {
	result := 0
	position := 50
	for _, line := range input {
		direction := rune(line[0])
		number, err := strconv.Atoi(line[1:])
		helper_functions_go.ErrorCheck(err)

		oldPos := position
		var newPos int

		if direction == 'L' {
			newPos = position - number
		} else {
			newPos = position + number
		}

		// Simple logic: count how many multiples of 100 we pass through or land on
		// Multiples of 100 represent position 0

		if oldPos == 0 {
			// Starting at 0 - don't count the start, only if we come back to 0
			if newPos == 0 {
				// Didn't move (shouldn't happen)
			} else if newPos > 0 {
				// Went right: count multiples at 100, 200, etc
				result += newPos / 100
				// Also check if we landed back on 0
				if newPos % 100 == 0 {
					// We're at 0 after wrapping (e.g., R100, R200, etc.)
					// This is already counted in newPos / 100
				}
			} else {
				// Went left: count multiples at -100, -200, etc, then back to 0
				result += (-newPos) / 100
				// Check if we landed back on 0
				if newPos % 100 == 0 {
					// Landed on 0 (e.g., L100, L200)
					// Need to add 1 more for landing on 0
					result += 1
				}
			}
		} else {
			// Not starting at 0
			if newPos >= 100 {
				// Went right past 100
				result += newPos / 100
			} else if newPos > 0 && newPos < 100 {
				// Stayed in positive range
				if newPos == 0 || (oldPos > newPos && newPos == 0) {
					// Landed on 0
					result++
				}
			} else if newPos == 0 {
				// Landed exactly on 0
				result++
			} else {
				// newPos < 0: went negative
				result += (-newPos) / 100 + 1
			}
		}

		position = ((newPos % 100) + 100) % 100
	}

	return result
}
