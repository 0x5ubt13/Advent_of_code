package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	input := helper_functions_go.GetInputAsArrayOfStringArrays("./day_08_test_input.txt")
	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))
}

func part1(input [][]string) int {
	result := 0
	
	for _, row := range input {
		for _, num := range row {
			// X,Y,Z coordinates
			coordinates := strings.Split(num, ",")
			x, err := strconv.Atoi(coordinates[0])
			helper_functions_go.ErrorCheck(err)
			y, err := strconv.Atoi(coordinates[1])
			helper_functions_go.ErrorCheck(err)
			z, err := strconv.Atoi(coordinates[2])
			helper_functions_go.ErrorCheck(err)
			fmt.Println(x, y, z)
		}
	}

	return result
}

func part2(input [][]string) int {
	result := 0
	// Implementation
	return result
}
