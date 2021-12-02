package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseString(s string) string {
	return s
}
func parseInteger(s string) (int, error) {
	return strconv.Atoi(s)
}
func problem1(measurements []int) int {
	return countIncreased(measurements, 1)
}
func sum(measurements []int) int {
	sum := 0
	for _, el := range measurements {
		sum += el
	}
	return sum
}
func problem2(measurements []int) int {
	return countIncreased(measurements, 3)
}
func countIncreased(measurements []int, windowsize int) int {
	counter := 0
	for i := 1; i+windowsize <= len(measurements); i++ {
		if sum(measurements[i-1:i-1+windowsize]) < sum(measurements[i:i+windowsize]) {
			counter++
		}
	}
	return counter
}
func readFile(filename string) []int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	measurements := []int{}
	for scanner.Scan() {
		instruction, _ := parseInteger(scanner.Text())
		measurements = append(measurements, instruction)
	}
	return measurements
}
func solve(filename string) (int, int) {
	measurements := readFile(filename)
	p1 := problem1(measurements)
	p2 := problem2(measurements)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 7))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 5))

	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 1292))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 1262))
}
