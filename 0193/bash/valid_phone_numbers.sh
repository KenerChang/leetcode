#!/bin/bash

cat file.txt | grep -E '(\(\d{3}\) |\d{3}-)\d{3}-\d{4}' # works in mac
# cat file.txt | grep -P '(\(\d{3}\) |\d{3}-)\d{3}-\d{4}'  works in linux environemnt