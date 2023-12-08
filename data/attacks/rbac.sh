#!/usr/bin/env bash


# message
echo "Starting RBAC attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

# make a single requests to perform RBAC attack
HOST1="$1/login"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/signin"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/auth"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/auth/login"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/auth/signin"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/signup"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/auth/signup"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/register"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

HOST1="$1/auth/register"
RESP1=$(curl "$HOST1")

if [ "$RESP1" == "HTTP/1.1 200" ] || [ "$RESP1" == "HTTP/1.1 202" ]
then
  echo "Got $RESP1";
  exit 0;
fi

sleep 5s;

echo "Secure RBAC."

exit 1;