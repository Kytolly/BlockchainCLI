Sure, here's a more detailed README in English for your Blockchain CLI project:

---

# Blockchain CLI (UESTC-Go语言和区块链技术-期末设计)

## Introduction

This project is a basic implementation of a blockchain system based on the Bitcoin protocol, encompassing all aspects of the final design project. Built upon the foundational code discussed in classes on "Blockchain Basic Prototype" and "Blockchain Proof of Work," this project further enhances the blockchain's underlying data structure by incorporating a Merkle Tree structure for transaction data.

Additionally, leveraging the cryptographic concepts covered in "Blockchain and Cryptography," this project utilizes public key cryptography for designing blockchain accounts and transaction structures. Transactions are signed with private keys and sent to blockchain nodes, which verify the signatures using public keys on the blockchain before recording them in a block.

The project also includes functionality to query transactions on the blockchain and utilize Merkle Trees to verify if a specific transaction exists within a block.

## Features

- [x] Mining and Proof of Work 
- [x] Blockchain Iteration 
- [x] Persistence and Indexing
- [x] Merkle Tree for Accelerated Querying
- [x] Digital Signature and Verification
- [x] Network Communication and Nodes
- [x] Command Line Interface (CLI)
- [x] Engineering: Log and Config

## Getting Started
you can `go build` it, and run `blockchain help` for more details.


## License

This project is licensed under the MIT License. This means you are free to use, modify, distribute, and incorporate the source code into your projects as long as the original copyright and license notice are included.

**Warning:** Do not cheat on your homework! (This is intended specifically for UESTC students.)

## Contributing

Maybe there are some bugs still to be fixed. If you want to contribute to it, you can fork free and dont forget to issue.

## References

- [blockchain_go](https://github.com/Jeiwan/blockchain_go) - An educational blockchain project in Go.
- [Bitcoin Protocol Documentation](https://en.bitcoin.it/wiki/Protocol_documentation) - Comprehensive documentation on the Bitcoin protocol.