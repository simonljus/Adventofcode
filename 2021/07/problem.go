package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	if a > b {
		return a - b
	}
	return b - a
}
func parseInteger(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func splitToInts(s string) []int {
	ints := []int{}
	for _, d := range strings.Split(s, ",") {
		ints = append(ints, parseInteger(d))
	}
	return ints
}
func count(elements []int) map[int]int {
	m := map[int]int{}
	for _, element := range elements {
		if _, exists := m[element]; !exists {
			m[element] = 0
		}
		m[element] += 1
	}
	return m
}

func arithmetic(n int) int {
	return (n * (n + 1)) / 2
}
func linear(n int) int {
	return n
}
func minimizeFuel(instructions []int, calculatefuel func(int) int) int {
	costs := map[int]int{}
	m := count(instructions)
	for to_pos := range m {
		cost := 0
		for from_pos, count := range m {
			distance := distance(from_pos, to_pos)
			cost += calculatefuel(distance) * count
		}
		costs[to_pos] = cost
	}
	mincost := -1
	for _, cost := range costs {
		if mincost == -1 {
			mincost = cost
		}
		if cost < mincost {
			mincost = cost
		}
	}
	return mincost
}
func readFile(filename string) []int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := splitToInts(scanner.Text())

	return instructions
}
func solve(filename string) (int, int) {
	instructions := readFile(filename)
	p1 := minimizeFuel(instructions, linear)
	p2 := minimizeFuel(instructions, arithmetic)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 37))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 168))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 357353))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 104822130))
}
