package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseString(s string) string {
	return s
}
func problem1(instructions []string) int {
	return 0
}
func problem2(instructions []string) int {
	return 0
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
	p1 := problem1(instructions)
	p2 := problem2(instructions)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 42))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 1337))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 42))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expexted %d", p2, 1337))
}
