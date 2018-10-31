#!/usr/bin/env python3

import random
import sys

def pick3(num_games):
    games = sorted(shuffled(range(1, num_games+1))[:3])
    return [(n, random.choice(['Home', 'Away'])) for n in games]

def shuffled(seq):
    seq = list(seq)
    random.shuffle(seq)
    return seq

def main():
    num_games = int(sys.argv[1]) if len(sys.argv) == 2 else 14
    print(pick3(num_games))

if __name__ == '__main__':
    main()
