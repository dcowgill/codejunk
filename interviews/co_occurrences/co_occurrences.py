#!/usr/bin/env python

import collections
import csv
import sys

# Parse the number of lists.
N = int(sys.stdin.readline())

# Read lists of ints from stdin and count how many times each one occurs.
items = collections.defaultdict(int)
rows = []
for row in csv.reader(sys.stdin):
    row = [int(s) for s in row]
    rows.append(row)
    for i in row:
        items[i] += 1

# Count occurrences of unique pairs of items in each row. We may immediately
# omit items which appeared fewer than N times in the input.
pairs = collections.defaultdict(int)
for row in rows:
    row = [i for i in row if items[i] >= N]
    for a in row:
        for b in row:
            if a < b:
                pairs[(a,b)] += 1

# Filter pairs with fewer than N appearances, then print in sorted order.
pairs = ((pair,n) for pair,n in pairs.iteritems() if n >= N)
for pair, n in sorted(pairs, lambda a,b: cmp(a[0], b[0])):
    print "%d,%d" % pair
