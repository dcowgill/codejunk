#!/usr/bin/env python

import sys

N = int(sys.stdin.readline())
days = []
for _ in xrange(N):
    days.append(int(sys.stdin.readline()))

highs = [1]
for i in xrange(1, N):
    d, j, n = days[i], i-1, 1
    while j >= 0 and d > days[j]:
        n += highs[j]
        j -= highs[j]
    highs.append(n)

print "\n".join(str(n) for n in highs)
