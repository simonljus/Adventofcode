package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func parseString(s string) int {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		return -1
	} else {
		return int(i)
	}
}
func addToMap(diff int, adapter int, m *map[int][]int) {
	worker := *m
	adapters, exists := worker[diff]
	if !exists {
		init := []int{}
		worker[diff] = append(init, adapter)
	} else {
		worker[diff] = append(adapters, adapter)
	}

}
func getAns(m *map[int][]int) int {
	worker := *m
	ones, _ := worker[1]
	threes, _ := worker[3]
	return len(ones) * len(threes)
}

func problem2(instructions []int) int {
	maxdiff := 3
	instructions = append(instructions, 0)
	sort.Ints(instructions)
	size := len(instructions)
	combinations := make([]int, size)
	combinations[0] = 1
	for i := 0; i < size; i++ {
		for j := i + 1; j < size && instructions[j]-instructions[i] <= maxdiff; j++ {
			combinations[j] += combinations[i]
		}
	}
	return combinations[size-1]
}

func problem1(instructions []int) int {
	m := make(map[int][]int)
	current := 0
	for i := 0; i < len(instructions); i++ {
		diff := instructions[i] - current
		current = instructions[i]
		addToMap(diff, current, &m)
	}
	myAdapter := current + 3
	addToMap(3, myAdapter, &m)
	return getAns(&m)
}

func readFile(filename string) []int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := []int{}
	for scanner.Scan() {
		instruction := parseString(scanner.Text())
		instructions = append(instructions, instruction)
	}
	sort.Ints(instructions)
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
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 35))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 8))
	t2p1, t2p2 := solve("test_input_2.txt")
	fmt.Println(fmt.Sprintf("Test file2 Problem 1: %d , expected %d", t2p1, 220))
	fmt.Println(fmt.Sprintf("Test file2 Problem 2: %d , expected %d", t2p2, 19208))

	p1, p2 := solve("input.txt")

	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 2484))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expexted %d", p2, 15790581481472))
}
