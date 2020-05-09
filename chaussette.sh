#!/bin/bash -

port=$1
max_ip=$2
stats=$3

if test  -z  $port || test  -z $max_ip ; then
   echo "specify port please and the max ip and \"stats\" in option"
   exit 0
fi

ss_result=$(ss -antu  "sport = :$port" | awk '{print$NF}' | grep -vEi "address|\*" | awk -F ":" '!($NF="")' | sort -n | uniq -c | sort -n | tail -n $max_ip)

counter=0
for i in  $ss_result ; do
   if ! [ -z $stats ] ; then
      echo -e $i
   else
      if ! (( $counter % 2 == 0 )); then
         echo $i
      fi
   fi
   ((counter=counter+1))
done
