package main

import (
	"fmt"
	// "strings"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	// input := helper_functions_go.GetInputAsCharGrid("./day_04_test_input.txt")
	input := helper_functions_go.GetInputAsCharGrid("./day_04_input.txt")

	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))
}

func part1(input [][]rune) int {
	result := 0

	coordinatesArray := makeCoordinatesArray(input, len(input), len(input[0]))

	// Now figure out the results
	for _, c := range coordinatesArray {
		// fmt.Printf("row: %v, column: %v, score: %d, value: %c\n", c.Y, c.X, c.Score, c.Value)
		if c.Value == '@' && c.Score < 4 {
			// fmt.Printf("GONER\n")
			result += 1
		}
	}

	return result
}

func part2(input [][]rune) int {
	result := 0
	
	coordinatesArray := makeCoordinatesArray(input, len(input), len(input[0]))

	control := 0
	for {
		for i := range coordinatesArray {
			// fmt.Printf("row: %v, column: %v, score: %d, value: %c\n", c.Y, c.X, c.Score, c.Value)
			if coordinatesArray[i].Value == '@' && coordinatesArray[i].Score < 4 {
				// fmt.Printf("GONER\n")
				result += 1
				// Update the grid
				input[coordinatesArray[i].Y][coordinatesArray[i].X] = '.'
				control += 1
			}
		}
		
		if control == 0 {
			break
		}

		// Recalculate coords
		coordinatesArray = makeCoordinatesArray(input, len(input), len(input[0]))
		control = 0
	}

	return result
}

// Coordinates will quickly assign coords
type Coordinates struct {
	X     int
	Y     int
	Score int
	Value rune
	North rune
	South rune
	East  rune
	West  rune
	NE    rune
	NW    rune
	SE    rune
	SW    rune
}

// checkAdjacents will look at the 8 adjacent positions in the grid
// and return the value of detected matches
func checkAdjacents(grid [][]rune, gridHeight, gridWidth, x, y int) Coordinates {
	c := Coordinates{
		X:     x,
		Y:     y,
		Score: 0,
		Value: grid[y][x],
	}

	if c.Value == '.' {
		return c
	}

	// Helper to avoid array boundary errors
	get := func(x, y int) rune {
		// fmt.Printf("Checking x: %v, y: %v, value: ", x, y)
		if x < 0 || y < 0 || x >= gridWidth || y >= gridHeight {
			fmt.Printf("|\n")
			// Sentinel value
			return '|'
		}
		// fmt.Printf("%c\n", grid[y][x])
		return grid[y][x]
	}

	c.North = get(x, y-1)
	if c.North == '@' {
		c.Score += 1
	}

	c.South = get(x, y+1)
	if c.South == '@' {
		c.Score += 1
	}

	c.East = get(x+1, y)
	if c.East == '@' {
		c.Score += 1
	}

	c.West = get(x-1, y)
	if c.West == '@' {
		c.Score += 1
	}

	c.NE = get(x+1, y-1)
	if c.NE == '@' {
		c.Score += 1
	}

	c.NW = get(x-1, y-1)
	if c.NW == '@' {
		c.Score += 1
	}

	c.SE = get(x+1, y+1)
	if c.SE == '@' {
		c.Score += 1
	}

	c.SW = get(x-1, y+1)
	if c.SW == '@' {
		c.Score += 1
	}

	return c
}

func makeCoordinatesArray(input [][]rune, gridHeight, gridWidth int) []Coordinates {
	coordinatesArray := []Coordinates{}
	for rowIdx, line := range input {
		// fmt.Printf("%v", line)

		for colIdx := range line {
			// fmt.Printf("Col: %c\n", col)
			// fmt.Printf("Getting data for rowIdx %v and colIdx %v, with value %c\n", rowIdx, colIdx, col)
			coordinatesArray = append(coordinatesArray, checkAdjacents(input, gridHeight, gridWidth, colIdx, rowIdx))
		}

		// fmt.Println(coordinatesArray)
	}
	return coordinatesArray
}