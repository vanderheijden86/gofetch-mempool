# Ethereum Mempool Explorer

Golang implementation of an ethereum mempool explorer. It uses the go-ethereum `ethclient` and `gethclient` packages
to connect to an Ethereum node. It connects to an Infura node in this case, but that could just as easily be a
locally running eth node if you change the configuration in [ethclients.go](cmd/datacollection/ethclients.go).
It starts a websocket server to publish the full mempool TXs in JSON format for other clients (like a front-end) to
consume.

## Setup

In the project's `root folder` directory, you need to place a file called `.env` in which to put the websocket and https
URLs of your ethereum node. Easiest is to rename the`example.env`file to `.env`and paste in your own URLs. These could
for example be URLs of a node you set up via [Infura](https://infura.io/).

```bash
ETH_NODE_HTTPS=<place your eth node https address here>
ETH_NODE_WS=<place your eth node websocket address here>
```

## Running

You run the [main.go](cmd/websocketserver/main.go) to start listening for mempool TXs and broadcasting them via
websocket to `localhost:8080/mempooltxs`.

You can then hook up websocat and pipe it to jq by running `websocat ws://localhost:8080/mempooltxs | jq` in your
terminal to verify you get output like:

```json
[
  {
    "type": "0x0",
    "nonce": "0x67fc81",
    "gasPrice": "0x627cc9f2f",
    "maxPriorityFeePerGas": null,
    "maxFeePerGas": null,
    "gas": "0x55730",
    "value": "0x484eb851b606c00",
    "input": "0x",
    "v": "0x25",
    "r": "0x5eedb535667cf7dffdf80a3e0a2bf7db9bc39d70e7558093d79e258e24ace3cd",
    "s": "0x32f8e48146dcb12787f3d369006d7d17ba1d0639977b723e41e073a7a0e6339e",
    "to": "0x9d33eebb0bf685dcb32265e0a7a3a5ee467a9bd9",
    "hash": "0x4909e4c93a35f468d69082930fdc7cfe4251fbe5ad857d33da48f83d52619b21"
  },
  {
    "type": "0x2",
    "nonce": "0x16baa",
    "gasPrice": null,
    "maxPriorityFeePerGas": "0x3b9aca00",
    "maxFeePerGas": "0xa00d967f1",
    "gas": "0x5208",
    "value": "0x7e772d03a84258",
    "input": "0x",
    "v": "0x0",
    "r": "0x3126ae554f96b38c60f8fef6d2bdb9419bad30b6d11aa64345f417882b77b9cb",
    "s": "0x6e4a69af189a61253406cbc37d7f45d351e46f60088abcb8f246a1908c2aa1fc",
    "to": "0x57a548dce32104ee529184d7d5b11a7e56ae37a8",
    "chainId": "0x1"
  }
]
```

