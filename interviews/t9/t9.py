#!/usr/bin/env python

import sys

DIGITS = range(2, 10)
KEYPAD = [
    "",     # 0
    "",     # 1
    "abc",  # 2
    "def",  # 3
    "ghi",  # 4
    "jkl",  # 5
    "mno",  # 6
    "pqrs", # 7
    "tuv",  # 8
    "wxyz", # 9
]

# Returns the set of all prefixes of the words in the given file.
def create_dictionary_from_file(path):
    dictionary = set()
    with open(path, 'r') as f:
        for word in f:
            word = word.strip().lower()
            dictionary.update(word[:n] for n in xrange(1, len(word)+1))
    return dictionary

# Returns the successor set of dictionary prefixes formed by appending each of
# the digit d's letters to each of the prefixes in ps.
def successor_prefixes(dictionary, d, ps):
    a = []
    for p in ps:
        for l in KEYPAD[d]:
            s = p + l
            if s in dictionary:
                a.append(s)
    return a

# Builds a tree beginning with the given prefixes.
def create_tree(dictionary, prefixes):
    if not prefixes:
        prefixes = [""]
    children = {}
    tree = (prefixes, children)
    for d in DIGITS:
        a = successor_prefixes(dictionary, d, prefixes)
        if a:
            children[d] = create_tree(dictionary, a)
    return tree

# Searches tree for the set of prefixes matching the keypad input in digits.
def lookup(tree, digits):
    for d in digits:
        if d not in DIGITS:
            continue
        tree = tree[1].get(d)
        if not tree: return
    return tree[0]

def main():
    dictionary = create_dictionary_from_file(sys.argv[1])
    tree = create_tree(dictionary, None)
    for s in sys.stdin:
        digits = [int(c) for c in s.strip()]
        prefixes = lookup(tree, digits)
        print ",".join(prefixes) if prefixes else ""

if __name__ == "__main__":
    main()
