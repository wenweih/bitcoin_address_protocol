# bitcoin_address_protocol
generate bitcoin address protocol by raw way

code for article https://www.jianshu.com/p/71a4454c74da

assume that there is Golang env for you machine.
1. go build
2. ./bitcoin_protocol

```shell
➜  bitcoin_protocol git:(master) ./bitcoin_protocol
0 - Having a private ECDSA key
5C992BE8E980E508402B62AE19B272C7089197445653C94D28B02ABAD47F9B0C
=======================
1 - Take the corresponding public key generated with it (65 bytes, 1 byte 0x04, 32 bytes corresponding to X coordinate, 32 bytes corresponding to Y coordinate)
raw public key B53DD83B65F131B45B44B04B5D45CA095AF1CEB0918C8AE024A414B93BC17811A77451344682B1E12561304943235FB887F8216D047372F6588539C3A57321B2
=======================
2 - Perform SHA-256 hashing on the public key
2EE04B5A2BF7CBBCB82D6B2334066CF91EF56B199CB48D803D138C23B43FBACD
=======================
3 - Perform RIPEMD-160 hashing on the result of SHA-256
51337C281FADA1EDAE5EB3AFD827FBE56031755C
=======================
4 - Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)
0051337C281FADA1EDAE5EB3AFD827FBE56031755C
=======================
5 - Perform SHA-256 hash on the extended RIPEMD-160 result
23CF88427B84E39760B9BD04D78F046115729C5AFF8578F7AF451DD3E86A8FDC
=======================
6 - Perform SHA-256 hash on the result of the previous SHA-256 hash
66A6F89BAD3E3C29162BD8E9C1A4513EFDCB3AAC305BF48A70D287001FC9117A
=======================
7 - Take the first 4 bytes of the second SHA-256 hash. This is the address checksum
66A6F89B
=======================
8 - Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from stage 4. This is the 25-byte binary Bitcoin Address.
0051337C281FADA1EDAE5EB3AFD827FBE56031755C66A6F89B
=======================
9 - Convert the result from a byte string into a base58 string using Base58Check encoding. This is the most commonly used Bitcoin Address format
18QMQn7poD737SBC7mqaQYs6bQHoe6mMqU
=======================
```

see the address in blockchain browser: [18QMQn7poD737SBC7mqaQYs6bQHoe6mMqU](https://blockchain.info/address/18QMQn7poD737SBC7mqaQYs6bQHoe6mMqU)
