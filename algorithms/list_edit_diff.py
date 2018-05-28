#!/usr/bin/env python

from __future__ import print_function

import random

random.seed(0)


def random_swap(a):
    i = random.randrange(0, len(a))
    j = random.randrange(0, len(a))
    a[i], a[j] = a[j], a[i]


def display(a, b, d):
    assert len(a) == len(b) == len(d)
    for i in range(len(a)):
        print("%3s %3s %3s" % (a[i], b[i], 'X' if d[i] else ''))


def find_edits(a, b):
    assert len(a) == len(b)
    b = [x for x in b] # make a copy of b
    d = [False] * len(a)
    done = False
    while not done:
        done = True
        for i in range(len(a)):
            if a[i] == b[i]:
                continue
            j = b.index(a[i])
            b[i], b[j] = b[j], b[i]
            d[i] = d[j] = True
            done = False
            break
    return d


def main():
    tests = [
        [range(1, 7), [4, 3, 1, 2, 5, 6]],
        [range(5), [4, 1, 2, 3, 0]],
        [range(1, 7), [1, 2, 5, 6, 3, 4]],
    ]
    # for num_swaps in range(1, 6):
    #     a = range(5)
    #     b = range(5)
    #     for _ in range(num_swaps):
    #         random_swap(b)
    #     tests.append([a, b])

    for a, b in tests:
        print('--------------------')
        display(a, b, find_edits(a, b))


if __name__ == '__main__':
    main()
