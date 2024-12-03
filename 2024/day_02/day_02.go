package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := GetInputAsArrayOfInt64Array(`2024/day_02/day_03_input.txt`)
	//fmt.Println(data)

	fmt.Printf("Part 1: %d\nPart 2: %d\n", part1(data), part2(data))
}

func part2(data [][]int64) int64 {
	var difference, safeLines int64
	for _, line := range data {
		increase, decrease, skipOne, penultimate := false, false, false, false
		for i, num := range line {
			if skipOne {
				continue
			}

			fmt.Println(i, num)

			// First number, check increase or decrease
			if i == 0 {
				if line[i+1] > num {
					increase = true
					if line[i+1]-num <= 3 && line[i+1]-num >= 0 {
						// 1 4 7 10 13 16
						//fmt.Println(line, "increases")
						continue
					} else {
						// 1 5, or
						// 5 1
						skipOne = true
						// second number not good, check third number
						if line[i+2] > num {
							increase = true
							if line[i+1]-num <= 3 && line[i+1]-num >= 0 {
								continue
							}
						}
						// Not good second and third. break out
						break
					}
				} else if line[i+1] < num {
					decrease = true
					// 11 9 7 5 3 1
					if num-line[i+1] <= 3 && num-line[i+1] >= 0 {
						continue
					}
					//fmt.Println(line, "decreases")
				} else {
					//fmt.Println(line, "unsafe")
					// part 2: ignore it instead and try next one
					if line[i+2] > num {
						// 1 1 4 7 10 13 16
						increase, skipOne = true, true
						//fmt.Println(line, "increases")
						if line[i+2]-num <= 3 && line[i+2]-num >= 0 {
							continue
						}
						// else break, 2 in a row
						break
					} else if line[i+2] < num {
						decrease, skipOne = true, true
						if num-line[i+2] <= 3 && num-line[i+2] >= 0 {
							// 16 16 13 10 7 4 1
							//fmt.Println(line, "decreases")
							continue
						} else {
							// else break, 2 in a row
							break
						}
					}
				}

				// Middle numbers, compare against others
				if i >= 1 && i < len(line)-1 {
					if increase {
						difference = line[i+1] - num
						fmt.Println("num:", num, "num+1:", line[i+1], "difference:", difference)
					}

					if decrease {
						difference = num - line[i+1]
						fmt.Println("num:", num, "num+1:", line[i+1], "difference:", difference)
					}

					if difference >= 4 || difference <= 0 {
						fmt.Println(num, line[i+1], difference, "detected unsafe. breaking")
						// Unsafe detected, check next one if not one borked already
						if skipOne {
							break
						}

						if i+1 == len(line)-1 {
							// Edge case: last chance to check if all good if everything went good until now and penultimate
							penultimate = true
						}

						if increase && !penultimate {
							difference = line[i+2] - num
							fmt.Println("num:", num, "num+1:", line[i+2], "difference:", difference)
							continue
						}

						if decrease && !penultimate {
							difference = num - line[i+2]
							fmt.Println("num:", num, "num+1:", line[i+2], "difference:", difference)
							continue
						}

						skipOne = true
					}
				}

				// Last number
				if i == len(line)-1 {
					fmt.Println(line, "is safe")
					safeLines += 1
				}
			}
		}
	}

	return safeLines
}

func part1(data [][]int64) int64 {
	var difference, safeLines int64
	for _, line := range data {
		increase, decrease := false, false
		for i, num := range line {
			fmt.Println(i, num)
			if i == 0 {
				// First number, check increase or decrease
				if line[i+1] > num && line[i+1]-num <= 3 && line[i+1]-num >= 0 {
					increase = true
					//fmt.Println(line, "increases")
				} else if line[i+1] < num && num-line[i+1] <= 3 && num-line[i+1] >= 0 {
					decrease = true
					//fmt.Println(line, "decreases")
				} else {
					// Unsafe detected, get out of this line
					//fmt.Println(line, "unsafe")
					break
				}
			}

			// Middle numbers, compare against others
			if i >= 1 && i < len(line)-1 {
				if increase {
					difference = line[i+1] - num
					fmt.Println("num:", num, "num+1:", line[i+1], "difference:", difference)
				}

				if decrease {
					difference = num - line[i+1]
					fmt.Println("num:", num, "num+1:", line[i+1], "difference:", difference)
				}

				if difference >= 4 || difference <= 0 {
					fmt.Println(num, line[i+1], difference, "detected unsafe. breaking")
					// Unsafe detected, get out of this line
					break
				}
			}

			// Last number
			if i == len(line)-1 {
				fmt.Println(line, "is safe")
				safeLines += 1
			}
		}
	}

	return safeLines
}

func GetInputAsArrayOfInt64Array(filename string) [][]int64 {
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
