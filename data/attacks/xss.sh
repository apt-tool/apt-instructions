#!/usr/bin/env bash


# message
echo "Starting XSS attack ...";

echo "performing at:";

# print date to set attack
date "+%Y/%m/%d %H:%M:%S";

# make a single requests to perform XSS attack
HOST1="$1"
RESP1=$(curl -s "$HOST1")

echo "Got $RESP1";

if [ "$RESP1" == "*<html>*" ]
then
  echo "Could add the following script which is a malware to your response!"
  echo "<script>for() { alert('got hacked!') }</script>"
  exit 0;
fi

echo "Secure to XSS!";

exit 1;