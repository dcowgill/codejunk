#!/usr/bin/env python

from __future__ import print_function
import random
import string


def gen_words(n):
    words = set()
    for _ in xrange(n):
        wordlen = random.randint(5,10)
        word = ''.join(random.choice(string.ascii_lowercase) for _ in xrange(wordlen))
        words.add(word)
    return words


def gen_big_example_test_case():
    terms = ["Landmark", "City", "Bridge"]
    text = "The George Washington Bridge in New York City is one of the oldest bridges ever constructed. It is now being remodeled because the bridge is a landmark. City officials say that the landmark bridge effort will create a lot of new jobs in the city. " * 10000
    return (terms, text)


# (A, B*n, C) and (A, C)
def gen_pathological_test_case(words, n):
    assert len(words) >= 3
    a, b, c = list(words)[:3]
    words = [a] + [b]*n + [c]
    terms = [a, c]
    return (terms, wrapjoin(words))


def wrapjoin(words, max_line_len=72):
    lines = []
    curline = []
    for w in words:
        curlen = sum(len(s) for s in curline) + len(curline)
        if curlen + len(w) + 1 > 72:
            lines.append(" ".join(curline))
            curline = []
        curline.append(w)
    lines.append(" ".join(curline))
    return "\n".join(lines)


if __name__ == '__main__':
    terms, text = gen_big_example_test_case()
    print(",".join(terms))
    print(text)
