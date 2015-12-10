"""http://adventofcode.com/day/4"""

import hashlib

m = hashlib.md5()
key = "ckczppom"

goodhash = None
counter = 1

while not goodhash:
    m = hashlib.md5()
    value = key + str(counter)
    m.update(value)
    hash = m.hexdigest()
    # print value + ":" + hash
    if hash.startswith("000000"):
        goodhash = hash
        break
    else:
        counter+=1

print "value:" + value
print "hash:" + goodhash
print "counter: "+str(counter)
