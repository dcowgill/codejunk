import collections
import random
import string

def choose(values, m):
    n = len(values)
    s = set()
    for j in xrange(n-m, n):
        t = random.randint(0, j)
        s.add(t if t not in s else j)
    return [values[i] for i in s]

def random_string():
    return ''.join(random.choice(string.ascii_lowercase) for _ in xrange(5))

values = [random_string() for _ in xrange(20)]
c = collections.Counter(values)
for _ in xrange(100*1000):
    sample = choose(values, 4)
    for s in sample:
        c[s] += 1

total = sum(c.values())
for v in values:
    print "%s: %0.2f%%" % (v, 100.0*c[v]/total)


# initialize set S to empty
# for J := N-M + 1 to N do
#     T := RandInt(1, J)
#     if T is not in S then
#         insert T in S
#     else
#         insert J in S
