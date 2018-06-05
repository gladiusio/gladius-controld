# Gladius Control Daemon

See the main [gladius-node](https://github.com/gladiusio/gladius-node) repository to see more.

## Cross compile

To compile for other systems you need to install [xgo](https://github.com/karalabe/xgo).
This is because of the Ethereum CGO bindings.

Run `make dependencies`

Then run `xgo --targets="windows/*,darwin/*,linux/*" --dest="./build/" ./cmd/gladius-controld`
from the project root. You can change the target to be whatever system you want.

## API Documentation

This document provides documentation for the Gladius Control Daemon to build interfaces on top of the Gladius Network with familiar REST API calls. If something needs more detail or explanation, please file an issue.

Throughout the document, you will see {{ETH_ADDRESS}}. This is a placeholder for either a node address or pool address in almost all cases.

## Requests
### **POST** - /api/p2p/message/sign

#### Description
Takes a message and returns a verifiable signature from the account at `/api/account/`

#### CURL

```sh
curl -X POST "http://localhost:3001/api/p2p/message/sign" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```json
{
  "message": "Message to sign goes here",
  "passphrase": "WoahASecurePassphrase"
}
```

## Requests
### **POST** - /api/p2p/message/verify

#### Description
Verifies a signature from `/api/p2p/message/sign` and checks to see if that
address is authorized for the pool.

#### CURL

```sh
curl -X POST "http://localhost:3001/api/p2p/state/verify" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```json
{
  "message": "base64encodedmessage",
  "hash": "base64encodedhash",
  "signature": "base64encodedsignature",
  "address": "0x4A97ACA4C808EE8a7C36175e31d46795d91F6CdD"
}
```

### **POST** - /api/keystore/pgp/create

#### Description
Creates a new PGP Key pair and stores the keys in `~/.config/gladius/keys` on Unix based systems and `C:\Users\USER\.gladius\keys` on Windows.

#### CURL

```sh
curl -X POST "http://localhost:3001/api/keystore/pgp/create" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"name\":\"Nate\",\"comments\":\"Anything goes here\",\"email\":\"someone@email.com\"}"
}
```

### **GET** - /api/keystore/pgp/view/public

#### Description
Retrieves the generated Public Keys from `/keystore/pgp/create`.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/keystore/pgp/view/public" \
    -H "Content-Type: text/plain"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "text/plain"
  ],
  "default": "text/plain"
}
```

### **POST** - /api/keystore/account/create

#### Description
Creates a new Ethereum account encrypted against the provided passphrase. The account will be stored in `~/.config/gladius/wallet` on Unix based systems and `C:\Users\USER\.gladius\wallet` on Windows.




#### CURL

```sh
curl -X POST "http://localhost:3001/api/keystore/account/create" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"passphrase\":\"***** Hidden credentials *****\"}"
}
```

### **GET** - /api/keystore/account

#### Description
Retrieves the list of generated accounts from `/keystore/account/create` as a JSON object.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/keystore/account"
```

### **POST** - /api/keystore/account/open

#### Description
Unlocks the Gladius account.

#### CURL

```sh
curl -X POST "http://localhost:3001/api/keystore/account/open" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"passphrase\":\"***** Hidden credentials *****\"}"
}
```

### **GET** - /api/status/tx/{{TX_HASH}}

#### Description
Retrieves the status of a given transaction hash. This info is similar to what is provided by etherscan.io. Links to the web are also provided to give to the user as feedback.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/status/tx/{{TX_HASH}}"
```

### **GET** - /api/node/

#### Description
Retrieve the node that is registered to the given wallet, as of now `/keystore/account`. An optional URL parameter can be provided to retrieve the node address of a wallet address.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/node/\
?account={{ETH_ADDRESS}}"
```

#### Query Parameters

- **account** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "{{ETH_ADDRESS}}"
  ],
  "default": "{{ETH_ADDRESS}}"
}
```

### **POST** - /api/node/create

#### Description
Creates and registers a node with the default wallet.

**`X-Authorization` is required to charge the wallet for creating a node.**

#### CURL

```sh
curl -X POST "http://localhost:3001/api/node/create" \
    -H "X-Authorization: ***** Hidden credentials *****"
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```

### **POST** - /api/node/{{ETH_ADDRESS}}/data

#### Description
Adds data to the node registered for the given wallet. The data payload *is* flexible but we are expecting the **Body Parameters** below. The data payload is encrypted automatically by using the generated PGP Keys. If these keys change, calling this again will replace that data with data encrypted by the new keys.

**`X-Authorization` is required to charge the wallet for submitting data.**

#### CURL

```sh
curl -X POST "http://localhost:3001/api/node/{{ETH_ADDRESS}}/data" \
    -H "X-Authorization: ***** Hidden credentials *****" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```
- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"name\":\"Node Demo 2\",\"email\":\"example2@gladius.io\",\"ip\":\"2.2.2.2\",\"status\":\"active\"}"
}
```

### **GET** - /api/node/{{ETH_ADDRESS}}/data

#### CURL

```sh
curl -X GET "http://localhost:3001/api/node/{{ETH_ADDRESS}}/data" \
    -H "Content-Type: text/plain; charset=utf-8" \
    --data-raw "$body"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

#### Header Parameters

- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "text/plain; charset=utf-8"
  ],
  "default": "text/plain; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "Retrieves the data POST'd in `/node/{{nodeAddress}}/data` and automatically decrypts the data using the private key that was previously generated."
}
```

### **POST** - /api/node/{{ETH_ADDRESS}}/apply/{{ETH_ADDRESS}}

#### Description
This endpoint retrieves submitted node data, decrypts it, re-encrypts it against the provided pool's public key, and submits an encrypted application for the pool. Only the pool owner can see this application and cannot be modified only overwritten with a new application.

**`X-Authorization` is required to charge the wallet for submitting an application.**

#### CURL

```sh
curl -X POST "http://localhost:3001/api/node/{{ETH_ADDRESS}}/apply/{{ETH_ADDRESS}}" \
    -H "X-Authorization: ***** Hidden credentials *****"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```
- **ResponseBodyPath_2** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```

### **GET** - /api/node/{{ETH_ADDRESS}}/application/{{ETH_ADDRESS}}

#### Description
Retrieves the current status of a submitted pool application. This endpoint also provides the available statuses as well as copy to use for displaying messages to the user.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/node/{{ETH_ADDRESS}}/application/{{ETH_ADDRESS}}"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```
- **ResponseBodyPath_2** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

### **GET** - /api/pool/{{ETH_ADDRESS}}

#### Description
Returns the public key of a given pool. You can find a list of available pools at `/market/pools`.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/pool/{{ETH_ADDRESS}}"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

### **GET** - /api/pool/{{ETH_ADDRESS}}/node/{{ETH_ADDRESS}}/application

#### Description
Retrieve the application for a given node address. The data is automatically decrypted with the private key in the keys directory.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/pool/{{ETH_ADDRESS}}/node/{{ETH_ADDRESS}}/application"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```
- **ResponseBodyPath_2** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

### **GET** - /api/pool/{{ETH_ADDRESS}}/nodes/approved

#### Description
Returns the nodes based on the provided filter. Below are the available filters:

Returns all nodes, regardless of status

- `/pool/{{ETH_ADDRESS}}/nodes/`

Returns approved nodes

- `/pool/{{ETH_ADDRESS}}/nodes/approved`

Returns rejected nodes

- `/pool/{{ETH_ADDRESS}}/nodes/rejected`


#### CURL

```sh
curl -X GET "http://localhost:3001/api/pool/{{ETH_ADDRESS}}/nodes/approved"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

### **POST** - /api/pool/{{ETH_ADDRESS}}/data

#### Description
Sets the public data payload for a pool.

**`X-Authorization` is required to charge the wallet for submitting this data.**

#### CURL

```sh
curl -X POST "http://localhost:3001/api/pool/{{ETH_ADDRESS}}/data" \
    -H "X-Authorization: ***** Hidden credentials *****" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```
- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"name\":\"Gladius Pool A\",\"location\":\"NYC - United States\",\"rating\":\"4.5\",\"nodeCount\":\"20\",\"maxBandwidth\":\"10\"}"
}
```

### **GET** - /api/pool/{{ETH_ADDRESS}}/data

#### Description
Retrieves the public data set on a given pool from `/pool/{{poolAddress}}/data`

#### CURL

```sh
curl -X GET "http://localhost:3001/api/pool/{{ETH_ADDRESS}}/data" \
    -H "X-Authorization: ***** Hidden credentials *****" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```
- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"name\":\"Gladius Pool A\",\"location\":\"NYC - United States\",\"rating\":\"4.5\",\"nodeCount\":\"20\",\"maxBandwidth\":\"10\"}"
}
```

### **PUT** - /api/pool/{{ETH_ADDRESS}}/node/{{ETH_ADDRESS}}/approve

#### Description
Approves the given node and adds it to the list of approved nodes for a pool.

**`X-Authorization` is required to charge the wallet for approving a node.**

#### CURL

```sh
curl -X PUT "http://localhost:3001/api/pool/{{ETH_ADDRESS}}/node/{{ETH_ADDRESS}}/approve" \
    -H "X-Authorization: ***** Hidden credentials *****" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```
- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"name\":\"Gladius Pool A\",\"location\":\"NYC - United States\",\"rating\":\"4.5\",\"nodeCount\":\"20\",\"maxBandwidth\":\"10\"}"
}
```

### **PUT** - /api/pool/{{ETH_ADDRESS}}/node/{{ETH_ADDRESS}}/reject

#### Description
Rejects the given node and adds it to the list of rejected nodes for a pool.

**`X-Authorization` is required to charge the wallet for rejecting a node.**

#### CURL

```sh
curl -X PUT "http://localhost:3001/api/pool/{{ETH_ADDRESS}}/node/{{ETH_ADDRESS}}/reject" \
    -H "X-Authorization: ***** Hidden credentials *****" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Path Parameters

- **ResponseBodyPath** should respect the following schema:

```
{
  "type": "string",
  "default": "{{ETH_ADDRESS}}"
}
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```
- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"name\":\"Gladius Pool A\",\"location\":\"NYC - United States\",\"rating\":\"4.5\",\"nodeCount\":\"20\",\"maxBandwidth\":\"10\"}"
}
```

### **POST** - /api/market/pools/create

#### Description
Creates a new pool with the provided public key. This key cannot be modified and a new pool would have to be created.

**`X-Authorization` is required to charge the wallet for creating a pool**

#### CURL

```sh
curl -X POST "http://localhost:3001/api/market/pools/create" \
    -H "X-Authorization: ***** Hidden credentials *****" \
    -H "Content-Type: application/json; charset=utf-8" \
    --data-raw "$body"
```

#### Header Parameters

- **X-Authorization** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "***** Hidden credentials *****"
  ],
  "default": "***** Hidden credentials *****"
}
```
- **Content-Type** should respect the following schema:

```
{
  "type": "string",
  "enum": [
    "application/json; charset=utf-8"
  ],
  "default": "application/json; charset=utf-8"
}
```

#### Body Parameters

- **body** should respect the following schema:

```
{
  "type": "string",
  "default": "{\"publicKey\":\"{{PUBLIC_KEY}}\"}"
}
```

### **GET** - /api/market/pools

#### Description
Returns a list of available pool addresses.

#### CURL

```sh
curl -X GET "http://localhost:3001/api/market/pools"
```

## References
