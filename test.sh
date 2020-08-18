#!/bin/bash
set -e

# all api
echo 'CSV2JSON all API Paths'
curl -i 192.168.31.168:1325/
echo ''

# CSV to JSON
echo 'CSV to JSON Test'
curl -i -X POST  192.168.31.168:1325/n3-csv2json/v0.2.5/csv2json --data-binary '@./data/Modules.csv'
echo ''