package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"

	//"io"
	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
	"os"
)

func main() {
	left, right := getSortedInput("./day_01_input.txt")
	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d", part1(left, right), part2(left, right))
}

func part2(left, right []int) int {
	occurrences := make(map[int]int)

	for _, leftNum := range left {
		// Skip if already checked
		if _, ok := occurrences[leftNum]; ok {
			continue
		}

		for _, rightNum := range right {
			if leftNum == rightNum {
				occurrences[leftNum]++
			}
		}
	}

	result := 0
	for _, leftNum := range left {
		result += leftNum * occurrences[leftNum]
	}

	return result
}

func part1(left, right []int) int {
	result := 0
	for i := range left {
		if left[i] > right[i] {
			result += left[i] - right[i]
		}

		if left[i] < right[i] {
			result += right[i] - left[i]
		}

		i++
	}

	return result
}

func getSortedInput(filename string) ([]int, []int) {
	f, err := os.Open(filename)
	helper_functions_go.ErrorCheck(err)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	left, right := make([]int, 0), make([]int, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//fmt.Println(strings.Split(strings.ReplaceAll(scanner.Text(), "   ", " "), " "))
		line := strings.Split(scanner.Text(), " ")
		leftNum, _ := strconv.Atoi(line[0])
		rightNum, _ := strconv.Atoi(line[len(line)-1])

		//fmt.Println(leftNum, rightNum)

		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	// Sort them
	sort.Ints(left)
	sort.Ints(right)

	return left, right
}
