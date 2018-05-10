#!/usr/bin/env python

"""
Problem description taken from
https://www.interviewcake.com/question/java/cake-thief

You are a renowned thief who has recently switched from stealing
precious metals to stealing cakes because of the insane profit margins.
You end up hitting the jackpot, breaking into the world's largest
privately owned stock of cakes—the vault of the Queen of England.

While Queen Elizabeth has a limited number of types of cake, she has an
unlimited supply of each type.

Each type of cake has a weight and a value, stored in objects of a
CakeType class:

    class CakeType {
        int weight;
        int value;
        public CakeType(int weight, int value) {
            this.weight = weight;
            this.value  = value;
        }
    }

For example:

    // weighs 7 kilograms and has a value of 160 pounds
    new CakeType(7, 160);

    // weighs 3 kilograms and has a value of 90 pounds
    new CakeType(3, 90);

You brought a duffel bag that can hold limited weight, and you want to
make off with the most valuable haul possible.

Write a function maxDuffelBagValue() that takes an array of cake type
objects and a weight capacity, and returns the maximum monetary value
the duffel bag can hold.

For example:

    CakeType[] cakeTypes = new CakeType[] {
        new CakeType(7, 160),
        new CakeType(3, 90),
        new CakeType(2, 15),
    };

    int capacity = 20;

    maxDuffelBagValue(cakeTypes, capacity);
    // returns 555 (6 of the middle type of cake and 1 of the last type of cake)

Weights and values may be any non-negative integer. Yes, it's weird to
think about cakes that weigh nothing or duffel bags that can't hold
anything. But we're not just super mastermind criminals—we're also
meticulous about keeping our algorithms flexible and comprehensive.
"""

def max_duffel_bag_value(cake_types, capacity):
    if capacity <= 0:
        return 0

    # Remove worthless cakes.
    cake_types = [(w, v) for w, v in cake_types if v > 0]

    # If any cake has zero weight, we can store infinite value.
    if any(weight <= 0 for weight, _ in cake_types):
        return float('inf')

    # Dynamic programming approach.
    cache = {}
    def f(n):
        if n == 0:
            return 0
        if cache.has_key(n):
            return cache[n]
        best = 0
        for weight, value in cake_types:
            if n >= weight:
                v = f(n - weight) + value
                if v > best:
                    best = v
        cache[n] = best
        return best

    return f(capacity)


def main():
    print(max_duffel_bag_value([(7, 160), (3, 90), (2, 15)], 20))


if __name__ == '__main__':
    main()
