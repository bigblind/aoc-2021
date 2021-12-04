def read_course():
    with open("input2.txt") as inp:
        return [parse_instruction(l) for l in inp.readlines()]

def parse_instruction(instr):
    instr = instr.split(" ")
    return (instr[0], int(instr[1]))

def part_1():
    forward = 0
    depth = 0
    for direction, amount in read_course():
        if direction == "forward":
            forward += amount
        elif direction == "up":
            depth -= amount
        else:
            depth += amount
    
    return depth * forward
        

def part_2():
    forward = 0
    depth = 0
    aim = 0
    for cmd, amount in read_course():
        if cmd == "forward":
            forward += amount
            depth += aim * amount
        elif cmd == "up":
            aim -= amount
        else:
            aim += amount
    
    return depth * forward

print("part 1:", part_1())
print("part 2:", part_2())


