package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	//"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
	"io"
	"os"
)

func main() {
	input := GetInputAsStringArray("2024/day_05/test_input.txt")
	fmt.Println(input)

	pagesMap := make(map[int][]string)
	sectionOne := true
	result := 0

	for i, pages := range input {
		if pages == "" {
			fmt.Println("empty line detected. Initiating second phase")
			sectionOne = false
			continue
		}

		// Classify pages
		if sectionOne {
			fmt.Println("index:", i, "pages:", pages)
			// get page numbers in [0] and [1]
			pagesMap[i] = strings.Split(pages, "|")

			continue
		}

		// Check order
		correctOrder := true
		pagesOrder := strings.Split(pages, ",")

		middlePage := len(pagesOrder) / 2
		for j, pageToOrder := range pagesOrder {
			fmt.Println("Checking order of page", j, pageToOrder)
			for k := range len(pagesMap) {
				fmt.Println(pagesMap[k])
				// To complete
				if pagesMap[k][0] == pageToOrder {
					fmt.Println("Page to order: ", pageToOrder)
				}
			}

			correctOrder = false
		}

		// Add middlePage to the result
		if correctOrder {
			n, e := strconv.Atoi(pagesOrder[middlePage])
			ErrorCheck(e)

			result += n
		}
	}

	//for i := range len(pagesMap) {
	//	fmt.Println(pagesMap[i])
	//}
}

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

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
