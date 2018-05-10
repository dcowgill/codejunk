#!/usr/bin/env python

from __future__ import division, print_function

import cPickle
import collections
import dill
import json
import pickle
import random
import string
import time

import ujson


class A(object):
    def __init__(self, **kwargs):
        self.field0 = kwargs['field0']
        self.field1 = kwargs['field1']
        self.field2 = kwargs['field2']
        self.field3 = kwargs['field3']
        self.field4 = kwargs['field4']
        self.field5 = kwargs['field5']
        self.field6 = kwargs['field6']
        self.field7 = kwargs['field7']
        self.field8 = kwargs['field8']
        self.field9 = kwargs['field9']


class B(object):
    def __init__(self, **kwargs):
        self.a0 = kwargs['a0']
        self.a1 = kwargs['a1']
        self.a2 = kwargs['a2']


class C(object):
    def __init__(self, **kwargs):
        self.b0 = kwargs['b0']
        self.b1 = kwargs['b1']


def random_number():
    return random.randint(-10000, +10000)


def random_string():
    return ''.join(random.choice(string.printable) for _ in xrange(random.randint(5, 20)))


def random_value():
    return random.choice([random_number, random_string])()


def random_a():
    return A(**{"field" + str(i): random_value() for i in xrange(10)})


def random_b():
    return B(**{"a" + str(i): random_a() for i in xrange(3)})


def random_c():
    return C(**{"b" + str(i): random_b() for i in xrange(2)})


def assert_As_are_equal(a1, a2):
    assert isinstance(a1, A) and isinstance(a2, A)
    for k, v in a1.__dict__.iteritems():
        assert v == getattr(a2, k), "%r == %r" % (k, getattr(a2, k))


def assert_Bs_are_equal(b1, b2):
    assert isinstance(b1, B) and isinstance(b2, B)
    for k, v in b1.__dict__.iteritems():
        assert_As_are_equal(v, getattr(b2, k))


def assert_Cs_are_equal(c1, c2):
    assert isinstance(c1, C) and isinstance(c2, C)
    for k, v in c1.__dict__.iteritems():
        assert_Bs_are_equal(v, getattr(c2, k))


def a_from_dict(d):
    return A(**d)


def b_from_dict(d):
    return B(a0=a_from_dict(d['a0']),
             a1=a_from_dict(d['a1']),
             a2=a_from_dict(d['a2']))


def c_from_dict(d):
    return C(b0=b_from_dict(d['b0']),
             b1=b_from_dict(d['b1']))


def main():
    NTRIALS = 1000

    c = random_c()

    assert_Cs_are_equal(c, cPickle.loads(cPickle.dumps(c, protocol=-1)))
    assert_Cs_are_equal(c, c_from_dict(ujson.loads(ujson.dumps(c))))

    def pickle_dumps(c, n):
        for _ in xrange(n):
            pickle.dumps(c, protocol=-1)

    def pickle_loads(c, n):
        s = pickle.dumps(c, protocol=-1)
        for _ in xrange(n):
            pickle.loads(s)

    def cpickle_dumps(c, n):
        for _ in xrange(n):
            cPickle.dumps(c, protocol=-1)

    def cpickle_loads(c, n):
        s = cPickle.dumps(c, protocol=-1)
        for _ in xrange(n):
            cPickle.loads(s)

    def dill_dumps(c, n):
        for _ in xrange(n):
            dill.dumps(c, protocol=-1)

    def dill_loads(c, n):
        s = dill.dumps(c, protocol=-1)
        for _ in xrange(n):
            dill.loads(s)

    def ujson_dumps(c, n):
        for _ in xrange(n):
            ujson.dumps(c)

    def ujson_loads(c, n):
        s = ujson.dumps(c)
        for _ in xrange(n):
            ujson.loads(s)

    def ujson_loads_and_create(c, n):
        s = ujson.dumps(c)
        for _ in xrange(n):
            c_from_dict(ujson.loads(s))

    def json_loads(c, n):
        s = ujson.dumps(c)
        for _ in xrange(n):
            json.loads(s)

    def json_loads_and_create(c, n):
        s = ujson.dumps(c)
        for _ in xrange(n):
            c_from_dict(json.loads(s))

    trials = [
        ("pickle.dumps", pickle_dumps),
        ("pickle.loads", pickle_loads),
        ("cPickle.dumps", cpickle_dumps),
        ("cPickle.loads", cpickle_loads),
        ("dill.dumps", dill_dumps),
        ("dill.loads", dill_loads),
        ("ujson.dumps", ujson_dumps),
        ("ujson.loads", ujson_loads),
        ("ujson.loads*", ujson_loads_and_create),
        ("json.loads", json_loads),
        ("json.loads*", json_loads_and_create),
    ]

    label_padding = max(len(t[0]) for t in trials)
    fmt = "%{}s: %7.2f us (%6d/s)".format(label_padding)
    for label, fn in trials:
        before = time.time()
        fn(c=c, n=NTRIALS)
        elapsed = time.time() - before
        elapsed_microseconds = 1000**2 * elapsed / NTRIALS
        nps = NTRIALS / elapsed
        print(fmt % (label, elapsed_microseconds, nps))

    print("===")

    print("       cPickle.dumps: %5d bytes" % (len(cPickle.dumps(c, protocol=-1)),))
    print("          dill.dumps: %5d bytes" % (len(dill.dumps(c)),))
    print("         ujson.dumps: %5d bytes" % (len(ujson.dumps(c)),))


if __name__ == '__main__':
    main()
