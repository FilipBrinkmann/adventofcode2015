"""adventofcode 5"""

import re


def hasMinVowels(amount, string):
    count = 0
    for i in xrange(0, len(string)):
        if string[i] in ['a','e','i','o','u']:
            count += 1
        if count >= amount:
            return True
    return False

def hasDoubleLetter(string):
    if len(string) == 0:
        return False
    lastletter = None
    for i in xrange(0, len(string)):
        if string[i] is lastletter:
            return True
        else:
            lastletter = string[i]
    return False

def containsBadSequence(string):
    for sequence in ['ab', 'cd', 'pq', 'xy']:
        if sequence in string:
            return True
    return False

def containsDoublePairs(line):
    if len(line) < 4:
        return False
    for i in xrange(0, len(line)-1):
        pattern = line[i:i+2]
        for j in xrange(i+2,len(line)-1):
            sequence = line[j:j+2]
            if len(sequence) is 2:
                # print "comparing " + pattern + " <-> " + sequence
                if pattern == sequence:
                    return True
    return False

def containsRepeatingLetter(line):
    if len(line) < 3:
        return False
    for i in xrange(0, len(line)-2):
        if line[i] == line [i+2]:
            return True
    return False

def isNice(line):
    # return (hasMinVowels(3, line) and hasDoubleLetter(line) and not containsBadSequence(line))
    return containsDoublePairs(line) and containsRepeatingLetter(line)

# Tests for containsDoublePairs()
assert not containsDoublePairs('')
assert not containsDoublePairs('aaa')
assert not containsDoublePairs('abc')
assert not containsDoublePairs('abcde')
assert containsDoublePairs('abfeab')
assert containsDoublePairs('abab')

# Tests for containsRepeatingLetter
assert containsRepeatingLetter('xyx')
assert containsRepeatingLetter('abcdefeghi')
assert containsRepeatingLetter('aaa')
assert not containsRepeatingLetter('abc')
assert not containsRepeatingLetter('')
assert not containsRepeatingLetter('a')
assert not containsRepeatingLetter('aa')
# exit()

with open('input5.txt') as infile:
    nicewords = 0
    line = infile.readline()
    while(line):
        if(isNice(line)):
            nicewords += 1
        line = infile.readline()
    print "nice words: " + str(nicewords)
