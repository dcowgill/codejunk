#!/usr/bin/env python

from __future__ import division, print_function

import arrow
import collections
import csv

Row = collections.namedtuple('Row', ['date', 'hour', 'eyeballs', 'zeroes',
                                     'completed', 'requests', 'unique_drivers'])

with open("1952852_CPJ7P.csv", "r") as fp:
    next(fp)  # skip header
    reader = csv.reader(fp)
    curdate = None
    rows = []
    for row in reader:
        if row[0]:
            curdate = row[0]
        else:
            row[0] = curdate
        date = arrow.get(row[0], "DD-MMM-YY")
        rows.append(Row(date, *map(int, row[1:])))

# Verify that there are no gaps in hours (simplifies subsequent work)
for i in range(1, len(rows)):
    assert (rows[i-1].hour + 1) % 24 == rows[i].hour

print("Which date had the most completed trips during the two week period?")
completed_by_date = collections.defaultdict(int)
for row in rows:
    completed_by_date[row.date] += row.completed
print(max(completed_by_date.items(), key=lambda p: p[1]))

print("What was the highest number of completed trips within a 24 hour period?")
print(max(sum(r.completed for r in rows[i : i+24]) for i in range(len(rows) - 24)))

print("Which hour of the day had the most requests during the two week period?")
completed_by_hour = collections.defaultdict(int)
for row in rows:
    completed_by_hour[row.hour] += row.requests
print(max(completed_by_hour.items(), key=lambda p: p[1]))

# Friday at 5pm to Sunday at 3am
def is_weekend(date):
    return ((date.weekday() == 4 and row.hour >= 17) or  # friday after 5pm
            (date.weekday() == 5) or                     # saturday
            (date.weekday() == 6 and row.hour < 3))      # sunday before 3am

print("What percentage of all zeroes during the two week period occurred on weekends?")
zeroes_by_type = collections.defaultdict(int)
for row in rows:
    zeroes_by_type['weekend' if is_weekend(row.date) else 'weekday'] += row.zeroes
print(100 * (zeroes_by_type['weekend'] / sum(zeroes_by_type.values())))

print("What is the weighted average ratio of completed trips per driver during the two week period?")
n = sum(r.completed * r.completed / r.unique_drivers for r in rows if r.unique_drivers > 0)
d = sum(r.completed for r in rows if r.unique_drivers > 0)
print(n / d)

print("In drafting a driver schedule in terms of 8 hour shifts, when are the busiest 8 consecutive hours over the two week period in terms of unique requests?")
print(max([(rows[i].hour, rows[i+7].hour, sum(r.requests for r in rows[i : i+8]))
           for i in range(len(rows) - 8)], key=lambda x: x[2]))
