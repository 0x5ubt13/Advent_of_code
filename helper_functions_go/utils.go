package helper_functions_go

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInputAsInt64Array(filename string) []int64 {
	lines := make([]int64, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		intLine, err2 := strconv.ParseInt(scanner.Text(), 10, 64)
		if err2 != nil {
			log.Fatal(err2)
		}

		lines = append(lines, intLine)
	}
	return lines
}

func GetInputAsArrayOfInt64Arrays(filename string) [][]int64 {
	lines := make([][]int64, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// 7 6 4 2 1
		numbersLine := strings.Split(scanner.Text(), " ")
		int64line := make([]int64, len(numbersLine))
		for i := 0; i < len(numbersLine); i++ {
			int64line[i], err = strconv.ParseInt(numbersLine[i], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}

		lines = append(lines, int64line)
	}
	return lines
}

func GetInputAsMapStringInt(filename string) []map[string]int {
	file, err := os.Open(filename)
	ErrorCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var directions []map[string]int
	for _, val := range lines {
		tempSlice := strings.Fields(val)
		key := tempSlice[0]
		value, err2 := strconv.Atoi(tempSlice[1])
		ErrorCheck(err2)
		m := map[string]int{key: value}
		directions = append(directions, m)
	}

	return directions
}

func GetInputAsArrayOfStringArrays(filename string) [][]string {
	lines := make([][]string, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), " "))
	}
	return lines
}

// GetInputAsStringArray processes input and returns []string
func GetInputAsStringArray(filename string) []string {
	f, err := os.Open(filename)
	ErrorCheck(err)
	defer f.Close()

	data := make([]string, 0)

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		data = append(data, string(line))
	}

	return data
}

// GetInputAsByteArray processes input and returns []byte
func GetInputAsByteArray(filename string) []byte {
	bytes, err := os.ReadFile(filename)
	ErrorCheck(err)
	return bytes
}

// GetInputAsString reads entire file as a single string
func GetInputAsString(filename string) string {
	bytes, err := os.ReadFile(filename)
	ErrorCheck(err)
	return string(bytes)
}

// GetInputAsCharGrid reads file as 2D rune array (character grid)
// Useful for grid-based puzzles like mazes, maps, etc.
func GetInputAsCharGrid(filename string) [][]rune {
	lines := GetInputAsStringArray(filename)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}

// GetInputAsIntArray reads file where each line is a single integer
func GetInputAsIntArray(filename string) []int {
	lines := make([]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		intLine, err2 := strconv.Atoi(scanner.Text())
		if err2 != nil {
			log.Fatal(err2)
		}
		lines = append(lines, intLine)
	}
	return lines
}

// GetInputAsArrayOfIntArrays reads file where each line contains space-separated integers
func GetInputAsArrayOfIntArrays(filename string) [][]int {
	lines := make([][]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		numbersLine := strings.Fields(scanner.Text())
		intline := make([]int, len(numbersLine))
		for i := 0; i < len(numbersLine); i++ {
			intline[i], err = strconv.Atoi(numbersLine[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		lines = append(lines, intline)
	}
	return lines
}

// GetInputAsIntArrayFromCSV reads a single line of comma-separated integers
// Example: "1,2,3,4,5" -> []int{1,2,3,4,5}
func GetInputAsIntArrayFromCSV(filename string) []int {
	content := GetInputAsString(filename)
	content = strings.TrimSpace(content)
	parts := strings.Split(content, ",")

	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		ErrorCheck(err)
		result[i] = num
	}
	return result
}

// GetInputAsIntArrayFromCSVWithDelimiter reads a single line with custom delimiter
func GetInputAsIntArrayFromCSVWithDelimiter(filename string, delimiter string) []int {
	content := GetInputAsString(filename)
	content = strings.TrimSpace(content)
	parts := strings.Split(content, delimiter)

	result := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		ErrorCheck(err)
		result[i] = num
	}
	return result
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

// GetInputAsParagraphs reads file split by blank lines
// Returns [][]string where each []string is a paragraph (group of lines)
// Useful for puzzles with grouped data
func GetInputAsParagraphs(filename string) [][]string {
	content := GetInputAsString(filename)
	paragraphs := strings.Split(content, "\n\n")

	result := make([][]string, len(paragraphs))
	for i, paragraph := range paragraphs {
		lines := strings.Split(strings.TrimSpace(paragraph), "\n")
		result[i] = lines
	}
	return result
}

// GetInputAsIntGrid reads file as 2D int array where each character is a digit
// Example: "123\n456" -> [][]int{{1,2,3},{4,5,6}}
// Useful for height maps, digit grids, etc.
func GetInputAsIntGrid(filename string) [][]int {
	lines := GetInputAsStringArray(filename)
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}
	return grid
}

// Condensing error handling
func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
