package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseTimestamp(s string) int {
	val, _ := strconv.ParseInt(s, 10, 64)
	return int(val)
}
func parseBuses(s string) []string {
	return strings.Split(s, ",")
}
func problem1(timestamp int, allbuses []string) int {
	minTime := math.MaxUint32
	selectedBus := -1
	buses := []int{}
	for _, val := range allbuses {
		id, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			buses = append(buses, int(id))
		}
	}
	for _, bus := range buses {
		waitTime := (bus - (timestamp % bus)) % bus
		if waitTime < minTime {
			minTime = waitTime
			selectedBus = bus
		}
	}
	return minTime * selectedBus
}
func problem2(instructions []string) int {
	mapped := make(map[int]int)
	ids := []int{}
	for i, instruction := range instructions {
		id, err := strconv.ParseInt(instruction, 10, 64)
		if err == nil {
			mapped[int(id)] = i
			ids = append(ids, int(id))
		}
	}
	sort.Ints(ids)
	factor := 1
	constant := 0
	for _, id := range ids {
		pos, _ := mapped[id]
		nFactor, nConstant := crt(factor, constant, id-pos, id)
		constant = factor*nConstant + constant
		factor = factor * nFactor

	}
	return constant
}

/*
https://play.golang.org/p/SmzvkDjYlb
*/
func getGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func crt(factor int, constant int, remainder int, mod int) (int, int) {
	nRemainder := (remainder - constant) % mod
	if nRemainder < 0 {
		nRemainder += mod
	}
	if factor%mod != 0 {
		factor = factor % mod
	}
	i := 0
	for factor != getGCD(factor, i*mod+nRemainder) {
		i++
	}
	gcdRemainder := (i*mod + nRemainder) / factor
	return mod, gcdRemainder % mod

}
func readFile(filename string) (int, []string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestamp := parseTimestamp(scanner.Text())
	scanner.Scan()
	buses := parseBuses(scanner.Text())
	return timestamp, buses
}
func testproblem2(busstr string, expected int) {
	buses := parseBuses(busstr)
	ans := problem2(buses)
	if expected != ans {
		fmt.Println("Test failed,", busstr, "Should be", expected, "Actual", ans)
	}

}
func solve(filename string) (int, int) {
	timestamp, buses := readFile(filename)
	p1 := problem1(timestamp, buses)
	p2 := problem2(buses)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	p1, p2 := solve("input.txt")
	testproblem2("7,13,x,x,59,x,31,19", 1068781)
	testproblem2("17,x,13,19", 3417)
	testproblem2("67,7,59,61", 754018)
	testproblem2("67,x,7,59,61", 779210)
	testproblem2("67,7,x,59,61", 1261476)
	testproblem2("1789,37,47,1889", 1202161486)
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 295))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 1068781))
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 6568))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 554865447501099))

}
