#!/usr/bin/env python

import unittest

"""
Good morning! Here's your coding interview problem for today.

This problem was asked by Jane Street.

cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and
last element of that pair. For example, car(cons(3, 4)) returns 3, and
cdr(cons(3, 4)) returns 4.

Given this implementation of cons:

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair
Implement car and cdr.
"""

def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair


def car(pair):
    return pair(lambda a, _: a)


def cdr(pair):
    return pair(lambda _, b: b)


class TestCarCdr(unittest.TestCase):
    def test_car(self):
        self.assertEqual(3, car(cons(3, 4)))

    def test_cdr(self):
        self.assertEqual(4, cdr(cons(3, 4)))

    def test_nested(self):
        ps = cons(3, cons(4, 5))
        self.assertEqual(3, car(cons(3, cons(4, 5))))
        self.assertEqual(4, car(cdr(cons(3, cons(4, 5)))))
        self.assertEqual(5, cdr(cdr(cons(3, cons(4, 5)))))


if __name__ == '__main__':
    unittest.main()
