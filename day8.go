package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type displays struct {
	digits []string
	output []string
}

func readInput() []displays {
	b, _ := os.ReadFile("input8.txt")
	s := string(b)
	lines := strings.Split(s, "\n")
	result := make([]displays, len(lines))
	for i, l := range lines {
		split := strings.Split(l, " | ")
		result[i] = displays{
			digits: strings.Split(split[0], " "),
			output: strings.Split(split[1], " "),
		}
	}
	return result
}

func part1() int {
	inp := readInput()
	res := 0
	for _, d := range inp {
		for _, dig := range d.output {
			l := len(dig)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				res += 1
			}
		}
	}

	return res
}

type condition func(digit string) bool

func length(l int) condition {
	return func(digit string) bool {
		return len(digit) == l
	}
}

func contains(seg rune) condition {
	return func(digit string) bool {
		return strings.ContainsRune(digit, seg)
	}
}

func containsAll(segs string) condition {
	return func(digit string) bool {
		for _, seg := range segs {
			if !strings.ContainsRune(digit, seg) {
				return false
			}
		}
		return true
	}
}

func not(cond condition) condition {
	return func(digit string) bool {
		return !cond(digit)
	}
}

func findDigits(digits []string, conds ...condition) []string {
	res := make([]string, 0)
outerLoop:
	for _, d := range digits {
		for _, cond := range conds {
			if !cond(d) {
				continue outerLoop
			}
		}
		res = append(res, d)
	}
	return res
}

func missingSegments(digit string) string {
	res := ""
	for _, r := range "abcdefg" {
		if !strings.ContainsRune(digit, r) {
			res += string(r)
		}
	}
	return res
}

// Analyzes the digits to make a an array mapping from the digit's value to its representation in the output.
// so the string at [0] is the representation of 0
func makeMapping(digits []string) []string {
	mapping := make([]string, 10)

	oneSegments := findDigits(digits, length(2))[0]
	mapping[1] = oneSegments
	sevenSegments := findDigits(digits, length(3))[0]
	mapping[7] = sevenSegments
	fourSegments := findDigits(digits, length(4))[0]
	mapping[4] = fourSegments
	eightSegments := findDigits(digits, length(7))[0]
	mapping[8] = eightSegments

	bAndD := ""
	for _, s := range fourSegments {
		if !strings.ContainsRune(oneSegments, s) {
			bAndD += string(s)
		}
	}

	nineAndZero := findDigits(digits, length(6), containsAll(oneSegments))
	nineSegments := findDigits(nineAndZero, containsAll(bAndD))[0]
	mapping[9] = nineSegments
	eSegment := rune(missingSegments(nineSegments)[0])
	zoroSegments := findDigits(nineAndZero, contains(eSegment))[0]
	mapping[0] = zoroSegments

	threeSegments := findDigits(digits, length(5), containsAll(oneSegments))[0]
	mapping[3] = threeSegments

	sixSegments := findDigits(digits, length(6), not(containsAll(oneSegments)))[0]
	mapping[6] = sixSegments

	twoAndFive := findDigits(digits, length(5), not(containsAll(threeSegments)))
	twoSegments := findDigits(twoAndFive, contains(eSegment))[0]
	fiveSegments := findDigits(twoAndFive, not(contains(eSegment)))[0]
	mapping[2] = twoSegments
	mapping[5] = fiveSegments

	return mapping
}

type bytes []byte

func (b bytes) Len() int           { return len(b) }
func (b bytes) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b bytes) Less(i, j int) bool { return b[i] < b[j] }

// Sorts all the segments within the digits
func standardizeDigits(digits []string) []string {
	result := make([]string, len(digits))
	for i, d := range digits {
		b := bytes(d)
		sort.Sort(b)
		result[i] = string(b)
	}
	return result
}

func parseOutput(out []string, mapping []string) int {
	res := 0
	for i, d := range out {
		var idx int
		found := false
		for j, d2 := range mapping {
			if d == d2 {
				idx = j
				found = true
				break
			}
		}
		if !found {
			panic(d + "-" + strings.Join(mapping, " "))
		}
		res += int(math.Pow10(4-(i+1))) * idx
	}
	return res
}

func part2() int {
	displays := readInput()
	sum := 0
	for _, d := range displays {
		out := standardizeDigits(d.output)
		mapping := standardizeDigits(makeMapping(d.digits))
		sum += parseOutput(out, mapping)
	}
	return sum
}

func main() {
	fmt.Printf("Part 1: %v\n", part1())
	fmt.Printf("Part 2: %v\n", part2())
}
