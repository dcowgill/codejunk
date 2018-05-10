#!/usr/bin/env python

# Finds connected component subgraphs

import collections
import re
import sys

edges = collections.defaultdict(list)
for line in sys.stdin:
    m = re.match(r"(\d+) (\d+)", line)
    v, u = int(m.group(1)), int(m.group(2))
    edges[v].append(u)
    edges[u].append(v)

def findcc(edges):
    start = edges.iterkeys().next()
    frontier = collections.deque([start])
    cc = []
    while frontier:
        v = frontier.popleft()
        if v in edges:
            frontier.extend(edges.pop(v))
            cc.append(v)
    return cc

while edges:
    print findcc(edges)
