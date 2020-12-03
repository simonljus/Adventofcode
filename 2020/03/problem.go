package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	col int
	row int
}

var p1Instructions = []Coord{Coord{3, 1}}
var p2Instructions = []Coord{Coord{1, 1}, Coord{3, 1}, Coord{5, 1}, Coord{7, 1}, Coord{1, 2}}

const tree = byte('#')

func move(matrix []string, pos *Coord, adder Coord) bool {
	pos.row = pos.row + adder.row
	pos.col = pos.col + adder.col
	nRows := len(matrix)
	nCols := len(matrix[0])
	if pos.row >= nRows {
		return false
	}
	return tree == matrix[pos.row][pos.col%nCols]
}
func traverse(matrix []string, adder Coord) int {
	pos := Coord{0, 0}
	counter := 0
	for pos.row < len(matrix) {
		if move(matrix, &pos, adder) {
			counter = counter + 1
		}
	}
	return counter
}

func readFile(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	var matrix []string
	for scanner.Scan() {
		matrix = append(matrix, scanner.Text())
	}
	return matrix
}
func problem1(filename string) int {
	return solve(filename, p1Instructions)
}
func problem2(filename string) int {
	return solve(filename, p2Instructions)
}
func solve(filename string, instructions []Coord) int {
	factor := 1
	matrix := readFile(filename)
	for _, instruction := range instructions {
		factor = factor * traverse(matrix, instruction)
	}
	return factor
}
func test(filename string, problemNr int, expected int, problemFunc func(s string) int) {
	if actual := problemFunc("test_input.txt"); actual != expected {
		fmt.Println(fmt.Sprintf("Test Problem %d failed: expected %d, should be %d", problemNr, expected, actual))
	}
}
func main() {
	test("test_input.txt", 1, 7, problem1)
	test("test_input.txt", 2, 336, problem2)
	fmt.Println("Problem1: ", problem1("input.txt"))
	fmt.Println("Problem2: ", problem2("input.txt"))
}
