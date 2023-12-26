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
echo "Starting Authentication access attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

# make a single requests to perform RBAC attack
HOST1="$1/login"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/signin"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/auth"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/auth/login"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/auth/signin"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/signup"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/auth/signup"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/register"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/auth/register"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/api/auth/login"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/user/login"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

HOST1="$1/users"
RESP1=$(curl "$HOST1")

echo "$HOST1";

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

for i in {1..1000}
do
  HOST1="$1/users/$i"
  RESP1=$(curl "$HOST1")

  echo "$HOST1";

  if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
  then
    echo "Got $RESP1";
    exit 0;
  fi

  sleep 2s;
done

for i in {1..1000}
do
  HOST1="$1/admins/$i"
  RESP1=$(curl "$HOST1")

  echo "$HOST1";

  if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
  then
    echo "Got $RESP1";
    exit 0;
  fi

  sleep 2s;
done

sleep 5s;

echo "Secure on Authentication access."

exit 1;
