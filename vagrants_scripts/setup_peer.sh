#!/bin/sh

# Run the conrold
/gladius/build/gladius-controld install
/gladius/build/gladius-controld start

# Wait for it to start
sleep 5s

# Create a new gladius account
body='{"passphrase":"password"}'

curl -s -X POST "http://localhost:3001/api/keystore/account/create" \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$body"

sleep 1s

# Fetch the discovery node state
body='{"ip": "'$1'", "passphrase":"password"}'

echo $body

curl -s -X POST "http://localhost:3001/api/p2p/state/pull" \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$body"

sleep 1s

# Get the local IP of the VM
address=$(ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | head -n 1 | tail -n 2)

# Sign a message with the VM ip
body='{"message": {"node": {"ip_address": "'$address'"}}, "passphrase": "password"}'

curl -s -X POST "http://localhost:3001/api/p2p/message/sign" \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$body"
