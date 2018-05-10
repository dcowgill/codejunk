#!/usr/bin/env python

from __future__ import print_function
import collections
import re
import sys


class TermHistogram(object):
    """
    A histogram of search terms.

    This class has one constraint: every term should appear a minimum a one
    time in the histogram. The ``valid`` method will return true IFF the
    constraint currently holds.

    Attempting to add to or remove from the histogram a string that isn't a
    valid term (as provided to the constructor) has no effect.
    """
    def __init__(self, terms, words=None):
        self._terms = set(terms)
        self._hist = collections.defaultdict(int)
        if words is not None:
            for w in words:
                if w in self._terms:
                    self._hist[w] += 1
        self._refresh_validity()

    def add(self, word):
        if word not in self._terms:
            return
        oldval = self._hist[word]
        self._hist[word] += 1
        # Did we just go from 0 to 1? If so, we might be valid again.
        if oldval == 0:
            self._refresh_validity()

    def remove(self, word):
        if word not in self._terms:
            return
        oldval = self._hist[word]
        self._hist[word] -= 1
        # Did we just go from 1 to 0? If so, we're invalid.
        if oldval == 1:
            self._valid = False

    def valid(self):
        return self._valid

    def _refresh_validity(self):
        self._valid = all(self._hist[t] for t in self._terms)


class ShortestSubstringProblem(object):
    """
    Represents a particular instance of the shortest-substring problem. This
    is just a dumb data structure; treat as immutable once constructed.

    Exposes the following attributes:

    * ``terms``: A set containing the (lowercased) search terms.
    * ``text``: The original, unadulterated content.
    * ``words``: The sequence of (lowercased) words found in ``text``.
    * ``word_positions``: The position (begin, end) of each word in ``words``.
    """
    def __init__(self, terms, text):
        self.terms = set(t.lower() for t in terms)
        self.text = text
        self.words = []
        self.word_positions = []

        for match in re.finditer(r"[a-z-]+", self.text, re.I):
            begin, end = match.span()
            self.words.append(self.text[begin : end].lower())
            self.word_positions.append((begin, end))

    def word_span(self, i, j):
        """Returns the span (begin, end) from word i through word j."""
        return (self.word_positions[i][0], self.word_positions[j][1])

    @classmethod
    def from_file(cls, fp):
        """
        Constructs a problem data set from an input stream. The input format is as
        follows: the first line must contain a comma-separated list of search
        terms; the remainder of the file is the text to search. For example:

            >>> with open("/tmp/example.input", "w") as f:
            ...     f.write("quick,the,dog\n")
            ...     f.write("The quick brown fox jumped over the lazy dog.\n")
            ...     f.write("The quick brown fox.")

            >>> repr(theoretical_solver("/tmp/example.input"))
            "'lazy dog.\\nThe quick'"
        """
        terms = [s for s in fp.readline().strip().split(",")]
        text = fp.read()
        return cls(terms, text)


def solve(problem):
    """
    Given a shortest-substring problem, returns the min-length span containing
    every search term at least once. If no such span exists, returns None.
    """
    histogram = TermHistogram(problem.terms, problem.words)
    if not histogram.valid():
        return None

    # Find the shortest substring beginning with the first word. We accomplish
    # this by removing words from the end of the string until our histogram
    # constraint is violated, then adding the last word back to the histogram.
    j = num_words = len(problem.words)
    while histogram.valid():
        j -= 1
        histogram.remove(problem.words[j])
    histogram.add(problem.words[j])
    substring = problem.word_span(0, j)

    # For each i in [0, N], remove the ith word from our histogram, then
    # advance j until the histogram constraint is restored. If that span is
    # shorter than the currently shortest one, it's the new winner.
    for i in xrange(num_words):
        histogram.remove(problem.words[i])
        while not histogram.valid() and j < num_words-1:
            j += 1
            histogram.add(problem.words[j])
        if not histogram.valid():
            break
        begin, end = problem.word_span(i+1, j)
        if end-begin < substring[1]-substring[0]:
            substring = (begin, end)

    return substring


def main():
    problem = ShortestSubstringProblem.from_file(sys.stdin)
    span = solve(problem)
    if span:
        begin, end = span
        print(problem.text[begin : end])


if __name__ == '__main__':
    main()
