#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting Payload attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

echo "\n";

cat /dev/urandom | tr -dc '[:alpha:]' | fold -w ${1:-20} | head -n 1

echo "\n";

ls -a;

ls -la;

cat /dev/urandom | tr -dc '[:alpha:]' | fold -w ${1:-20} | head -n 100

for _ in {1..10}
do
  TMP=$(curl -XPOST -T "big-file.iso" "$HOST")
  echo Upload result: "$TMP"

  if [[ "$TMP" == *"200"* ]]; then
    echo "System is not secure on Payload attack!"
    exit 0;
  fi

  sleep 1s;
done

sleep 5s;

echo "System is secure on Payload attack!"
