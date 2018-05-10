#!/usr/bin/env python

import sys
import string
from collections import deque, defaultdict

def findTerm( search, text ):

    # Unique set of search terms (lowercase)
    terms = set( s for s in search.lower().split(" ") )

    # List of all the words (lowercase)
    words = [t for t in text.lower().split(" ")]

    # A dictionary of term -> positions
    tPositions = defaultdict( deque )

    # The last term seen in each iteration
    lastTerm = None

    # The best slice we've found so far
    bestSlice = []

    for i, w in enumerate( words ):

        # Is this one of the search terms?
        if w in terms:

            # Add its position to the hash
            tPositions[ w ].append( i )

            # Have we seen this term more than once?
            if len( tPositions[ w ] ) > 1:

                # If this is the only term we've seen so far, drop the last position and use the latest
                if len( tPositions ) == 1:
                    tPositions[ w ].popleft()

                elif lastTerm and lastTerm != w:
                    newDistance = i - tPositions[ lastTerm ][ -1 ]
                    oldDistance = tPositions[ lastTerm ][ -1 ] - tPositions[ w ][ -2 ]

                    if newDistance < oldDistance:
                        tPositions[ w ].popleft()
                    elif len( tPositions ) > 2:
                        tPositions[ w ].pop()

            if len( tPositions ) == len( terms ):

                newSlice = sorted( p[-1] for p in tPositions.values() )
                newSliceSize = newSlice[-1] - newSlice[0]

                if not bestSlice:
                    bestSlice = newSlice
                elif newSliceSize < bestSlice[-1] - bestSlice[0]:
                    bestSlice = newSlice

            # Update the lastTerm
            lastTerm = w

    return words[ bestSlice[0] : bestSlice[-1] + 1 ]


text = "The George Washington Bridge in New York City is one of the oldest bridges ever constructed. It is now being remodeled because the bridge is a landmark. City officials say that the landmark bridge effort will create a lot of new jobs in the city. " * 10000

search = "Landmark City Bridge"

text = "x a b c a"
search = "a c y"

result = findTerm( search, text )
print result
