"""http://adventofcode.com/day/3"""

def is_santa(position):
    return (position % 2 is 0)


def is_robot(position):
    return not is_santa(position)

with open('input3.txt') as inputfile:
    x_santa = 0
    y_santa = 0
    x_robot = 0
    y_robot = 0
    visited_houses = []
    housecounter = 0

    line = inputfile.readline()
    for i in xrange(0, len(line)):
        if(is_santa(i)):
            x = x_santa
            y = y_santa
        elif(is_robot(i)):
            x = x_robot
            y = y_robot
        if line[i] is '^':
            y += 1
        elif line[i] is '>':
            x += 1
        elif line[i] is '<':
            x -= 1
        elif line[i] is 'v':
            y -= 1
        else:
            print "Invalid input: "+line[i]
            exit()
        if is_santa(i):
            x_santa = x
            y_santa = y
            if not (x, y) in visited_houses:
                housecounter += 1
                visited_houses.append((x, y))
        elif is_robot(i):
            x_robot = x
            y_robot = y
            if not (x, y) in visited_houses:
                housecounter += 1
                visited_houses.append((x, y))
print "Visited houses:" + str(housecounter)
