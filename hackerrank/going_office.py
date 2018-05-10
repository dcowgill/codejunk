#!/usr/bin/env python

import collections
import heapq
import re
import sys

def input_iter():
    for line in sys.stdin:
        yield [int(s) for s in re.findall(r"\d+", line)]

def read_input():
    it = input_iter()
    num_cities, num_roads = next(it)
    roads = collections.defaultdict(dict)
    for _ in xrange(num_roads):
        u, v, w = next(it)
        roads[u][v] = roads[v][u] = w
    src, dst = next(it)
    num_queries, = next(it)
    queries = []
    for _ in xrange(num_queries):
        u, v = next(it)
        queries.append((u, v))
    return (num_cities, roads, src, dst, queries)

def sssp(num_vertexes, edges, src, dst):
    nodes = [[sys.maxint if u != src else 0, u, True]
             for u in xrange(0, num_vertexes)]
    queue = list(nodes)
    prev = collections.defaultdict(list)
    while queue:
        heapq.heapify(queue)
        curr_dist, curr, _ = heapq.heappop(queue)
        for u,w in edges[curr].iteritems():
            if not nodes[u][2]:
                continue # already visited
            new_dist = curr_dist + w
            if new_dist < nodes[u][0]:
                nodes[u][0] = new_dist
                prev[u] = [curr]
            elif new_dist == nodes[u][0]:
                prev[u].append(curr)
        nodes[curr][2] = False # mark curr as visited
        if curr == dst:
            print nodes
            print queue
            return nodes[dst][0], prev
    return None, None

# Converts previous-link array to sets of edges
def prev2paths(prev, y, z):
    if y not in prev:
        return [set()]
    paths = []
    for x in prev[y]:
        for p in prev2paths(prev, x, y):
            p.add((x, y))
            paths.append(p)
    return paths

def main():
    data = read_input()
    num_vertexes, edges, src, dst, queries = data

    # Find edges which are on all best paths.
    best_d, prev = sssp(num_vertexes, edges, src, dst)
    best_paths = prev2paths(prev, dst, None)
    best_path_edges = set.intersection(*best_paths)

    exit(0)

    # Special case: no path from src->dst.
    if best_d < 0:
        for _ in xrange(len(queries)):
            print "Infinity"
            exit(0)

    # Run queries; skip those that don't remove a best-path edge.
    for u,v in queries:
        if (u,v) not in best_path_edges:
            print best_d
            continue
        w = edges[u].pop(v)
        edges[v].pop(u)
        d, _ = sssp(num_vertexes, edges, src, dst)
        print "Infinity" if d<0 else d
        edges[u][v] = edges[v][u] = w

if __name__ == '__main__':
    main()
