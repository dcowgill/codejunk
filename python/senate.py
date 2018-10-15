#!/usr/bin/env python3

import csv

populations = {'Alabama': 4863300, 'Alaska': 741894, 'Arizona': 6931071, 'Arkansas': 2988248, 'California': 39250017, 'Colorado': 5540545, 'Connecticut': 3576452, 'Delaware': 952065, 'Florida': 20612439, 'Georgia': 10310371, 'Hawaii': 1428557, 'Idaho': 1683140, 'Illinois': 12801539, 'Indiana': 6633053, 'Iowa': 3134693, 'Kansas': 2907289, 'Kentucky': 4436974, 'Louisiana': 4681666, 'Maine': 1331479, 'Maryland': 6016447, 'Massachusetts': 6811779, 'Michigan': 9928300, 'Minnesota': 5519952, 'Mississippi': 2988726, 'Missouri': 6093000, 'Montana': 1042520, 'Nebraska': 1907116, 'Nevada': 2940058, 'New Hampshire': 1334795, 'New Jersey': 8944469, 'New Mexico': 2081015, 'New York': 19745289, 'North Carolina': 10146788, 'North Dakota': 757952, 'Ohio': 11614373, 'Oklahoma': 3923561, 'Oregon': 4093465, 'Pennsylvania': 12784227, 'Rhode Island': 1056426, 'South Carolina': 4961119, 'South Dakota': 865454, 'Tennessee': 6651194, 'Texas': 27862596, 'Utah': 3051217, 'Vermont': 624594, 'Virginia': 8411808, 'Washington': 7288000, 'West Virginia': 1831102, 'Wisconsin': 5778708, 'Wyoming': 585501}

def main():
    fp = open("./us-senate.csv")
    reader = csv.DictReader(fp)
    rows = [row for row in reader]
    senators = {'D': 0, 'R': 0}
    constituents = {'D': 0, 'R': 0}
    states = {'D': set(), 'R': set()}
    for row in rows:
        key = 'R' if row['party'] == 'republican' else 'D' # both independents caucus w/ dems
        state = row['state_name']
        population = populations[state]
        senators[key] += 1
        constituents[key] += population // 2
        states[key].add(state)
    dem_pct = 100.0 * constituents['D'] / (constituents['D'] + constituents['R'])

    print("senators = %r, constituents = %r" % (senators, constituents))
    print("each democrat senator represents {:,} people".format(constituents['D'] // senators['D']))
    print("each republican senator represents {:,} people".format(constituents['R'] // senators['R']))
    print("democrats represent {:.1f}% of the population in {} states".format(dem_pct, len(states['D'])))
    print("republicans represent {:.1f}% of the population in {} states".format(100-dem_pct, len(states['R'])))

if __name__ == '__main__':
    main()
