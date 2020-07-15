#!/bin/bash

awk '
{
  for (i = 1; i <= NF; i++) {
        if(NR == 10) {
            s = $0;
        } else {
          continue
        }
    }
} 
END {
  print s
}
' file.txt

# or use sed
sed -n 10p file.txt