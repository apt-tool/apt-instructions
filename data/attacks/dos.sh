#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting DoS attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

echo "\n";

# make a loop to perform DoS attack
for _ in {1..100}
do
  TMP=$(curl -i -X GET "$HOST")
  echo Got: "$TMP"

  if [[ "$TMP" == *"400"* ]]; then
    echo "System is secure on DoS attack!"
    exit 1;
  fi

  sleep 1s;
done

echo "Sent 100 http requests!";

sleep 5s;

echo "System did not block me!";

exit 0;
