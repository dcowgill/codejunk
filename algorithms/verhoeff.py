#!/usr/bin/env python3

# https://en.wikipedia.org/wiki/Verhoeff_algorithm

import math

def isprime(n):
    for i in range(2, int(math.sqrt(n))+1):
        if n%i == 0:
            return False
    return True

dTab = [
    [0,1,2,3,4,5,6,7,8,9],
    [1,2,3,4,0,6,7,8,9,5],
    [2,3,4,0,1,7,8,9,5,6],
    [3,4,0,1,2,8,9,5,6,7],
    [4,0,1,2,3,9,5,6,7,8],
    [5,9,8,7,6,0,4,3,2,1],
    [6,5,9,8,7,1,0,4,3,2],
    [7,6,5,9,8,2,1,0,4,3],
    [8,7,6,5,9,3,2,1,0,4],
    [9,8,7,6,5,4,3,2,1,0],
]

def d(j, k):
    return dTab[j][k]

pTab = [
    [0,1,2,3,4,5,6,7,8,9],
    [1,5,7,6,2,8,3,0,9,4],
    [5,8,0,3,7,9,6,1,4,2],
    [8,9,1,6,0,4,3,5,2,7],
    [9,4,5,3,1,2,6,8,7,0],
    [4,2,8,6,5,7,3,9,0,1],
    [2,7,9,3,8,0,6,4,1,5],
    [7,0,4,6,9,1,3,2,5,8],
]

def p(pos, digit):
    return pTab[pos%8][digit]

invTab = [0,4,3,2,1,5,6,7,8,9]

def inv(digit):
    return invTab[digit]

def checksum(n):
    ns = [0] + list(map(int, reversed(str(n))))
    c = 0
    for i in range(len(ns)):
        c = d(c, p(i%8, ns[i]))
    return inv(c)

def is_valid(n):
    ns = list(map(int, reversed(str(n))))
    c = 0
    for i in range(len(ns)):
        c = d(c, p(i%8, ns[i]))
    return c == 0

def primes(a, b):
    ps = []
    for i in range(a, b+1):
        if isprime(i) and is_valid(i):
            ps.append(i)
    return ps
