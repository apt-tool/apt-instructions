#!/usr/bin/env bash


echo "System Status Report:"
echo "---------------------"

# Display current date and time
echo "Date and Time:"
date

# Display system uptime
echo -e "\nSystem Uptime:"
uptime

# Display system load average
echo -e "\nSystem Load Average:"
cat /proc/loadavg

# Display memory usage
echo -e "\nMemory Usage:"
free -h

# Display disk space usage
echo -e "\nDisk Space Usage:"
df -h

# Display top CPU-consuming processes
echo -e "\nTop CPU-consuming Processes:"
ps -eo pid,ppid,cmd,%cpu,%mem --sort=-%cpu | head -n 6

# Display top memory-consuming processes
echo -e "\nTop Memory-consuming Processes:"
ps -eo pid,ppid,cmd,%cpu,%mem --sort=-%mem | head -n 6


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
