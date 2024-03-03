#!/bin/bash
for file in /proc/*/status ; do awk '/^Pid|VmSwap|Name/{printf "%-16s", $2$3}END{ print ""}' $file; done | sort -k 3 -n -r | head -n 16
