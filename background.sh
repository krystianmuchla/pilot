#!/usr/bin/env bash
rm -rf log.txt
touch log.txt
nohup ./pilot > log.txt 2>&1 &
