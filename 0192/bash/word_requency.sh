#!/bin/bash

cat words.txt |tr ' ' '\n' | sort | uniq -c | sort -k1 -n -r | awk '{print$2" "$1}'