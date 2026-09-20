#!/usr/bin/env bash
rm -rf log.txt
touch log.txt
nohup go run ./cmd/ > log.txt 2>&1 &
