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

// Condensing error handling
func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
