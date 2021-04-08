#!/bin/sh
r=0; e=`tput colors`
while [ $r -lt $e ]
do  for c in 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15
    do  col=$((r+c))
        tput setaf $col
        echo -n " $col"
    done
    tput sgr0
    echo
    r=$((r+16))
done
