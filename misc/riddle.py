#!/usr/bin/env python

def f(a, rd):
    if not rd:
        print(a)
        return
    d, rd = rd[0], rd[1:]
    for i in xrange(len(a)):
        if a[i] == None:
            j = i + d + 1
            if j < len(a) and a[j] == None:
                a[i] = a[j] = d
                f(a, rd)
                a[i] = a[j] = None

digits = [1,2,3,4]
num = [None] * (len(digits)*2)
f(num, digits)
