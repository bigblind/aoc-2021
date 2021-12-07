package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []int {
	b, _ := os.ReadFile("input6.txt")
	s := string(b)
	fishes := make([]int, 0)
	for _, stringyfish := range strings.Split(s, ",") {
		fish, _ := strconv.ParseInt(stringyfish, 10, 64)
		fishes = append(fishes, int(fish))
	}
	return fishes
}

func simulate(days int) int {
	fishes := readInput()
	fishByDay := make([]int, 9)
	for _, f := range fishes {
		fishByDay[f] += 1
	}

	for d := 0; d < days; d++ {
		births := fishByDay[0]
		for i := 0; i < 8; i++ {
			fishByDay[i] = fishByDay[i+1]
		}
		fishByDay[8] = births
		fishByDay[6] += births
	}

	result := 0
	for _, fish := range fishByDay {
		result += fish
	}
	return result
}

func main() {
	fmt.Printf("Part 1: %v\n", simulate(80))
	fmt.Printf("Part 2: %v\n", simulate(256))
}
