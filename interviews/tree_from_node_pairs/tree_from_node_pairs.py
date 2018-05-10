#!/usr/bin/env python

from __future__ import print_function
import collections
import sys


# Given a list of pairs `(a,b)` where `a` is the number of a node and `b` is its
# parent, where the nodes and edges form a tree with a single root, construct
# the tree and return its root.

class Node(object):
    def __init__(self):
        self.name = None
        self.children = {}
        self.parent = None

    def __repr__(self):
        return "Node<%r, %r>" % (self.name, self.children)


def build_tree(pairs):
    # Build the tree
    nodes = collections.defaultdict(Node)
    for c, p in pairs:
        parent = nodes[p]
        child = nodes[c]
        parent.name = p
        child.name = c
        parent.children[c] = child
        child.parent = p

    # Choose an arbitrary node, then traverse up to the root
    n = next(nodes.itervalues())
    while n.parent is not None:
        n = n.parent
    return n


def main():
    # Read the pairs from stdin. Expected format: "A B\nC B\n" etc.
    pairs = []
    for line in sys.stdin:
        pairs.append(line.strip().split())

    tree = build_tree(pairs)
    print(tree.name)


if __name__ == '__main__':
    main()
