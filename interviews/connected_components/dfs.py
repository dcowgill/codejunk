#!/usr/bin/env python

import collections
import re
import sys

class Graph(object):
    def __init__(self):
        self._edges = collections.defaultdict(set)

    def add_edge(self, v, u):
        self._edges[v].add(u)
        self._edges[u].add(v)

    def vertexes(self):
        return self._edges.iterkeys()

    def neighbors(self, v):
        return self._edges[v]

# Depth-first search
def cc_dfs(graph, start):
    cc = set()
    def f(v):
        if v in cc: return
        cc.add(v)
        for u in graph.neighbors(v):
            f(u)
    f(start)
    return cc

# Breadth-first search
def cc_bfs(graph, start):
    cc = set()
    frontier = collections.deque([start])
    while frontier:
        v = frontier.popleft()
        if v in cc: continue
        frontier.extend(graph.neighbors(v))
        cc.add(v)
    return cc

# Build the graph.
g = Graph()
for line in sys.stdin:
    m = re.match(r"(\d+) (\d+)", line)
    assert m
    v, u = int(m.group(1)), int(m.group(2))
    g.add_edge(v, u)

# Find all connected components
def ccs(graph, ccfn):
    explored = set()
    ccs = []
    for v in g.vertexes():
        if v not in explored:
            cc = ccfn(g, v)
            explored.update(cc)
            ccs.append(cc)
    return ccs

print ccs(g, cc_dfs)
print ccs(g, cc_bfs)
