
# PROBLEM DESCRIPTION
#
# You work at Foursquare, and they've given you a problem to solve:
# given an array of historical check-in counts by day, they want you to
# compute what they call the "N-Day High" for each of those days.
#
# Foursquare defines the "N-Day High" for any particular day as follows:
# the number of consecutive prior days with fewer check-ins, plus one.
# (They always count the current day.)
#
# Your job is to write a function that computes the N-Day Highs given a
# list of check-in counts. Here's an example input and output to such a
# function:

checkins = [3, 1, 3, 5, 7, 5, 6, 9] # one possible input
dayhighs = [1, 1, 2, 4, 5, 1, 2, 8] # corresponding expected output

def quadratic(days):
    if not days:
        return []
    highs = [1]
    for i in xrange(1, len(days)):
        d, j, n = days[i], i-1, 1
        while j >= 0 and d > days[j]:
            n += 1
            j -= 1
        highs.append(n)
    return highs

def foursquare(days):
    # ops = 0
    if not days:
        return []
    highs = [1]
    for i in xrange(1, len(days)):
        # ops += 1
        d, j, n = days[i], i-1, 1
        while j >= 0 and d > days[j]:
            # ops += 1
            n += highs[j]
            j -= highs[j]
        highs.append(n)
    # print(ops)
    return highs

# def n_day_highs(check_ins):
#     pending = []
#     results = [None] * len(check_ins)
#     i = len(check_ins) - 1
#     heapq.heappush(pending, (check_ins[i], i))
#     # ops = 0
#     while i > 0:
#         # ops += 1
#         # print("OUTER: i=%d, pending=%r, check_ins=%r" % (i, pending, check_ins))
#         i -= 1
#         while pending and check_ins[i] >= pending[0][0]:
#             # ops += 1
#             # print("INNER: check_ins[%d]=%d, pending=%r" % (i, check_ins[i], pending))
#             smallest = heapq.heappop(pending)
#             results[smallest[1]] = smallest[1] - i
#         heapq.heappush(pending, (check_ins[i], i))
#     for item in pending:
#         results[item[1]] = item[1] + 1
#     # print(ops)
#     return results

import random
import timeit

def randomdays(n):
    return [random.randrange(1, 100) for _ in xrange(n)]

results = {}
n = 1
while n <= 2**15:
    d1, d2, d3 = [], [], []
    for j in xrange(10):
        data = range(n) # randomdays(n)
        d1.append(timeit.timeit(stmt="foursquare(data)", setup="from __main__ import foursquare, data", number=10))
        d2.append(timeit.timeit(stmt="n_day_highs(data)", setup="from __main__ import n_day_highs, data", number=10))
        d3.append(timeit.timeit(stmt="quadratic(data)", setup="from __main__ import quadratic, data", number=10))
    results[n] = [d1, d2, d3]
    n *= 2

results0 = {k: v[0] for k, v in results.items()}
results1 = {k: v[1] for k, v in results.items()}
results2 = {k: v[2] for k, v in results.items()}

mean_results0 = {k: sum(v)/len(v) for k, v in results0.items()}
mean_results1 = {k: sum(v)/len(v) for k, v in results1.items()}
mean_results2 = {k: sum(v)/len(v) for k, v in results2.items()}
