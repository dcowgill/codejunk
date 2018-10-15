#!/usr/bin/env python3

import random

def pick3(num_games = 15):
    games = sorted(shuffled(range(1, num_games+1))[:3])
    return [(n, random.choice(['Home', 'Away'])) for n in games]

def shuffled(seq):
    seq = list(seq)
    random.shuffle(seq)
    return seq

def main():
    print(pick3())

if __name__ == '__main__':
    main()
