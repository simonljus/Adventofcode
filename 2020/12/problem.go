package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
Vector a vector
*/
type Vector struct {
	east  int
	north int
}

/*
* Ship a ship
 */
type Ship struct {
	east      int
	north     int
	direction rune
}

func toVec(direction rune) Vector {
	north := 0
	east := 0
	switch direction {
	case 'N':
		north = 1
	case 'E':
		east = 1
	case 'S':
		north = -1
	case 'W':
		east = -1
	}
	return Vector{east: east, north: north}

}

func executeCommand(command string, count int, ship *Ship) {
	switch command {
	case "F":
		vec := toVec(ship.direction)
		ship.east += count * vec.east
		ship.north += count * vec.north
	case "R":
		coords := "NESW"
		ind := strings.IndexRune(coords, ship.direction)
		addIndex := ((count + 360) % 360) / 90
		ind += addIndex
		ship.direction = rune(coords[ind%len(coords)])
	case "N":
		ship.north += count
	case "E":
		ship.east += count
	default:
		fmt.Println("UNKNOWN", command)
	}
}
func executeWaypoint(command string, count int, ship *Ship, wp *Vector) {
	switch command {
	case "F":
		ship.east += count * wp.east
		ship.north += count * wp.north
	case "R":
		addIndex := (((count + 360) % 360) / 90)
		for i := 0; i < addIndex; i++ {
			eastFactor := 1
			northFactor := -1
			if wp.north == 0 && wp.east == 0 {
				continue
			}
			oldEast := wp.east
			oldNorth := wp.north
			wp.east = oldNorth * eastFactor
			wp.north = oldEast * northFactor
		}
	case "N":
		wp.north += count
	case "E":
		wp.east += count
	default:
		fmt.Println("UNKNOWN", command)
	}
}

func parseString(s string) (string, int) {
	count, _ := strconv.Atoi(s[1:])
	command := string(s[0])
	directions := []string{"L", "S", "W"}
	opposites := []string{"R", "N", "E"}
	for i, d := range directions {
		if command == d {
			return opposites[i], -1 * count
		}
	}
	return command, count
}
func problem1(filename string) int {
	ship := readFile(filename)
	return int(math.Abs(float64(ship.north))) + int(math.Abs(float64(ship.east)))
}
func problem2(filename string) int {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	ship := Ship{direction: 'E'}
	waypoint := Vector{east: 10, north: 1}
	for scanner.Scan() {
		command, count := parseString(scanner.Text())
		executeWaypoint(command, count, &ship, &waypoint)
	}
	return int(math.Abs(float64(ship.north))) + int(math.Abs(float64(ship.east)))
}
func readFile(filename string) Ship {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	ship := Ship{direction: 'E'}
	for scanner.Scan() {
		command, count := parseString(scanner.Text())
		executeCommand(command, count, &ship)
	}
	return ship
}
func solve(filename string) (int, int) {

	p1 := problem1(filename)
	p2 := problem2(filename)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")

	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 25))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 286))
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 381))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 28591))
}
