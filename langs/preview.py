import random
from collections import defaultdict
 
printdead, printlive = '-#'
maxgenerations = 3
cellcount = 3, 3
celltable = defaultdict(int, {
 (1, 2): 1, 
 (1, 3): 1, 
 (0, 3): 1, 
 } )

u = universe = defaultdict(int)
u[(1, 0)], u[(1, 1)], u[(1, 2)] = 1, 1, 1

for i in range(maxgenerations):
    print "\nGeneration %3i:" % ( i, )
    for row in range(cellcount[1]):
        print "  ", ''.join(str(universe[(row, col)])
                            for col in range(cellcount[0])).replace(
                                '0', printdead).replace('1', printlive)
    nextgeneration = defaultdict(int)
    for row in range(cellcount[1]):
        for col in range(cellcount[0]):
            nextgeneration[(row, col)] = celltable[
                ( universe[(row, col)], 
                  -universe[(row, col)] + sum(universe[(r, c)]
                                             for r in range(row-1, row+2)
                                             for c in range(col-1, col+2) )
                ) ]
    universe = nextgeneration
