#!/bin/bash -

set -x

port=$1
max_minute=$2

if test  -z  $port || test -z $max_minute ; then
   echo "specify port and max minute please"
   exit 0
fi

if ! type at &> /dev/null ; then
   echo "at not found, trying to install it.."
   apt install -y at
fi

ipt_rule="iptables -A INPUT -p tcp --syn --dport $port -m connlimit --connlimit-above 5 -j DROP"

$ipt_rule

echo $ipt_rule | sed -e 's/-A/-D/' | at now+${max_minute}minutes
