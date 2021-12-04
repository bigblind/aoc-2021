def read_depths():
    with open("input1.txt") as inp:
        return [int(x) for x in inp.readlines()]

def count_increments(numbers):
    result = 0
    prev = numbers[0]
    for n in numbers[1:]:
        if n > prev:
            result += 1
        prev = n
    return result

def part_1():
    return count_increments(read_depths())

def part_2():
    depths = read_depths()
    trios = [sum(depths[i:i+3]) for i in range(len(depths) - 2)]
    return count_increments(trios)

print("part 1:", part_1())
print("part 2:", part_2())


