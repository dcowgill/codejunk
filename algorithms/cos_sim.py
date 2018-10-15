#!/usr/bin/env python3

import math

def cos(a, b): return dot(a, b) / (magnitude(a) * magnitude(b))
def dot(a, b): return sum(a[i]*b[i] for i in range(len(a)))
def magnitude(a): return math.sqrt(sum(x*x for x in a))

def main():
    a = [0, 1, 0, 1, 1, 1]
    b = [1, 1, 0, 1, 1, 0]
    c = [0, 0, 0, 1, 1, 1]
    print(cos(a, b))
    print(cos(a, c))
    print(cos(b, c))

if __name__ == '__main__':
    main()
