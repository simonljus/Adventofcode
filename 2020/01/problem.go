package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func problem1() int32 {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	uniques := make(map[int32]int)
	const sum = 2020
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		val32 := int32(val)
		_, exists := uniques[sum-val32]
		if exists {
			return (sum - val32) * val32
		}
		uniques[int32(val)] = 0
	}
	return -1
}
func findsum(m map[int32]int, sum int32) int32 {
	for k, _ := range m {
		_, exists := m[sum-k]
		if exists {
			return (sum - k) * k
		}
	}
	return 0
}
func problem2() int32 {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	uniques := make(map[int32]int)
	const sum = 2020
	for scanner.Scan() {
		val, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		val32 := int32(val)
		product := findsum(uniques, sum-val32)
		if product != 0 {
			return product * val32
		}
		uniques[int32(val)] = 0
	}
	return 0
}
func main() {
	fmt.Printf("Problem 1: %d", problem1())
	fmt.Printf("Problem 2: %d", problem2())
}
