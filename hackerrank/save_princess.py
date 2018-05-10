#!/usr/bin/env python

def displayPathtoPrincess(n, grid):
    if grid[0][0] == 'p':
        dirs = ('UP', 'LEFT')
    elif grid[0][n-1] == 'p':
        dirs = ('UP', 'RIGHT')
    elif grid[n-1][0] == 'p':
        dirs = ('DOWN', 'LEFT')
    elif grid[n-1][n-1] == 'p':
        dirs = ('DOWN', 'RIGHT')
    else:
        raise Exception("princess not found")

    distance = int(n / 2)
    for _ in range(distance):
        print(dirs[0])
    for _ in range(distance):
        print(dirs[1])

if __name__ == '__main__':
    grid = [
        ['-', '-', 'p'],
        ['-', 'm', '-'],
        ['-', '-', '-'],
    ]
    displayPathtoPrincess(len(grid[0]), grid)
