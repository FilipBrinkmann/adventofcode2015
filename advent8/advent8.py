"""Solver for advent of code 2015, day 8"""

def calculate_code_length(string):
    return len(string)


def main():
    # testing helper methods
    assert calculate_code_length('') == 0
    assert calculate_code_length('"') == 1
    assert calculate_code_length('""') == 2
    assert calculate_code_length('"abc"') == 5
    assert calculate_code_length('"aaa\"aaa"') == 10
    with open('input8.txt') as infile:
        line = infile.readline().strip("\r\n")
        while line:
            line = infile.readline()
            print line + str(calculate_code_length(line))

if __name__ == '__main__':
    main()
