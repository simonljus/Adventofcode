package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseBingonumbers(s string) []int {
	var numbers []int
	for _, d := range strings.Split(s, ",") {
		numbers = append(numbers, parseInteger(d))
	}
	return numbers
}
func parseInteger(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func bingoboard(rows []string) []map[int]bool {
	size := len(rows)
	lines := make([]map[int]bool, size*2)
	for i, r := range rows {
		for j, numstring := range strings.Fields(r) {
			number := parseInteger(numstring)
			if lines[i] == nil {
				lines[i] = make(map[int]bool, size)
			}
			lines[i][number] = true
			if lines[size+j] == nil {
				lines[size+j] = make(map[int]bool, size)
			}
			lines[size+j][parseInteger(numstring)] = true
		}

	}
	return lines
}
func getSum(board []map[int]bool) int {
	sum := 0
	unique := make(map[int]bool, len(board)*len(board))
	for _, row := range board {
		for n := range row {
			if _, ok := unique[n]; !ok {
				unique[n] = true
				sum += n
			}
		}
	}
	return sum
}
func playBingo(numbers []int, boards [][]map[int]bool) (int, int) {
	firstBingoIndex := -1
	lastBingoIndex := -1
	var p1 = -1
	var p2 = -1
	hasBingo := map[int]int{}
	bingoCount := 0
	for _, n := range numbers {
		if bingoCount == len(boards) {
			break
		}
		for boardIndex, board := range boards {
			if _, exists := hasBingo[boardIndex]; exists {
				continue
			}
			for _, row := range board {
				delete(row, n)
				if len(row) == 0 {
					if _, exists := hasBingo[boardIndex]; exists {
						continue
					}
					bingoCount += 1
					if bingoCount == len(boards) {
						lastBingoIndex = boardIndex
					}
					hasBingo[boardIndex] = len(hasBingo) + 1
					if bingoCount == 1 {
						firstBingoIndex = boardIndex
					}
				}
			}
		}
		if bingoCount == 1 && p1 == -1 {
			p1 = getSum(boards[firstBingoIndex]) * n
		}
		if bingoCount == len(boards) && p2 == -1 {
			p2 = getSum(boards[lastBingoIndex]) * n

		}
	}
	return p1, p2
}
func readFile(filename string) ([]int, [][]map[int]bool) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := []string{}
	scanner.Scan()
	numbers := parseBingonumbers(scanner.Text())
	var boards [][]map[int]bool
	scanner.Scan()
	for scanner.Scan() {
		row := scanner.Text()
		if len(row) == 0 {
			board := bingoboard(instructions)

			boards = append(boards, board)
			instructions = []string{}
		} else {
			instructions = append(instructions, row)
		}

	}
	if len(instructions) > 0 {
		boards = append(boards, bingoboard(instructions))
	}
	return numbers, boards
}
func solve(filename string) (int, int) {
	numbers, boards := readFile(filename)
	p1, p2 := playBingo(numbers, boards)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 4512))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 1924))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 63424))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 23541))
}
