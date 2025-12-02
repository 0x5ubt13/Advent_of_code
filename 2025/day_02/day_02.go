package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := GetInputAsStringArrayFromCSVWithDelimiter("./day_02_input.txt", ",")
	fmt.Printf("Part 1 solution: %d\nPart 2 solution: %d\n", part1(input), part2(input))
}

func part1(input []string) int {
	result := 0

	// fmt.Printf("Input: %v\n", input)

	for _, line := range input {
		// fmt.Printf("Line: %s\n", line)

		// Check every line (puzzle range) for any duplicates
		// Divide length by 2 and check if numbers are their doubles
		startStr, endStr := strings.Split(line, "-")[0], strings.Split(line, "-")[1]

		start, err := strconv.Atoi(startStr)
		ErrorCheck(err)
		end, err := strconv.Atoi(endStr)
		ErrorCheck(err)

		// fmt.Printf("Start: %d, End: %d\n", start, end)

		for i := start; i <= end; i++ {
			iStr := strconv.Itoa(i)
			iLen := len(iStr)
			firstHalf := iStr[:iLen/2]
			secondHalf := iStr[iLen/2:]

			if firstHalf == secondHalf {
				// fmt.Printf("Found double: %s\n", iStr)
				result += i
			}
		}

	}

	return result
}

func part2(input []string) int {
	result := 0

	// fmt.Printf("\n\n--- Part 2 ---\n\n")

	for _, line := range input {
		startStr, endStr := strings.Split(line, "-")[0], strings.Split(line, "-")[1]

		start, err := strconv.Atoi(startStr)
		ErrorCheck(err)
		end, err := strconv.Atoi(endStr)
		ErrorCheck(err)

		for i := start; i <= end; i++ {
			currentNumber := strconv.Itoa(i)
			currentNumberDigits := len(currentNumber)

			// Skip single-digit numbers (can't be repeated patterns)
			if currentNumberDigits < 2 {
				continue
			}

			found := false

			// Try all possible divisions: 2 parts, 3 parts, 4 parts, etc.
			for parts := 2; parts <= currentNumberDigits; parts++ {
				// Only check if the length divides evenly
				if currentNumberDigits%parts != 0 {
					continue
				}

				partLen := currentNumberDigits / parts
				firstPart := currentNumber[:partLen]
				allEqual := true

				// Check if all parts are equal to the first part
				for p := 1; p < parts; p++ {
					nextPart := currentNumber[p*partLen : (p+1)*partLen]
					if firstPart != nextPart {
						allEqual = false
						break
					}
				}

				if allEqual {
					// fmt.Printf("Found repeated pattern: %s (pattern: %s, %d times)\n", currentNumber, firstPart, parts)
					result += i
					found = true
					break
				}
			}

			if found {
				continue
			}
		}
	}

	return result

	// 24573450272 too low
}

// GetInputAsString reads entire file as a single string
func GetInputAsString(filename string) string {
	bytes, err := os.ReadFile(filename)
	ErrorCheck(err)
	return string(bytes)
}

// GetInputAsStringArrayFromCSVWithDelimiter reads a single line with custom delimiter and returns an array of strings
func GetInputAsStringArrayFromCSVWithDelimiter(filename string, delimiter string) []string {
	content := GetInputAsString(filename)
	content = strings.TrimSpace(content)
	parts := strings.Split(content, delimiter)

	result := make([]string, len(parts))
	for i, part := range parts {
		partTrimmed := strings.TrimSpace(part)
		result[i] = partTrimmed
	}
	return result
}

// Condensing error handling
func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
