%pylab

import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

import csv

def read_data(filename):
    rows = []
    with open(filename, 'rb') as fp:
        for row in csv.reader(fp):
            rows.append(row)
    rows = rows[1:] # omit header line
    return [(map(float, row)) for row in rows]

data = read_data('./data.csv')
df = pd.DataFrame(data)

years = df[0]
violent_per_100k = (10**5)*df[3]/df[1]
murders_per_100k = (10**5)*df[5]/df[1]

fig = plt.figure()
ax0 = fig.add_subplot(1, 1, 1)
ln0 = ax0.plot(years, violent_per_100k, color='blue', label='violent crime rate')
ax0.grid()
ax0.set_xlabel('Year')
ax0.set_ylabel('Violent crimes per 100k')
ax1 = ax0.twinx()
ln1 = ax1.plot(years, murders_per_100k, color='red', label='murder rate')
ax1.set_ylabel('Murders per 100k')
lns = ln0 + ln1
labels = [ln.get_label() for ln in lns]
ax0.legend(lns, labels, loc=0)

close('all')
