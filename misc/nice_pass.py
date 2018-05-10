#!/usr/bin/env python

# Generate a "comfortable" password. Alternate keypresses will be on
# alternate sides of the keyboard, and no key will require the shift
# key to be depressed.

import itertools
import random
import sys

L = ["q", "w", "e", "r", "t",
     "a", "s", "d", "f", "g",
     "z", "x", "c", "v", "b",
     "1", "2", "3", "4", "5"]

R = ["y", "u", "i", "o", "p",
     "h", "j", "k", "l", ";",
     "n", "m", ",", ".", "/",
     "6", "7", "8", "9", "0",
     "-", "=", "[", "]", "'"]

try:
    length = max(0, int(sys.argv[1]))
except (IndexError, ValueError):
    print "Usage: {} length".format(sys.argv[0])
    sys.exit(1)

sides = random.choice(([L,R], [R,L]))
chars = (random.choice(sides[i%2]) for i in xrange(length))
print "".join(chars)
