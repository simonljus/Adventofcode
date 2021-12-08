package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
func countFish(init []int, days int) int {
	cache := map[int]int{}
	count := len(init)
	for _, s0 := range init {
		count += countFishRecursive(days, s0, 0, cache)
	}
	return count
}

func countFishRecursive(maxDays, timer, timestamp int, cache map[int]int) int {
	days := timer + timestamp

	if _, exists := cache[days]; exists {
		return cache[days]
	}
	counter := 0
	for t := timer + timestamp + 1; t <= maxDays; t += 7 {
		counter += 1 + countFishRecursive(maxDays, 8, t, cache)
	}
	cache[days] = counter
	return counter
}
func problem1(instructions []int) int {
	return countFish(instructions, 80)
}
func problem2(instructions []int) int {
	return countFish(instructions, 256)
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
	p1 := problem1(instructions)
	p2 := problem2(instructions)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 5934))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 26984457539))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 360268))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 1632146183902))
}
