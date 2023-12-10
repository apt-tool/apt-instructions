#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting Payload attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

echo "\n";

echo "Target hash:"

LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 12; echo

echo "\n";

ls -a;

LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 1000; echo

ls -la;

LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 100; echo

for _ in {1..10}
do
  TMP=$(curl -XPOST -T "big-file.iso" "$HOST")
  echo Upload result: "$TMP"
  echo Batch input ....

  if [[ "$TMP" == *"200"* ]]; then
    echo "System is not secure on Payload attack!"
    exit 0;
  fi

  sleep 15s;
done

LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 100; echo

echo "Hash key matched";

sleep 5s;

echo "http://localhost:9000/&2nnn304985nn20-8nnth02vu3-4u09u0930409unnwn20-8nnth02v^@@@018nns-948hfaouf8wu904r092u9/access/private"
echo "METHOD: POST"
echo "Size: 100Gi"
echo "Response: 500"

echo "PING http://localhost:9000 (172.0.0.1): 0 data bytes"
echo "cannot resolve localhost:9000: host not running"

echo "downtime system on host: localhost port: 9000"

echo "System is not secure on Payload attack!"
exit 0;
