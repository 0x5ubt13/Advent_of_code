package main

import (
	"fmt"
	"strconv"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	// input := helper_functions_go.GetInputAsArrayOfStringArrays("./day_06_test_input.txt")
	// inputLines := helper_functions_go.GetInputAsStringArray("./day_06_test_input.txt")
	input := helper_functions_go.GetInputAsArrayOfStringArrays("./day_06_input.txt")
	inputLines := helper_functions_go.GetInputAsStringArray("./day_06_input.txt")

	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(inputLines))
}

func part1(input [][]string) int {
	result := 0
	orderedOperations := CleanOutput(input)
	
	for _, op := range orderedOperations {
		opResult := 0
		switch op.Operand {
		case "+":
			for _, num := range op.Numbers {
				opResult += num
			}
		case "*":
			opResult = 1
			for _, num := range op.Numbers {
				opResult *= num
			}
		}
		result += opResult
	}
	
	return result
}

func part2(inputLines []string) int {
	result := 0
	orderedOperations := CleanOutputPart2(inputLines)

	for _, op := range orderedOperations {
		opResult := 0
		switch op.Operand {
		case "+":
			for _, num := range op.Numbers {
				opResult += num
			}
		case "*":
			opResult = 1
			for _, num := range op.Numbers {
				opResult *= num
			}
		}
		result += opResult
	}

	return result
}

type Operations struct {
	Numbers []int
	Operand string
}

func CleanOutput(input [][]string) map[int]Operations {
	ops := make(map[int]Operations)
	for i, line := range input {
		counter := 0
		for _, content := range line {
			if content == "" {
				continue
			}

			if content == "*" || content == "+" {
				op := ops[counter]
				op.Operand = content
				ops[counter] = op
				counter += 1
				continue
			}

			num, err := strconv.Atoi(content)
			helper_functions_go.ErrorCheck(err)

			if i == 0 {
				ops[counter] = Operations{Numbers: []int{num}}
			} else {
				op := ops[counter]
				op.Numbers = append(op.Numbers, num)
				ops[counter] = op
			}
			counter += 1
		}
	}
	return ops
}

func CleanOutputPart2(inputLines []string) map[int]Operations {
	// Convert input lines to character grid, preserving spaces
	maxLen := 0
	for _, line := range inputLines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Build character grid from raw lines
	grid := make([][]rune, len(inputLines))
	for i, line := range inputLines {
		grid[i] = make([]rune, maxLen)
		for j, ch := range line {
			grid[i][j] = ch
		}
		// Fill remaining with spaces if line is shorter
		for j := len(line); j < maxLen; j++ {
			grid[i][j] = ' '
		}
	}

	ops := make(map[int]Operations)
	opCounter := 0

	// Iterate columns from right to left
	for col := maxLen - 1; col >= 0; col-- {
		var numRunes []rune
		operand := ""
		isEmpty := true

		for row := 0; row < len(grid); row++ {
			ch := grid[row][col]

			if ch == ' ' {
				continue
			}

			isEmpty = false

			// Check if it's an operator
			if ch == '*' || ch == '+' {
				operand = string(ch)
			} else {
				// It's a digit
				numRunes = append(numRunes, ch)
			}
		}

		// If this column was all spaces, skip it
		if isEmpty {
			continue
		}

		// Convert runes to number
		if len(numRunes) > 0 {
			numStr := string(numRunes)
			num, err := strconv.Atoi(numStr)
			helper_functions_go.ErrorCheck(err)

			op := ops[opCounter]
			op.Numbers = append(op.Numbers, num)
			ops[opCounter] = op
		}

		// If we found an operator, move to next operation
		if operand != "" {
			op := ops[opCounter]
			op.Operand = operand
			ops[opCounter] = op
			opCounter++
		}
	}

	return ops
}
