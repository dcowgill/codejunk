#!/usr/bin/env python3

"""

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a list of integers, return the largest product that can be made by
multiplying any three integers.

For example, if the list is [-10, -10, 5, 2], we should return 500, since that's
-10 * -10 * 5.

You can assume the list has at least three integers.

"""

import random
import unittest

def brute_force(a):
    x = 0
    for i in range(len(a)):
        for j in range(len(a)):
            if i != j:
                for k in range(len(a)):
                    if i != k and j != k:
                        x = max(x, a[i] * a[j] * a[k])
    return x

def max3product(a):
    b = sorted(a)
    return max(b[0]*b[1]*b[-1], b[-3]*b[-2]*b[-1])

def random_ints(n):
    return [random.randint(-1000, 1000) for _ in range(n)]

class TestMax3Product(unittest.TestCase):
    def test_random(self):
        ntrials = 100
        for i in range(ntrials):
            a = random_ints(50)
            self.assertEqual(brute_force(a), max3product(a))

if __name__ == '__main__':
    unittest.main()
