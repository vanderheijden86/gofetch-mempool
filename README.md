# Ethereum Mempool Explorer

Golang implementation of an ethereum mempool explorer. For now it's mainly a way to gain experience with accessing
the ethereum mempool (pending transactions), refresh my understanding of the go programming language
and work with the go-ethereum `ethclient` and `gethclient` packages.

## Setup

In the project's `root folder` directory, you need to place a file called `.
env` in which to put your Infura key for ethereum mainnet. Easiest is to rename the `example.env` file to `.env` 
and paste in your own key.

```bash
INFURA_KEY=<place your INFURA key here>
```

## Running

You run the tests in `datacollection_test.go` to see the functionality to get the mempool TX hashes and transaction
details.

```text
=== RUN   TestMain_streamMemPoolTxs
Channel length:  0
Channel capacity:  100
0x58b232edfac1325a15df16da8a5c7b053f8767995fa592a58ebc4a9c2739211c
Channel length:  0
Channel capacity:  100
0x5430a13353856c8324247af58dc552316e8f5638d34a8e1a508b8abfc396b9ea
Channel length:  0
Channel capacity:  100
0x330e7c22df45f524d4a76ce18bc38bb00f21b4e8b0e82aa1f32e65eff742b99a
Channel length:  0
Channel capacity:  100
0xbecb2757392ba5d4488135324ffecbc7f4c0e677b3ff94479bb4a4fc140188c4
Channel length:  0
```

