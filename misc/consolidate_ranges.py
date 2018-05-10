def consolidate_ranges(ranges):
    if not ranges:
        return
    ranges = map(sorted, sorted(ranges))
    curr_range = ranges[0]
    for next_range in ranges[1:]:
        if curr_range[1] >= next_range[0]:
            curr_range[1] = max(curr_range[1], next_range[1])
        else:
            yield curr_range
            curr_range = next_range
    yield curr_range


ranges = [
    [],
    [[1, 2], [2, 3], [3, 4]],            # => [[1, 4]]
    [[1, 2], [3, 4], [5, 6]],            # => [[1, 2], [3, 4], [5, 6]]
    [[10, 8], [0, -5], [1, 2], [5, 8]],  # => [[-5, 0], [1, 2], [5, 10]]
    [[1, 10], [5, 6]],                   # => [[1, 10]]
    [[1, 2], [0, 2], [4, 10000000]],     # => [[0, 2], [4, 10000000]]
]


for r in ranges:
    print list(consolidate_ranges(r))
