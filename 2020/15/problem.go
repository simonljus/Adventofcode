package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseString(s string) []int64 {
	ints := []int64{}
	for _, v := range strings.Split(s, ",") {
		i, _ := strconv.ParseInt(v, 10, 64)
		ints = append(ints, i)
	}
	return ints
}
func problem(numbers []int64, rounds int) int64 {
	mem := make(map[int64][]int64)
	lastSpoken := int64(0)
	firstTimeSpoken := true
	for i, n := range numbers {
		indexes := []int64{}
		indexes = append(indexes, int64(i+1))
		lastSpoken = n
		mem[n] = indexes
	}
	for i := len(mem) + 1; i <= rounds; i++ {
		if i%1000 == 0 {
			fmt.Println(i)
		}
		spoken := int64(-1)
		if firstTimeSpoken {
			spoken = int64(0)
		} else {
			turns, _ := mem[lastSpoken]
			spoken = turns[len(turns)-1] - turns[len(turns)-2]
		}
		lastSpoken = spoken
		turns, exists := mem[spoken]
		if !exists {
			firstTimeSpoken = true
			arr := []int64{int64(i)}
			mem[spoken] = arr
		} else {
			firstTimeSpoken = false
			mem[spoken] = append(turns, int64(i))
		}
	}
	return lastSpoken
}
func solve(input string) (int64, int64) {
	instructions := parseString(input)
	p1 := problem(instructions, 2020)
	p2 := problem(instructions, 30000000)
	return p1, p2
}

func main() {
	/*tp1, tp2 := solve("0,3,6")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 436))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 1337))*/
	p1, p2 := solve("13,0,10,12,1,5,8")

	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 260))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 950))
}
