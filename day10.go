package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func readInput() []string {
	b, _ := os.ReadFile("input10.txt")
	lines := strings.Split(string(b), "\n")
	return lines
}

type stack struct {
	items []string
}

func newStack() *stack {
	return &stack{
		items: make([]string, 0),
	}
}

func (s *stack) Push(str string) {
	s.items = append(s.items, str)
}

func (s *stack) Pop() string {
	if len(s.items) == 0 {
		return "empty"
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val
}

var braceMap = map[string]string{
	"(": ")",
	"<": ">",
	"{": "}",
	"[": "]",
}

var scoreMap = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var score2Map = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func solve() (int, int) {
	part1Score := 0
	lineScores := make([]int, 0)
	for _, line := range readInput() {
		s := newStack()
		corrupted := false
		for _, r := range line {
			if closer, ok := braceMap[string(r)]; ok {
				s.Push(closer)
			} else {
				expected := s.Pop()
				if string(r) != expected {
					part1Score += scoreMap[string(r)]
					corrupted = true
					break
				}
			}
		}
		lineScorePart2 := 0
		if !corrupted {
			for v := s.Pop(); v != "empty"; v = s.Pop() {
				lineScorePart2 = lineScorePart2*5 + score2Map[v]
			}
			lineScores = append(lineScores, lineScorePart2)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lineScores)))
	part2Score := lineScores[len(lineScores)/2]
	return part1Score, part2Score
}

func main() {
	p1, p2 := solve()
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}
