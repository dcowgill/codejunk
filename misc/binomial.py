import math

def binomial(n, k):
    return math.factorial(n) / (math.factorial(k) * math.factorial(n-k))

def bell(n):
    if n <= 0:
        return 1
    b = 0
    for k in range(n):
        b += binomial(n-1, k) * bell(k)
    return b

def rgs(s, r):
    n = len(s)
    v = []
    for i in range(n):
        x = r[i]
        while len(v) <= x:
            v.append([])
        v[x].append(s[i])
    return v
