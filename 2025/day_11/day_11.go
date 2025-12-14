package main

import (
	"fmt"
	"strings"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	input := helper_functions_go.GetInputAsArrayOfStringArrays("./day_11_test_input.txt")
	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))
}

func part1(input [][]string) int {
	result := 0
	// Map every first element of each line array to the second and third elements of the same line array
	stringMap := make(map[string][]string)
	queue := []string{}
	
	for _, line := range input {
		// Take away the colon from the first element
		line[0] = strings.TrimSuffix(line[0], ":")
		queue = append(queue, line[0])
		for _, str := range line {
			if str != line[0] {
				stringMap[line[0]] = append(stringMap[line[0]], str)
			}
		}
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, str := range stringMap[curr] {
				
		}

	fmt.Printf("String Map: %v\n", stringMap)
	return result
}

func part2(input [][]string) int {
	result := 0
	// Implementation
	return result
}
