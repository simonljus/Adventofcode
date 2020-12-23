package main

import (
	"bufio"
	"fmt"
	"os"
)

func countActive(layers [][][]byte, z, y, x int) int {
	factors := []int{-1, 0, 1}
	minZ := 0
	maxZ := len(layers) - 1
	minY := 0
	maxY := len(layers[0]) - 1
	minX := 0
	maxX := len(layers[0][0]) - 1
	count := 0
	maxCount := 3
	for _, dz := range factors {
		nz := z + dz
		if nz > maxZ || nz < minZ {
			continue
		}
		for _, dy := range factors {
			ny := y + dy
			if ny > maxY || ny < minY {
				continue
			}
			for _, dx := range factors {
				nx := x + dx
				if nx > maxX || nx < minX {
					continue
				}
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				if layers[nz][ny][nx] == '#' {
					count++
					if count > maxCount {
						return count
					}
				}
			}
		}
	}
	return count
}
func countActiveW(fourD [][][][]byte, w, z, y, x int) int {
	factors := []int{-1, 0, 1}
	minW := 0
	maxW := len(fourD) - 1
	minZ := 0
	maxZ := len(fourD[0]) - 1
	minY := 0
	maxY := len(fourD[0][0]) - 1
	minX := 0
	maxX := len(fourD[0][0][0]) - 1
	count := 0
	maxCount := 3
	for _, dw := range factors {
		nw := w + dw
		if nw > maxW || nw < minW {
			continue
		}
		for _, dz := range factors {
			nz := z + dz
			if nz > maxZ || nz < minZ {
				continue
			}
			for _, dy := range factors {
				ny := y + dy
				if ny > maxY || ny < minY {
					continue
				}
				for _, dx := range factors {
					nx := x + dx
					if nx > maxX || nx < minX {
						continue
					}
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					if fourD[nw][nz][ny][nx] == '#' {
						count++
						if count > maxCount {
							return count
						}
					}
				}
			}
		}
	}
	return count
}
func parseString(s string) []byte {
	return []byte(s)
}
func countCubes(layers [][][]byte) int {
	count := 0
	for _, rows := range layers {
		for _, row := range rows {
			for _, cube := range row {
				if cube == '#' {
					count++
				}
			}
		}
	}
	return count
}
func createEmptyRow(nCols int) []byte {
	row := []byte{}
	for j := 0; j < nCols; j++ {
		row = append(row, '.')
	}
	return row
}
func createEmptyMatrix(nRows, nCols int) [][]byte {
	matrix := [][]byte{}
	emptyRow := createEmptyRow(nCols)
	for i := 0; i < nRows; i++ {
		matrix = append(matrix, emptyRow)
	}
	return matrix
}
func printlayers(i int, layers [][][]byte) {
	fmt.Println("ITERATION", i)
	for z, rows := range layers {
		fmt.Println("LAYERS", z)
		for _, row := range rows {
			fmt.Println(string(row))
		}
	}
}
func expandRow(row []byte) []byte {
	before := []byte{'.'}
	after := []byte{'.'}
	before = append(before, row...)
	before = append(before, after...)
	return before
}
func expandLayer(layer [][]byte) [][]byte {
	updated := [][]byte{}
	row := []byte{}
	for i, r := range layer {
		layer[i] = expandRow(r)
	}
	for range layer[0] {
		row = append(row, '.')
	}
	updated = append(updated, row)
	updated = append(updated, layer...)
	updated = append(updated, row)
	return updated
}
func expandLayers(layers [][][]byte) [][][]byte {
	updatedLayers := [][][]byte{}
	for i, layer := range layers {
		layers[i] = expandLayer(layer)
	}
	emptyMatrix := createEmptyMatrix(len(layers[0]), len(layers[0][0]))
	updatedLayers = append(updatedLayers, emptyMatrix)
	updatedLayers = append(updatedLayers, layers...)
	updatedLayers = append(updatedLayers, emptyMatrix)
	return updatedLayers
}
func createEmptyW(z, y, x int) [][][]byte {
	layers := [][][]byte{}
	emptyMatrix := createEmptyMatrix(y, x)
	for i := 0; i < z; i++ {
		layers = append(layers, emptyMatrix)
	}
	return layers
}
func expandW(w [][][][]byte) [][][][]byte {
	updatedW := [][][][]byte{}
	for i, layers := range w {
		w[i] = expandLayers(layers)
	}
	emptyLayers := createEmptyW(len(w[0]), len(w[0][0]), len(w[0][0][0]))
	updatedW = append(updatedW, emptyLayers)
	updatedW = append(updatedW, w...)
	updatedW = append(updatedW, emptyLayers)
	return updatedW
}
func problem1(matrix [][]byte) int {
	layers := [][][]byte{}
	layers = append(layers, matrix)
	results := []int{}
	for i := 0; i < 6; i++ {
		results = append(results, countCubes(layers))
		layers = expandLayers(layers)
		updatedLayers := [][][]byte{}
		for z, rows := range layers {
			updatedLayer := [][]byte{}
			for y, row := range rows {
				updatedRow := []byte{}
				for x, cube := range row {
					count := countActive(layers, z, y, x)
					updatedCube := getCubeState(cube, count)
					updatedRow = append(updatedRow, updatedCube)
				}
				updatedLayer = append(updatedLayer, updatedRow)
			}
			updatedLayers = append(updatedLayers, updatedLayer)
		}
		layers = updatedLayers

	}
	ans := countCubes(layers)
	results = append(results, ans)
	fmt.Println("problem 1 cycles", results)
	return ans
}
func countCubesInW(w [][][][]byte) int {
	sum := 0
	for _, layers := range w {
		sum += countCubes(layers)
	}
	return sum
}
func getCubeState(cube byte, count int) byte {
	if cube == '.' && count == 3 {
		return '#'
	}
	if cube == '#' && !(count == 3 || count == 2) {
		return '.'
	}
	return cube
}
func problem2(matrix [][]byte) int {
	initlayers := [][][]byte{}
	initlayers = append(initlayers, matrix)
	fourD := [][][][]byte{}
	fourD = append(fourD, initlayers)
	results := []int{}
	for i := 0; i < 6; i++ {
		results = append(results, countCubesInW(fourD))
		updatedW := [][][][]byte{}
		fourD = expandW(fourD)
		for w, layers := range fourD {
			updatedLayers := [][][]byte{}
			for z, rows := range layers {
				updatedLayer := [][]byte{}
				for y, row := range rows {
					updatedRow := []byte{}
					for x, cube := range row {
						count := countActiveW(fourD, w, z, y, x)
						updatedCube := getCubeState(cube, count)
						updatedRow = append(updatedRow, updatedCube)
					}
					updatedLayer = append(updatedLayer, updatedRow)
				}
				updatedLayers = append(updatedLayers, updatedLayer)
			}
			updatedW = append(updatedW, updatedLayers)
		}
		fourD = updatedW
	}
	ans := countCubesInW(fourD)
	results = append(results, ans)
	fmt.Println("problem 2 cycles", results)
	return ans
}
func readFile(filename string) [][]byte {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	rows := [][]byte{}
	for scanner.Scan() {
		row := parseString(scanner.Text())
		rows = append(rows, row)
	}
	return rows
}
func solve(filename string) (int, int) {
	rows := readFile(filename)
	p1 := problem1(rows)
	p2 := problem2(rows)
	return p1, p2
}
func main() {
	tp1, tp2 := solve("test_input.txt")
	fmt.Println(fmt.Sprintf("Test Problem 1: %d , expected %d", tp1, 112))
	fmt.Println(fmt.Sprintf("Test Problem 2: %d , expected %d", tp2, 848))
	p1, p2 := solve("input.txt")
	fmt.Println(fmt.Sprintf("Problem 1: %d , expected %d", p1, 271))
	fmt.Println(fmt.Sprintf("Problem 2: %d , expected %d", p2, 2064))

}
