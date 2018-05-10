#!/usr/bin/env python

# A simulation-based proof of the Monty Hall problem for those who
# refuse to be persuaded by the mathematical explanation.

import json
import random

GOAT = 'G'
CAR = 'C'
NUM_DOORS = 3

def random_door():
    return random.randrange(0, NUM_DOORS)

def new_doors():
    doors = [GOAT] * NUM_DOORS
    doors[random_door()] = CAR
    return doors

def play(switch):
    doors = new_doors()
    all_choices = set(range(NUM_DOORS))
    car = doors.index(CAR)
    choice = random_door()
    possible_reveals = all_choices - {choice, car}
    reveal = random.choice(list(possible_reveals))
    alternatives = all_choices - {choice, reveal}
    if switch:
        choice = random.choice(list(alternatives - {choice}))
    return doors[choice] == CAR

def simulate(switch, ntrials):
    wins = 0
    for _ in range(ntrials):
        if play(switch):
            wins += 1
    return 100.0 * wins / ntrials

if __name__ == '__main__':
    print("Win rate for stay:   %.1f%%" % simulate(False, 1000))
    print("Win rate for switch: %.1f%%" % simulate(True, 1000))
