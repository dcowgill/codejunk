#!/usr/bin/env python

import sys


class PrimeSieve(object):
    def __init__(self):
        self.x = 1
        self.composites = {}
        self.primes = []
        self.n = 0

    def reset(self):
        self.n = 0

    def next(self):
        if self.n >= len(self.primes):
            while True:
                self.x += 1
                if self.x in self.composites:
                    primes = self.composites[self.x]
                    for prime in primes:
                        k = self.x + prime
                        if k in self.composites:
                            self.composites[k].append(prime)
                        else:
                            self.composites[k] = [prime]
                    self.composites.pop(self.x)
                else:
                    self.composites[self.x*self.x] = [self.x]
                    break
            self.primes.append(self.x)
        self.n += 1
        return self.primes[self.n-1]

    def is_prime(self, x):
        self.reset()
        p = self.next()
        while p < x:
            p = self.next()
        return p==x


sieve = PrimeSieve()
memo = {}

def is_prime(n):
    global memo, sieve
    b = memo.get(n)
    if b is not None:
        return n
    p = sieve.is_prime(n)
    memo[n] = p
    return p


def is_lucky(n):
    digits = [int(d) for d in str(n)]
    sum_of_digits = sum(d for d in digits)
    if not is_prime(sum_of_digits):
        return False
    sum_of_square_of_digits = sum(d*d for d in digits)
    if not is_prime(sum_of_square_of_digits):
        return False
    return True


if __name__ == '__main__':
    N = int(sys.stdin.readline())
    for _ in xrange(N):
        a, b = [int(x) for x in sys.stdin.readline().strip().split(' ')]
        print a, b
        n = sum(1 for i in xrange(a, b+1) if is_lucky(i))
        print n
