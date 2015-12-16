"""Solver for advent of code no.7"""

WIRES = { 'a' : 'me'}

def value_of(wire):
    global WIRES
    return WIRES[wire]

def process_line():
    pass

def main():
    with open('input7.txt') as infile:
        line = infile.readline()
        while line:
            process_line(line)
    print value_of('a')

if __name__ == '__main__':
    main()
