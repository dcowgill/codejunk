#!/usr/bin/env python

import sys

dialpad = [
    (4, 6),    # 0
    (6, 8),    # 1
    (7, 9),    # 2
    (4, 8),    # 3
    (3, 9, 0), # 4
    (),        # 5
    (1, 7, 0), # 6
    (2, 6),    # 7
    (1, 3),    # 8
    (2, 4),    # 9
    ]

def memoize(f):
    cache = {}
    def f2(*args):
        if args in cache:
            return cache[args]
        v = f(*args)
        cache[args] = v
        return v
    return f2

@memoize
def dial(n, digits):
    if n <= 1: return len(digits)
    return sum(dial(n-1, dialpad[d]) for d in digits)

N = int(sys.stdin.readline())
print dial(N, xrange(10))
