#!/usr/bin/env python

# Inspired by this tweet:
# https://twitter.com/mxcl/status/608682016205344768

from __future__ import print_function

def invert_tree(root, left_child, right_child):
    nodes = []
    def fn(node):
        nodes.append(node)
        if right_child[node]:
            fn(right_child[node])
        if left_child[node]:
            fn(left_child[node])
    fn(root)
    print(" ".join(map(str, nodes)))

invert_tree(root=3, left_child=[0, 0, 0, 2, 1, 0], right_child=[0, 0, 0, 4, 5, 0])
