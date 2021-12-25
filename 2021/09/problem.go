package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Coord struct {
	x, y int
}

func getRiskLevel(x, y, xMax, yMax int, grid [][]int) int {
	v := getValue(x, y, xMax, yMax, grid)
	directions := [4](Coord){Coord{x: 0, y: -1}, Coord{x: 0, y: 1}, Coord{x: -1, y: 0}, Coord{x: 1, y: 0}}
	for _, dir := range directions {
		if getValue(x+dir.x, y+dir.y, xMax, yMax, grid) <= v {
			return 0
		}
	}
	return v + 1
}
func getValue(x, y, xMax, yMax int, grid [][]int) int {
	if x < 0 || y < 0 || x > xMax || y > yMax {
		return math.MaxInt
	}
	return grid[y][x]
}
func parseString(s string) []int {
	ints := make([]int, len(s))
	for i, r := range s {
		val, _ := strconv.Atoi(string(r))
		ints[i] = val
	}
	return ints
}
func getBasinSize(x, y, xMax, yMax int, grid [][]int, m map[int]bool) {
	v := getValue(x, y, xMax, yMax, grid)
	if v >= 9 {
		return
	}
	index := (xMax+1)*y + x
	if _, exists := m[index]; exists {
		return
	}
	m[index] = true

	directions := [4](Coord){Coord{x: 0, y: -1}, Coord{x: 0, y: 1}, Coord{x: -1, y: 0}, Coord{x: 1, y: 0}}
	for _, dir := range directions {
		getBasinSize(x+dir.x, y+dir.y, xMax, yMax, grid, m)
	}
	return
}
func solveProblems(instructions [][]int) (int, int) {
	sum := 0
	basins := []int{}
	yMax := len(instructions) - 1
	if yMax < 0 {
		return -1, -1
	}
	xMax := len(instructions[0]) - 1
	for y, row := range instructions {
		for x := range row {
			part := getRiskLevel(x, y, xMax, yMax, instructions)
			sum += part
			if part > 0 {
				m := map[int]bool{}
				getBasinSize(x, y, xMax, yMax, instructions, m)
				basins = append(basins, len(m))

			}
		}
	}
	sort.Ints(basins)
	fmt.Println()
	sum2 := 1
	for _, size := range basins[len(basins)-3:] {
		sum2 *= size
	}
	return sum, sum2
}
func problem2(instructions [][]int) int {
	return 0
}
func readFile(filename string) [][]int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := [][]int{}
	for scanner.Scan() {
		instruction := parseString(scanner.Text())
		instructions = append(instructions, instruction)
	}
	return instructions
}
func solve(filename string) (int, int) {
	instructions := readFile(filename)
	p1, p2 := solveProblems(instructions)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 15))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 1134))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 566))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 891684))
}
