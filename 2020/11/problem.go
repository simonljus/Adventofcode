package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countInDirection(row int, col int, rowadder int, coladder int, rows [][]string) int {
	rowpos := row
	colpos := col
	nrows := len(rows)
	ncols := len(rows[0])
	for factor := 1; true; factor++ {
		rowpos += rowadder
		colpos += coladder
		if colpos < 0 || colpos >= ncols {
			return 0
		}
		if rowpos < 0 || rowpos >= nrows {
			return 0
		}
		seat := rows[rowpos][colpos]
		if seat == "#" {
			return 1
		}
		if seat == "L" {
			return 0
		}
	}
	return 0
}
func p2updater(row int, col int, rows [][]string) string {
	adjacent := []int{-1, 0, 1}
	occupied := 0
	seat := rows[row][col]
	if seat != "#" && seat != "L" {
		return seat
	}
	for _, rowadder := range adjacent {
		for _, coladder := range adjacent {
			if rowadder == 0 && coladder == 0 {
				continue
			}
			occupied += countInDirection(row, col, rowadder, coladder, rows)
			if seat == "#" && occupied >= 5 {
				return "L"
			}
		}
	}
	if seat == "L" && occupied == 0 {
		return "#"
	}
	return seat
}

func countTotalOccupied(rows [][]string) int {
	count := 0
	for _, row := range rows {
		for _, seat := range row {
			if seat == "#" {
				count += 1
			}
		}
	}
	return count
}

func countOccupied(row int, col int, rows [][]string) int {
	adjacent := []int{-1, 0, 1}
	nrows := len(rows)
	ncols := len(rows[0])
	occupied := 0
	for _, rowadder := range adjacent {
		for _, coladder := range adjacent {
			rowpos := row + rowadder
			colpos := col + coladder
			if rowpos < 0 || colpos < 0 {
				continue
			}
			if colpos >= ncols || rowpos >= nrows {
				continue
			}
			if rowadder == 0 && coladder == 0 {
				continue
			}
			adjacentSeat := rows[rowpos][colpos]
			if adjacentSeat == "#" {
				occupied = occupied + 1
			}
		}
	}
	return occupied
}

func p1updater(row int, col int, rows [][]string) string {
	seat := rows[row][col]
	if seat != "L" && seat != "#" {
		return seat
	}
	count := countOccupied(row, col, rows)
	if seat == "L" && count == 0 {
		return "#"
	} else if seat == "#" && count >= 4 {
		return "L"
	}
	return seat
}
func iterateSeats(rows [][]string, seatupdater func(int, int, [][]string) string) int {
	changed := true
	for iterations := 0; changed == true; iterations++ {
		updated := [][]string{}
		changed = false
		for i, row := range rows {
			updatedrow := []string{}
			for j, seat := range row {
				updatedSeat := seatupdater(i, j, rows)
				if updatedSeat != seat {
					changed = true
				}
				updatedrow = append(updatedrow, updatedSeat)

			}
			updated = append(updated, updatedrow)
		}
		rows = updated
		if !changed {
			return countTotalOccupied(rows)
		}
	}
	return -1
}
func parseString(s string) []string {
	return strings.Split(s, "")
}

func readFile(filename string) [][]string {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	instructions := [][]string{}
	for scanner.Scan() {
		instruction := parseString(scanner.Text())
		instructions = append(instructions, instruction)
	}
	return instructions
}

func solve(filename string) (int, int) {
	instructions := readFile(filename)
	p1 := iterateSeats(instructions, p1updater)
	p2 := iterateSeats(instructions, p2updater)
	return p1, p2
}

func main() {
	tp1, tp2 := solve("test_input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 37))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 26))
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 2243))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expexted %d", p2, 2027))
}
