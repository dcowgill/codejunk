#!/usr/bin/env python3

"""

Good morning! Here's your coding interview problem for today.

This problem was asked by Facebook.

Given a multiset of integers, return whether it can be partitioned into two
subsets whose sums are the same.

For example, given the multiset {15, 5, 20, 10, 35, 15, 10}, it would return
true, since we can split it up into {15, 5, 10, 15, 10} and {20, 35}, which both
add up to 55.

Given the multiset {15, 5, 20, 10, 35}, it would return false, since we can't
split it up into two subsets that add up to the same sum.

"""

# TODO: this solution generates all possible two-subset partitions, so it's not
# going to work for large N; try to find a better approach.

def all_partitions_into_two_subsets(xs):
    if 0 <= len(xs) <= 1:
        yield
    elif len(xs) == 2:
        yield (xs[:1], xs[1:])
    else:
        x, rest = xs[0], xs[1:]
        yield ([x], rest)
        for p in all_partitions_into_two_subsets(rest):
            yield ([x] + p[0], p[1])
            yield (p[0], [x] + p[1])

def partition_into_two_equal_subsets(xs):
    for p in all_partitions_into_two_subsets(xs):
        first, second = p
        if sum(first) == sum(second):
            return p
    return None

def main():
    xs = [15, 5, 20, 10, 35, 15, 10]
    part = partition_into_two_equal_subsets(xs)
    if part:
        for s in part:
            print("sum(%s) = %d" % (s, sum(s)))
    else:
        print("failed")

if __name__ == '__main__':
    main()
