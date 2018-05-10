#!/usr/bin/env python

import re
import sys
import heapq
from collections import defaultdict

TEXT = """
The George Washington Bridge in New York City is one of the oldest bridges ever constructed. It is
now being remodeled because the bridge is a landmark. City officials say that the landmark bridge
effort will create a lot of new jobs in the city.
""".strip()

SEARCH_TERMS = ["Landmark", "City", "Bridgex"]

class Solution( object ):

    NUM_INSTANCES = 0

    def __init__( self, sortedChoices, remainingClusters ):
        self.sortedChoices = sortedChoices
        self.remainingClusters = remainingClusters
        Solution.NUM_INSTANCES += 1

    def __cmp__( self, other ):
        return cmp( self.minCost(), other.minCost() )

    # True if this solution is complete
    def isComplete( self ):
        return not self.remainingClusters

    # The distance between two scalar ranges
    def rangeDist( self, lo1, hi1, lo2, hi2 ):
        if hi1 > lo2 and lo1 < hi2:
            return 0
        elif hi2 > lo1 and lo2 < hi1:
            return 0
        elif hi1 < lo2:
            return lo2 - hi1
        else:
            return lo1 - hi2

    # The minimum possible cost for this solution
    def minCost( self ):

        if not hasattr( self, "_minCost" ):

            realCost = self.sortedChoices[ -1 ] - self.sortedChoices[ 0 ]

            minRemainingCost = min(
                self.rangeDist( self.sortedChoices[ 0 ], self.sortedChoices[ -1 ], c[ 0 ], c[ -1 ] )
                for c in self.remainingClusters ) if self.remainingClusters else 0

            self._minCost = realCost + minRemainingCost

        return self._minCost


def main():

    # Split up the input text
    words = TEXT.split()

    # A pattern to clean the input
    patt = re.compile( r"\W+" )

    # Make a map of the indices that words appear at
    wordsByPos = defaultdict( list )
    for i, word in enumerate( words ):
        wordsByPos[ patt.sub( "", word.lower() ) ].append( i )

    # For each search term, make a list of the indices it appears at
    termsByPos = dict( (term, wordsByPos.get( term.lower(), [] )) for term in SEARCH_TERMS )

    # The values from the dict are the positions we need to pick one of to get the solution.
    # Put the smallest ones first to go most constrained values first.
    clusters = sorted( filter( bool, termsByPos.values() ), key = lambda c: len( c ) )
    optimalSolution = None
    heap = []
    iterations = 0

    # Make a bunch of solutions for the choices in the first cluster
    for i in clusters[ 0 ]:
        heapq.heappush( heap, Solution( [i], clusters[ 1: ] ) )

    # Keep making new solutions until one is complete
    while True:

        s = heapq.heappop( heap )

        # If it's complete, we're done!
        if s.isComplete():
            optimalSolution = s
            break

        # Otherwise make children of this solution and put them in the queue
        else:
            for i in s.remainingClusters[ 0 ]:
                childSolution = Solution( sorted( s.sortedChoices + [i] ), s.remainingClusters[ 1: ] )
                heapq.heappush( heap, childSolution )

        iterations += 1

    print words[ optimalSolution.sortedChoices[ 0 ] : optimalSolution.sortedChoices[ -1 ] + 1 ]
    print "Took %d total iterations, %d partial solutions" % (iterations, Solution.NUM_INSTANCES)

    return 0


if __name__ == "__main__":

#     import hotshot
#     import hotshot.stats

#     prof = hotshot.Profile( "snippet.prof" )
#     prof.start()

    sys.exit( main() )

#     prof.stop()
#     prof.close()
