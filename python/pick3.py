#!/usr/bin/env python3

import random
import sys

def pick3(num_games):
    games = sorted(shuffled(range(1, num_games+1))[:3])
    return [(n, random.choice(['home', 'away'])) for n in games]

def shuffled(seq):
    seq = list(seq)
    random.shuffle(seq)
    return seq

def main():
    try:
        num_games = int(sys.argv[1])
    except (IndexError, ValueError):
        sys.exit("usage: pick3.py num_games")
    for p in pick3(num_games):
        print("%2d %s" % p)

if __name__ == '__main__':
    main()
