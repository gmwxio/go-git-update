#!/bin/bash

set -x 

mkdir ~/git-remote
cd ~/git-remote
mkdir gcon.git
cd gcon.git
git init --bare
cd ..
pwd
git daemon \
    --reuseaddr \
    --verbose \
    --base-path=`pwd` \
    --export-all \
    --enable=receive-pack \
    ~/git-remote