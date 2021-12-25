package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func removeCharacters(input string, characters string) string {
	filter := func(r rune) rune {
		if strings.IndexRune(characters, r) < 0 {
			return r
		}
		return -1
	}
	result := strings.Map(filter, input)
	return result

}
func parseInteger(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	sorted := strings.Join(s, "")
	return sorted
}
func sortArrayByLength(arr []string) []string {
	sort.Slice(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})
	return arr
}
func findSignals(numbers []string) map[string]int {
	m := map[string]int{}
	m2 := map[int]string{}
	maxLength := len(numbers)
	for len(m) != maxLength {
		for _, s := range numbers {
			if _, exists := m[s]; exists {
				continue
			}
			stringLength := len(s)
			if stringLength == 2 {
				//1
				m2[1] = s
				m[s] = 1
			} else if stringLength == 3 {
				//7
				m2[7] = s
				m[s] = 7
			} else if stringLength == 4 {
				//4
				m2[4] = s
				m[s] = 4
			} else if stringLength == 5 {
				//2,3,5
				if _, exists := m2[1]; !exists {
					continue
				}
				isADG := removeCharacters(s, m2[1])
				if len(isADG) == 3 {
					m2[3] = s
					m[s] = 3
				} else if s9, exists := m2[9]; exists {
					isEmpty := removeCharacters(s, s9)
					if len(isEmpty) == 0 {
						//5
						m2[5] = s
						m[s] = 5
					} else {
						m2[2] = s
						m[s] = 2
					}
				}
			} else if stringLength == 6 {
				//0,6,9
				if _, exists := m2[4]; !exists {
					continue
				}
				candidateAG := removeCharacters(s, m2[4])
				if len(candidateAG) == 2 {
					m2[9] = s
					m[s] = 9
				} else if s5, exists := m2[5]; exists {
					newCEorE := removeCharacters(s, s5)
					if len(newCEorE) == 1 {
						m2[6] = s
						m[s] = 6
					} else {
						m2[0] = s
						m[s] = 0
					}
				}
			} else if stringLength == 7 {
				//8
				m2[8] = s
				m[s] = 8
			}
		}
	}
	return m
}

func problem1(instructions []string) int {
	count := 0
	for _, line := range instructions {
		p := strings.Split(line, "|")[1]
		for _, signal := range strings.Split(p, " ") {
			signalLength := len(signal)
			if signalLength == 2 || signalLength == 3 || signalLength == 4 || signalLength == 7 {
				count += 1
			}
		}
	}
	return count
}

func problem2(instructions []string) int {
	count := 0
	for _, line := range instructions {
		numbers := strings.Split(line, "|")[0]
		sorted := []string{}
		for _, number := range strings.Split(numbers, " ") {
			if len(number) > 0 {
				sorted = append(sorted, sortString(number))
			}

		}
		m := findSignals(sortArrayByLength(sorted))
		p := strings.Split(line, "|")[1]
		displayed := []string{}
		for _, signal := range strings.Split(p, " ") {
			v := m[sortString(signal)]
			displayed = append(displayed, strconv.Itoa(v))
		}
		count += parseInteger(strings.Join(displayed, ""))

	}
	return count
}
func readFile(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := []string{}
	for scanner.Scan() {
		instruction := scanner.Text()
		instructions = append(instructions, instruction)
	}
	return instructions
}
func solve(filename string) (int, int) {
	instructions := readFile(filename)
	p1 := problem1(instructions)
	p2 := problem2(instructions)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 26))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 61229))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 237))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 1009098))
}
