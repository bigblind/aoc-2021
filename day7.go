package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() []int {
	b, _ := os.ReadFile("input7.txt")
	s := string(b)
	crabs := make([]int, 0)
	for _, c := range strings.Split(s, ",") {
		crab, _ := strconv.ParseInt(c, 10, 64)
		crabs = append(crabs, int(crab))
	}
	return crabs
}

func getFuelCost(crabs []int, pos int) int {
	res := 0
	for _, crab := range crabs {
		res += int(math.Abs(float64(crab - pos)))
	}
	return res
}

func getFuelCostPart2(crabs []int, pos int) int {
	res := 0
	for _, crab := range crabs {
		dist := math.Abs(float64(crab - pos))
		cost := (dist * (dist + 1)) / 2
		res += int(cost)
	}
	return res
}

func part1() int {
	crabs := sort.IntSlice(readInput())
	crabs.Sort()
	if len(crabs)%2 == 0 {
		pos := crabs[len(crabs)/2]
		return getFuelCost(crabs, pos)
	} else {
		pos := (crabs[len(crabs)/2] + crabs[len(crabs)/2+1]) / 2
		return getFuelCost(crabs, pos)
	}
}

func part2() int {
	crabs := readInput()
	sum := 0
	for _, crab := range crabs {
		sum += crab
	}
	avg := float64(sum) / float64(len(crabs))
	low := getFuelCostPart2(crabs, int(math.Floor(avg)))
	hi := getFuelCostPart2(crabs, int(math.Ceil(avg)))
	return int(math.Min(float64(low), float64(hi)))
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}
