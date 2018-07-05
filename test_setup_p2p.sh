#!/bin/sh

# Run the conrold
/gladius/build/gladius-controld install
/gladius/build/gladius-controld start

# Wait for it to start
sleep 5s

# Create a new gladius account
body='{"passphrase":"password"}'

curl -X POST "http://localhost:3001/api/keystore/wallet/create" \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$body"

sleep 1s

# Fetch the discovery node state
body='{"ip":"192.168.10.2.", "passphrase":"password"}'

curl -X POST "http://localhost:3001/api//state/pull" \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$body"

sleep 1s

# Push a message with the machine IP in it
body='{"ip":"172.28.128.2", "passphrase":"password"}'

curl -X POST "http://localhost:3001/api//state/pull" \
    -H "Content-Type: application/json; charset=utf-8" \
    -d "$body"
