import json

def read_diagnostics():
    with open("input3.txt") as inp:
        return [x.strip() for x in inp.readlines()]

def parse_instruction(instr):
    instr = instr.split(" ")
    return (instr[0], int(instr[1]))

def get_gamma_bits(diagnostics):
    nbits = len(diagnostics[0])
    ndiagnostics = len(diagnostics)
    gamma_binary = ""
    for i in range(nbits):
        ones = len([d for d in diagnostics if d[i] == "1"])
        if ones > (ndiagnostics / 2):
            gamma_binary += "1"
        else:
            gamma_binary += "0"

    return gamma_binary

def invert_bits(bits):
    return "".join([{"0": "1", "1": "0"}[b] for b in bits])

def part_1():
    gamma_binary = get_gamma_bits(read_diagnostics())
    epsilon_binary = invert_bits(gamma_binary)

    return int(gamma_binary, 2) * int(epsilon_binary, 2)

def part_2():
    diagnostics = read_diagnostics()
    nbits = len(diagnostics[0])
    potential_oxygen = diagnostics
    potential_co2 = list(diagnostics)
    for i in range(nbits):
        if len(potential_oxygen) > 1:
            ones = len([d for d in potential_oxygen if d[i] == "1"])
            most_popular = "1" if ones >= len(potential_oxygen) / 2 else "0"
            potential_oxygen = list([d for d in potential_oxygen if d[i] == most_popular])
        
        if len(potential_co2) > 1:
            ones = len([d for d in potential_co2 if d[i] == "1"])
            least_popular = "0" if ones >= len(potential_co2) / 2 else "1"
            potential_co2 = list([d for d in potential_co2 if d[i] == least_popular])

    ox = int(potential_oxygen[0], 2)
    co = int(potential_co2[0], 2)

    return int(potential_oxygen[0], 2) * int(potential_co2[0], 2)


print("part 1:", part_1())
print("part 2:", part_2())
