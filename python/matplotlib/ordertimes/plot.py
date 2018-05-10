# %pylab

import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

import csv

def read_data(filename):
    rows = []
    with open(filename, 'rb') as fp:
        for row in csv.reader(fp):
            rows.append(map(float, row))
    return rows[1:]

data0 = read_data('./55a6aad1b4c50673b900025e_items.csv')
data1 = read_data('./55a6ac88c1f3aa45db000239_items.csv')
data2 = read_data('./55a7ef8dc1f3aa4f610002e7_items.csv')

df0 = pd.DataFrame(data0)
df1 = pd.DataFrame(data1)
df2 = pd.DataFrame(data2)

def get_bins():
    return [i*5 for i in range(61)]

def plot_hist(ax, x, label):
    ax.hist(x, bins=get_bins(), label=label)

fig, axes = plt.subplots(nrows=3, ncols=1, sharex=True, sharey=True)
ax0, ax1, ax2 = axes
plot_hist(ax0, df0.icol(0), 'Tue')
plot_hist(ax1, df1.icol(0), 'Wed')
plot_hist(ax2, df2.icol(0), 'Thu')
