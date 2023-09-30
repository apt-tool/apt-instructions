#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting BGP attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

# make a loop to perform LP attack
for i in {1..10}
do
  BYTES=$((i * 100))
  echo "Sending $BYTES Gb data into $HOST"

  TMP=$(curl -X POST "$HOST")
  echo Got: "$TMP"
done

exit 0;