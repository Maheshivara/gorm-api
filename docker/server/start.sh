#!/bin/sh

echo "Waiting 5s for database start"
echo -n "["
i=0
while [ "$i" -le 25 ]; do
    echo -n "#"
    sleep 0.2
    i=$(( i + 1 ))
done 
echo "]"

echo "Starting API server"
/api/app