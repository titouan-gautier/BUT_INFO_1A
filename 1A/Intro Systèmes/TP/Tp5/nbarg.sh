# !/bin/bash

((a=1))
((b=2))
((c=4))

echo $a
echo $b
echo $c

echo $#

((d=2*$#))

echo $d

wc -l nbarg.sh