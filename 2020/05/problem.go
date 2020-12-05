package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

const nCols = 8
const nRows = 128
const maxId = nCols*nRows - 1

func getSeatId(pos string) (int, error) {
	rowmin := 0
	rowmax := nRows - 1
	colmin := 0
	colmax := nCols - 1
	for _, c := range pos {
		if c == 'F' || c == 'B' {
			rowmin, rowmax = move(c, rowmin, rowmax)
		} else {
			colmin, colmax = move(c, colmin, colmax)
		}
	}
	if colmin != colmax {
		return -1, errors.New(fmt.Sprintf("column min  %d is not the same as column max %d", colmin, colmax))
	} else if rowmin != rowmax {
		return -1, errors.New(fmt.Sprintf("row min  %d is not the same as row max %d", rowmin, rowmax))
	}
	return rowmin*nCols + colmin, nil

}
func move(pos rune, min, max int) (int, int) {
	if pos == 'F' || pos == 'L' {
		max = max - (max-min+1)/2
	} else {
		min = min + (max-min+1)/2
	}
	return min, max
}
func testFunction(value string, expected int, f func(s string) (int, error)) {
	if actual, err := f(value); expected != actual {
		if err != nil {
			fmt.Println(fmt.Sprintf("Test failed: invalid data expected %d got error %s for value %s ", expected, err.Error(), value))
		} else {
			fmt.Println(fmt.Sprintf("Test failed: expected %d, was %d for value %s", expected, actual, value))
		}

	}
}
func test() {
	testFunction("BFFFBBFRRR", 567, getSeatId)
	testFunction("FFFBBBFRRR", 119, getSeatId)
	testFunction("BBFFBBFRLL", 820, getSeatId)

}
func solve(filename string) (int, int) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	seats := []int{}
	for scanner.Scan() {
		if seatId, err := getSeatId(scanner.Text()); err != nil {
			fmt.Println(err.Error())
		} else {
			seats = append(seats, seatId)
		}
	}
	sort.Ints(seats)
	maxId := seats[len(seats)-1]
	for i, seatId := range seats {
		if seats[i+1]-seatId == 2 {
			return maxId, seatId + 1
		}
	}
	return maxId, -1
}

func main() {
	test()
	p1, p2 := solve("input.txt")
	// 915 699
	fmt.Println("Problem1:", p1, "Problem2:", p2)
}
