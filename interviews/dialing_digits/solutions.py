#!/usr/bin/env python

import sys

adj = {
     1: [6, 8],    2: [7, 9], 3: [4, 8],
     4: [3, 9, 0], 5: [],     6: [1, 7, 0],
     7: [2, 6],    8: [1, 3], 9: [2, 4],
                   0: [4, 6]
}

# Brute force solution enumerates all permutations: O(3^n).
def naive(n):
     count = [0]
     def dfs(v, d):
          d.append(v)
          if len(d) == n:
               count[0] += 1
               return
          for w in adj[v]:
               dfs(w, list(d))
     for v in adj.keys():
          dfs(v, [])
     return count[0]

# Use Dynamic Programming for a O(n) solution.
def smart(n):
     cnt1 = [1] * 10
     cnt2 = [0] * 10
     for i in range(n-1):
          for d, neighbours in adj.iteritems():
               for n in neighbours:
                    cnt2[n] += cnt1[d]
          cnt1 = cnt2
          cnt2 = [0] * 10
     return sum(cnt1)

N = int(sys.stdin.readline().strip())
print smart(N)
