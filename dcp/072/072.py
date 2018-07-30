#!/usr/bin/env python3

"""
Good morning! Here's your coding interview problem for today.

This problem was asked by Google.

In a directed graph, each node is assigned an uppercase letter. We define a
path's value as the number of most frequently-occurring letter along that path.
For example, if a path in the graph goes through "ABACA", the value of the path
is 3, since there are 3 occurrences of 'A' on the path.

Given a graph with n nodes and m directed edges, return the largest value path
of the graph. If the largest value is infinite, then return null.

The graph is represented with a string and an edge list. The i-th character
represents the uppercase letter of the i-th node. Each tuple in the edge list
(i, j) means there is a directed edge from the i-th node to the j-th node.
Self-edges are possible, as well as multi-edges.

For example, the following input graph:

ABACA
[(0, 1),
 (0, 2),
 (2, 3),
 (3, 4)]

Would have maximum value 3 using the path of vertices [0, 2, 3, 4], (A, A, C, A).

The following input graph:

A
[(0, 0)]

Should return null, since we have an infinite loop.
"""

import collections

# TODO: find a more efficient solution!

class Cycle(Exception):
    pass

def brute_force(vertices, edge_list):
    # Convert array of edges to a mapping from src -> [dst1, dst2, ...]
    edges = collections.defaultdict(list)
    for edge in edge_list:
        edges[edge[0]].append(edge[1])

    # Search depth-first for the path with the highest value.
    visited = [False] * len(vertices)
    def dfs(vert, start_letter):
        if visited[vert]: raise Cycle
        visited[vert] = True
        value = 1 if vertices[vert] == start_letter else 0
        for dst in edges.get(vert, []):
            value += dfs(dst, start_letter)
        visited[vert] = False
        return value

    # Run a dfs from every possible starting vertex.
    try:
        return max(dfs(vert, letter) for vert, letter in enumerate(vertices))
    except Cycle:
        return None

def main():
    print(brute_force('ABACA', [(0, 1), (0, 2), (2, 3), (3, 4)]))
    print(brute_force('A', [(0, 0)]))

if __name__ == '__main__':
    main()
