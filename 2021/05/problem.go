package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseString(s string) string {
	return s
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func distance(a, b int) int {
	return max(a, b) - min(a, b)
}

func drawHorizontal(x1, y1, x2, y2 int, rows map[int]map[int]int) {
	if y1 != y2 {
		return
	}
	drawPoints(x1, y1, x2, y2, rows)
}
func drawVertical(x1, y1, x2, y2 int, rows map[int]map[int]int) {
	if x1 != x2 {
		return
	}
	drawPoints(x1, y1, x2, y2, rows)
}
func drawDiagonal(x1, y1, x2, y2 int, rows map[int]map[int]int) {
	if !(distance(x1, x2) == distance(y1, y2)) {
		return
	}
	drawPoints(x1, y1, x2, y2, rows)

}
func drawPoints(x1, y1, x2, y2 int, rows map[int]map[int]int) {
	dx := 0
	if x1 > x2 {
		dx = -1
	} else if x1 < x2 {
		dx = 1
	}
	dy := 0
	if y1 > y2 {
		dy = -1
	} else if y1 < y2 {
		dy = 1
	}
	for x1 != x2 || y1 != y2 {
		drawPoint(x1, y1, rows)
		x1 += dx
		y1 += dy
	}
	drawPoint(x1, y1, rows)
}
func drawPoint(x, y int, rows map[int]map[int]int) {
	if _, ok := rows[y]; !ok {
		rows[y] = map[int]int{}
	}
	row := rows[y]
	if _, ok := row[x]; !ok {
		row[x] = 0
	}
	row[x] += 1
}
func countScore(m map[int]map[int]int) int {
	count := 0
	for _, row := range m {
		for _, cell := range row {
			if cell >= 2 {
				count += 1
			}

		}
	}
	return count
}
func problem1(instructions []string, m map[int]map[int]int) {
	for _, instruction := range instructions {
		var x1, y1, x2, y2 int
		fmt.Sscanf(instruction, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		drawHorizontal(x1, y1, x2, y2, m)
		drawVertical(x1, y1, x2, y2, m)
	}
}
func problem2(instructions []string, m map[int]map[int]int) {
	for _, instruction := range instructions {
		var x1, y1, x2, y2 int
		fmt.Sscanf(instruction, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		drawDiagonal(x1, y1, x2, y2, m)
	}
}
func readFile(filename string) []string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := []string{}
	for scanner.Scan() {
		instruction := parseString(scanner.Text())
		instructions = append(instructions, instruction)
	}
	return instructions
}
func solve(filename string) (int, int) {
	instructions := readFile(filename)
	m := map[int]map[int]int{}
	problem1(instructions, m)
	p1 := countScore(m)
	problem2(instructions, m)
	p2 := countScore(m)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 5))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 12))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 6666))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 19081))
}
