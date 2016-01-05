"""Solver for advent of code 2015, day 8"""

DEBUG = True

def printd(message):
    if DEBUG:
        print message

def find_double_ticks(string):
    count = 0
    for i in range(len(string)-1):
        if i == 1 and string[i] == "\"":
            count += 1
        elif string[i] != '\\':
            if string[i+1] == r'"':
                count += 1
    if len(string)>0 and string[len(string)-1] == r'"':
        count += 1
    return count

def find_escaped_ticks(string):
    count = 0
    for i in range(len(string)-1):
        if string[i] == '\\':
            if string[i+1] == '"':
                count += 1
    return count

def find_double_backslashes(string):
    count = 0
    offset = 0
    for i in range(len(string)-1):
        if i+offset+1 > len(string):
            break
        if string[i+offset] == '\\':
            if string[i+offset+1] == '\\':
                count += 1
                offset += 1
    return count

def find_hex(string):
    count = 0
    for i in range(len(string)-3):
        if string[i] == '\\':
            if string[i+1] == 'x':
                count += 3
                i += 3
    return count

def calculate_code_length(string, is_code):
    printd("----------------")
    printd("String: " + string.encode())
    code_characters = 0
    code_characters += find_double_ticks(string)
    printd("code characters after double ticks: " + str(code_characters))
    code_characters += find_escaped_ticks(string)
    printd("code characters after escaped ticks: " + str(code_characters))
    code_characters += find_double_backslashes(string)
    printd("code characters after double backslashes: " + str(code_characters))
    code_characters += find_hex(string)
    # print re.findall(r"(\\x[0-9A-F][0-9A-F])+", string)
    printd("code characters after hex values: " + str(code_characters))
    printd("length of string:" + str(len(string)))
    printd("length without code characters: " + str(len(string) - code_characters))
    printd("----------------")

    # if len(matches) > 0:
    #     print len(string)
    #     print len(matches)
    if is_code:
        return len(string)
    elif len(string) >= code_characters and not is_code:
        return len(string)-code_characters
    else:
        raise ValueError("more code characters than in memory length")


def main():
    # testing helper methods
    assert calculate_code_length(r'', True) == 0
    assert calculate_code_length(r'', False) == 0
    assert calculate_code_length(r'""', False) == 0
    assert calculate_code_length(r'""', True) == 2
    assert calculate_code_length(r'"abc"', False) == 3
    assert calculate_code_length(r'"abc"', True) == 5
    assert calculate_code_length(r'"aaa\"aaa"', True) == 10
    assert calculate_code_length(r'"aaa\"aaa"', False) == 7
    assert calculate_code_length(r'"\x27"', True) == 6
    assert calculate_code_length(r'"\x27"', False) == 1

    with open('input8.txt', 'r') as infile:
        memory = 0
        code = 0
        for line in infile:
            line = line.rstrip()
            code += calculate_code_length(line, True)
            memory += calculate_code_length(line, False)
        print str(code) + " - " + str(memory) + " = " + str(code - memory)

if __name__ == '__main__':
    main()