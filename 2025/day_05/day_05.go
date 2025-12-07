package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
)

func main() {
	// input := helper_functions_go.GetInputAsStringArray("./day_05_test_input.txt")
	input := helper_functions_go.GetInputAsStringArray("./day_05_input.txt")

	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))
}

func part1(input []string) int {
	result := 0
	// fmt.Printf("%s\n", input)

	ranges, ingredientIDs := GetRangesAndIngredientIDs(input, false)

	valids := []int{}
	for _, r := range ranges {
		rInt := strings.Split(r, "-")
		start, err := strconv.Atoi(rInt[0])
		helper_functions_go.ErrorCheck(err)
		end, err := strconv.Atoi(rInt[1])
		helper_functions_go.ErrorCheck(err)

		for _, ingredientID := range ingredientIDs {

			if ingredientID >= start && ingredientID <= end && !slices.Contains(valids, ingredientID) {
				valids = append(valids, ingredientID)
				result += 1
			}
		}
	}

	return result
}

type Range struct {
	start int
	end   int
}

func part2(input []string) int {
	rangeStrings, _ := GetRangesAndIngredientIDs(input, true)

	// Parse ranges
	ranges := make([]Range, 0, len(rangeStrings))
	for _, r := range rangeStrings {
		rInt := strings.Split(r, "-")
		start, err := strconv.Atoi(rInt[0])
		helper_functions_go.ErrorCheck(err)
		end, err := strconv.Atoi(rInt[1])
		helper_functions_go.ErrorCheck(err)
		ranges = append(ranges, Range{start, end})
	}

	// Sort ranges by start position
	slices.SortFunc(ranges, func(a, b Range) int {
		return a.start - b.start
	})

	// Merge overlapping ranges and calculate total count
	result := 0
	if len(ranges) == 0 {
		return 0
	}

	currentStart := ranges[0].start
	currentEnd := ranges[0].end

	for i := 1; i < len(ranges); i++ {
		if ranges[i].start <= currentEnd+1 {
			// Ranges overlap or are adjacent, merge them
			if ranges[i].end > currentEnd {
				currentEnd = ranges[i].end
			}
		} else {
			// No overlap, count the current range and start a new one
			result += currentEnd - currentStart + 1
			currentStart = ranges[i].start
			currentEnd = ranges[i].end
		}
	}

	// Add the last range
	result += currentEnd - currentStart + 1

	return result
}

func GetRangesAndIngredientIDs(input []string, isPart2 bool) ([]string, []int) {
	ranges := []string{}
	readingRanges := true
	ingredientIDs := []int{}

	for _, line := range input {
		if line == "" {
			if isPart2 {
				return ranges, nil
			}
			readingRanges = false
			continue
		}

		if readingRanges {
			ranges = append(ranges, line)
			continue
		}

		ingredientID, err := strconv.Atoi(line)
		helper_functions_go.ErrorCheck(err)
		ingredientIDs = append(ingredientIDs, ingredientID)
	}

	return ranges, ingredientIDs
}
