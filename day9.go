package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() [][]int {
	b, _ := os.ReadFile("input9.txt")
	lines := strings.Split(string(b), "\n")
	cols := len(lines[0])
	rows := len(lines)
	grid := make([][]int, rows)
	for i, l := range lines {
		row := make([]int, cols)
		for j, c := range l {
			value, _ := strconv.ParseInt(string(c), 10, 64)
			row[j] = int(value)
		}
		grid[i] = row
	}
	return grid
}

func getNeighbours(grid [][]int, row, col int) []int {
	res := make([]int, 0, 4)
	if row > 0 {
		res = append(res, grid[row-1][col])
	}
	if row < len(grid)-1 {
		res = append(res, grid[row+1][col])
	}
	if col > 0 {
		res = append(res, grid[row][col-1])
	}
	if col < len(grid[0])-1 {
		res = append(res, grid[row][col+1])
	}
	return res
}

func isLowPoint(value int, neighbours []int) bool {
	for _, v := range neighbours {
		if value >= v {
			return false
		}
	}
	return true
}

func part1() int {
	grid := readInput()
	res := 0
	for i, row := range grid {
		for j, value := range row {
			if isLowPoint(value, getNeighbours(grid, i, j)) {
				res += value + 1
			}
		}
	}
	return res
}

func getBasinSize(grid [][]int, row, col int) int {
	if grid[row][col] == 9 {
		return 0
	}

	res := 1
	value := grid[row][col]
	grid[row][col] = -1 // prevent it from being scanned again
	if row > 0 && grid[row-1][col] >= value {
		res += getBasinSize(grid, row-1, col)
	}
	if row < len(grid)-1 && grid[row+1][col] >= value {
		res += getBasinSize(grid, row+1, col)
	}
	if col > 0 && grid[row][col-1] >= value {
		res += getBasinSize(grid, row, col-1)
	}
	if col < len(grid[0])-1 && grid[row][col+1] >= value {
		res += getBasinSize(grid, row, col+1)
	}
	return res
}

func part2() int {
	grid := readInput()
	sizes := make([]int, 0)
	for i, row := range grid {
		for j, value := range row {
			if isLowPoint(value, getNeighbours(grid, i, j)) {
				sizes = append(sizes, getBasinSize(grid, i, j))
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}
