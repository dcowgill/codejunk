#!/usr/bin/env python

import collections
import itertools
import sys
import time


# Return all words that are an anagram of 'target'
def cowgill(target):
    target = sorted(target.lower())
    with open("/usr/share/dict/words", "r") as fp:
        for word in fp:
            word = word.strip().lower()
            if sorted(word) == target:
                yield word


def anderson(target):
    target=list(target.lower())
    corpus=[]
    with open("/usr/share/dict/words", "r") as f:
        for word in f:
            if len(word.strip()) == len(target):
                corpus.append(word.strip().lower())
    for perm in itertools.permutations(target):
        if perm in corpus:
            yield perm


def make_table(filename):
    anagrams = collections.defaultdict(set)
    with open(filename, "r") as fp:
        for w in fp:
            w = w.strip().lower()
            anagrams["".join(sorted(w))].add(w)
    return anagrams


def splits(s, n, m=0):
    if not s or not n:
        yield "", s
    else:
        for i in range(m, len(s) // 2 + 1):
            c, t = s[i], s[:i] + s[i+1:]
            for u, v in splits(t, n-1, i):
                yield c+u, v


def find_anagrams(seq, dct):
    if seq in dct:
        for s in dct[seq]:
            yield [s]
    for n in range(3, len(seq) + 1):
        for s, t in splits(seq, n):
            if s in dct:
                for v in find_anagrams(t, dct):
                    for u in dct[s]:
                        yield [u] + v


def main():
    dictionary, word = sys.argv[1], sys.argv[2]

    print("making anagram lookup dictionary...")
    table = make_table(dictionary)

    print("finding and sorting anagrams...")
    seq = "".join(sorted(word))
    anagrams = find_anagrams(seq, table)
    anagrams = sorted(anagrams, key=len)
    print("done")

    unique_anagrams = {frozenset(a) for a in anagrams}
    print("anagrams: {} ({} unique)", len(anagrams), len(unique_anagrams))

    for a in sorted(unique_anagrams, key=len):
        print(" ".join(a))

    # for a in anagrams:
    #     print(a)


if __name__ == "__main__":
    main()
