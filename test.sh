#!/bin/bash
set -e

R=`tput setaf 1`
G=`tput setaf 2`
Y=`tput setaf 3`
B=`tput setaf 4`
W=`tput sgr0`

printf "\n"

ip="192.168.31.168:1325/"
base=$ip"n3-csv2json/v0.3.0/"

# all api
title="CSV2JSON all API Paths"
url=$ip
scode=`curl --write-out "%{http_code}" --silent --output /dev/null $url`
if [ $scode -ne 200 ]; then
    echo "${Y}${title}${W}"
else
    echo "${G}${title}${W}"
fi
curl -i $url
printf "\n"

# CSV to JSON
title="CSV to JSON Test"
url=$base"csv2json"
file='@./data/Modules.csv'
scode=`curl -X POST $url --data-binary $file -w "%{http_code}" -s -o /dev/null`
if [ $scode -ne 200 ]; then
    echo "${Y}${title}${W}"
else
    echo "${G}${title}${W}"
fi
curl -X POST $url --data-binary $file
printf "\n"
