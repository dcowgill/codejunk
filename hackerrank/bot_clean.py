#!/usr/bin/env python

import sys


def best_path(start, board):
    memo = {}
    def search(curpos, targets, distance, path):
        if not targets:
            return distance, path

        shortest_distance = sys.maxint
        shortest_path = None

        for i, dst in enumerate(targets):
            new_targets = targets[:i] + targets[i+1:]
            distance_to_dst = abs(curpos[0] - dst[0]) + abs(curpos[1] - dst[1])
            new_distance = distance + distance_to_dst

            key = (dst, str(new_targets))
            if memo.get(key, sys.maxint) < new_distance:
                continue
            memo[key] = new_distance

            d, p = search(dst, new_targets, new_distance, path + [dst])
            if d < shortest_distance:
                shortest_distance, shortest_path = d, p

        return shortest_distance, shortest_path

    num_rows, num_cols = len(board), len(board[0])
    dirty_nodes = []
    for x in xrange(num_rows):
        for y in xrange(num_cols):
            if board[x][y] == 'd':
                dirty_nodes.append((x, y))
    dirty_nodes.sort()

    _, best_path = search(start, dirty_nodes, 0, [])
    return best_path


_best_path_cache = None

def best_move(posx, posy, board):
    if board[posx][posy] == 'd':
        return 'CLEAN'

    global _best_path_cache
    if _best_path_cache is None:
        _best_path_cache = best_path((posx, posy), board)

    try:
        dst = next((x,y) for x,y in _best_path_cache if board[x][y]=='d')
    except StopIteration:
        return None

    if posx > dst[0]: return 'UP'
    if posx < dst[0]: return 'DOWN'
    if posy > dst[1]: return 'LEFT'
    if posy < dst[1]: return 'RIGHT'


def next_move(posx, posy, board):
    print best_move(posx, posy, board)


if __name__ == '__main__':
    grid = [
        ['b', 'd', '-', '-', 'd', ],
        ['-', 'd', '-', '-', 'd', ],
        ['-', '-', '-', 'd', '-', ],
        ['-', '-', '-', 'd', '-', ],
        ['d', 'd', 'd', '-', 'd', ],
    ]
    next_move(0, 0, grid)
