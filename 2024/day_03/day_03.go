package main

import (
	"fmt"
	"github.com/0x5ubt13/Advent_of_code/helper_functions_go"
	"strconv"
)

func main() {
	data := helper_functions_go.GetInputAsStringArray("2024/day_03/day_03_test_input.txt")
	goodStrings := make([]string, 0)
	for _, v := range data {
		fmt.Println(v)
		skipOne := 0
		for index, letterRune := range v {
			if skipOne > 0 {
				skipOne--
				continue
			}

			letter := string(letterRune)
			fmt.Println(index, letter)

			if letter == "x" {
				if v[index+1] == "u" {
					skipOne++
					if v[index+2] == "l" {
						skipOne++
						if v[index+3] == "(" {
							skipOne++
							if isNumber(v[index+4]) {
								skipOne++
								// next can be either number or comma. Implement comma
								if isNumber(v[index+5]) {
									skipOne++
									// next can be either number or comma. Implement comma
									if isNumber(v[index+6]){
										skipOne++
										// next can only be a comma

									}

								} else if
							} else {
								continue
							}
						}
					}
				}
			}

		}
	}
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
