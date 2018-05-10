
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

data = read_data('./lunch_data.csv')
df = pd.DataFrame(data)

def create_fill_plot(ax, g):
    x = g.mean().index
    y = g.mean().icol(1)
    yerr = 2 * g.sem().icol(1)
    ax.grid(True)
    # ax.set_ylabel('seconds')
    ax.fill_between(x, y-yerr, y+yerr, color="#3f5d7d")
    ax.plot(x, y, color='#4f6d8d', lw=2)

def create_figure(df):
    fig, axs = plt.subplots(nrows=2, ncols=1, sharey=True)
    fig.set_size_inches(10*1.33, 10, forward=True)
    create_fill_plot(axs[0], df.groupby(0))
    create_fill_plot(axs[1], df.groupby(1))
    return fig, axs

fig, axs = create_figure(df)
