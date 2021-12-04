def read_input():
    with open("input4.txt") as f:
        lines = f.readlines()
        numbers = [int(x) for x in lines[0].split(",")]
        boards = [lines[i:i+5] for i in range(2, len(lines), 6)]
        boards = [[[int(x) for x in row.split()] for row in board] for board in boards]
        return (numbers, boards) 


def has_bingo(numbers, board):
    for row in board:
        if all([x in numbers for x in row]):
            return True
    
    for i in range(5):
        if all([x in numbers for x in [row[i] for row in board]]):
            return True
    
    return False


def get_score(board, numbers):
    score = 0
    for row in board:
        for n in row:
            if n not in numbers:
                score += n
    return score


def part_1(numbers, boards):
    for i in range(5, len(numbers)):
        winners = [b for b in boards if has_bingo(numbers[:i], b)]
        if len(winners) > 0:
            score = max([get_score(b, numbers[:i]) for b in winners])
            return score * numbers[i-1]


def part_2(numbers, boards):
    for i in range(5, len(numbers)):
        if len(boards) > 1:
            boards = [b for b in boards if not has_bingo(numbers[:i], b)]
        elif has_bingo(numbers[:i], boards[0]):
            return get_score(boards[0], numbers[:i]) * numbers[i-1]

numbers, boards = read_input()
print("part 1:", part_1(numbers, boards))
print("part 2:", part_2(numbers, boards))