#!/bin/bash

# Merges repositories, preserving history.
# Don't use this.

set -e
set -x

project=$1

rm -rf $project
git clone git@github.com:omakasecorp/$project.git
cd $project
git remote rm origin
files=$(find . -depth 1 ! -name .git ! -name $project)
mkdir -p $project
git mv ${files} $project
git commit -q -m "Prepare for mono-repo merge" >/dev/null
cd ../omakasemono
git remote add $project ../$project
git pull -q --no-edit $project master
git remote rm $project
cd ..
