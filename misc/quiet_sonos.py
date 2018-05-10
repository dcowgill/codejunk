#!/usr/bin/env python

from __future__ import print_function

import sys

import soco

def main():
    # By default, do nothing.
    scale = 1.0

    # Parse command line.
    if len(sys.argv) >= 2:
        scale_arg = sys.argv[1]
        try:
            scale = float(scale_arg)
        except e:
            sys.exit("invalid scale factor: " + scale_arg)

    # Multiply current volume of each zone by scale.
    for zone in soco.discover():
        try:
            volume = int(zone.volume * scale)
            if volume != zone.volume:
                print("%s - changing volume from %d to %d" % (zone.player_name, zone.volume, volume))
                zone.volume = volume
        except:
            print('OOPS: an error occurred')

if __name__ == '__main__':
    main()
