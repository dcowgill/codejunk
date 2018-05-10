#!/usr/bin/env python3

import math, random

qt = [
    [0,3,1,7,5,9,8,6,4,2],
    [7,0,9,2,1,5,4,8,6,3],
    [4,2,0,6,8,7,1,3,5,9],
    [1,7,5,0,9,8,3,4,2,6],
    [6,1,2,3,0,4,5,9,7,8],
    [3,6,7,4,2,0,9,5,8,1],
    [5,8,6,9,7,2,0,1,3,4],
    [8,9,4,5,3,6,2,0,1,7],
    [9,4,3,8,6,1,7,2,0,5],
    [2,5,8,1,4,3,6,7,9,0],
]

def checksum(n):
    i = 0
    for d in list(map(int, str(n))):
        i = qt[i][d]
    return i

def is_valid(n):
    return checksum(n) == 0

def print_random_test_cases(n):
    print("// checksums")
    print("var checksums = []struct{ n, c string }{")
    for _ in range(n):
        x = random.randint(0, 10**9)
        print('	{"%d", "%d"},' % (x, checksum(x)))
    print("}\n")

    print("var valid = []string{")
    for _ in range(n):
        x = random.randint(0, 10**9)
        c = str(x) + str(checksum(x))
        print('	"%s",' % c)
    print("}\n")

    print("var invalid = []string{")
    for _ in range(n):
        x = random.randint(0, 10**9)
        c = str(x) + str((checksum(x)+1)%10)
        print('	"%s",' % c)
    print("}")


################################################################################

def isprime(n):
    if n < 2: return False
    for i in range(2, int(math.sqrt(n))+1):
        if n%i == 0:
            return False
    return True

def primes(a, b):
    ps = []
    for i in range(a, b+1):
        if isprime(i) and is_valid(i):
            ps.append(i)
    return ps
