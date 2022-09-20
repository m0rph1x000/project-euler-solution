#!/usr/bin/env python3

from itertools import takewhile


def sieve():
    D = {}
    prime = 2
    while 1:
        if prime not in D:
            yield prime
            D[prime**2] = [prime]
        else:
            for b in D[prime]:
                D.setdefault(b+prime, []).append(b)
            del D[prime]
        prime += 1


print(sum(takewhile(lambda x: x < 2_000_000, sieve())))
