#!/usr/bin/env python3

"""

Good morning! Here's your coding interview problem for today.

This problem was asked by Yelp.

Given a mapping of digits to letters (as in a phone number), and a digit string,
return all possible letters the number could represent. You can assume each
valid number in the mapping is a single digit.

For example if {“2”: [“a”, “b”, “c”], 3: [“d”, “e”, “f”], …} then “23” should
return [“ad”, “ae”, “af”, “bd”, “be”, “bf”, “cd”, “ce”, “cf"].

"""

import unittest

def all_strings(mapping, digits):
    if not digits: return [""]
    result = []
    for letter in mapping[digits[0]]:
        for s in solve(mapping, digits[1:]):
            result.append(letter + s)
    return result

class TestAllStrings(unittest.TestCase):
    def test_example(self):
        mapping = {2: ["a", "b", "c"], 3: ["d", "e", "f"]}
        result = all_strings(mapping, [2, 3])
        expected = ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()
