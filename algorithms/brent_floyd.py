#!/usr/bin/env python

class Node(object):
    def __init__(self, value, next_node=None):
        self.value = value
        self.next = next_node

    def __repr__(self):
        return str(self.value)

def make_list():
    # 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 3
    n1 = Node(1, Node(2, Node(3, Node(4, Node(5, Node(6, Node(7)))))))
    n7 = n1.next.next.next.next.next.next
    n7.next = n1.next.next
    return n1

def floyd(f, x0):
    tortoise = x0
    hare = f(x0)
    while hare and tortoise != hare:
        tortoise = f(tortoise)
        hare = f(f(hare))
    return hare is not None

def brent(f, x0):
    p = n = 1
    tortoise = x0
    hare = f(x0)
    while hare and tortoise != hare:
        if p == n:
            tortoise = hare
            p *= 2
            n = 0
        hare = f(hare)
        n += 1
    return hare is not None

calls = [0]
def f(n):
    calls[0] += 1
    return n.next if n else None

def test():
    n = make_list()
    calls[0] = 0
    b = floyd(f, n)
    print "Floyd says:", b
    print "# of calls to f:", calls[0]
    calls[0] = 0
    b = brent(f, n)
    print "Brent says:", b
    print "# of calls to f:", calls[0]

test()
