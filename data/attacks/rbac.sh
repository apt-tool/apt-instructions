#!/usr/bin/env bash


# message
echo "Starting RBAC attack ...";

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

echo "Keygen"
LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 100; echo

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

HOST1="$1/api/register"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/api/login"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/api/users"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/api/users/register"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/api/users/1"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/api/admins"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/auth/users/2"
RESP1=$(curl "$HOST1")

echo "$HOST1";

sleep 5s;

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

LC_ALL=C tr -dc '[:graph:]' </dev/urandom | head -c 100; echo

echo "20 valid endpoints checked, some may got failer error."

echo "Secure on RBAC attack."

exit 1;
