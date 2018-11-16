#!/usr/bin/env python3

def power_set(xs):
    ps = 2**len(xs)
    for c in range(ps):
        subset = []
        for i, x in enumerate(xs):
            if c&(1<<i):
                subset.append(x)
        yield subset

def all_partitions(xs):
    if len(xs) == 1:
        yield [[xs]]
    elif len(xs) == 0:
        yield []
    else:
        x = xs[0]
        for p in all_partitions(xs[1:]):
            yield [[x]] + p
            for i in range(len(p)):
                yield [[x] + s if i==j else s for j, s in enumerate(p)]

def all_part2(xs):
    if 0 <= len(xs) <= 1:
        raise Exception("need at least two values")
    elif len(xs) == 2:
        yield (xs[:1], xs[1:])
    else:
        x, rest = xs[0], xs[1:]
        yield ([x], rest)
        for p in all_part2(rest):
            yield ([x] + p[0], p[1])
            yield (p[0], [x] + p[1])
