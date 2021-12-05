from typing import NamedTuple


def read_input():
    with open("input5.txt") as f:
        lines = f.readlines()
        lines = [line.split(" -> ") for line in lines]
        lines = [(parse_point(line[0]), parse_point(line[1])) for line in lines]
        return lines


def parse_point(str):
    return [int(n) for n in str.split(",")]

def cmp(a, b):
    return (a > b) - (a < b) 

def solve(lines, diagonals):
    seen = set()
    seentwice = set()
    for line in lines:
        x1, y1 = line[0]
        x2, y2 = line[1]
        xinc = cmp(x2, x1)
        yinc = cmp(y2, y1)
        if diagonals == False and not (x1 == x2 or y1 == y2):
            continue

        pt = (x1, y1)
        while True:
            if pt in seen:
                seentwice.add(pt)
            seen.add(pt)
            if pt[0] == x2 and pt[1] == y2:
                break
            pt = (pt[0] + xinc, pt[1] + yinc)
    
    return len(seentwice)

lines =read_input()
print("part 1", solve(lines, False))
print("part 2", solve(lines, True))


