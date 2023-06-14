# Hedera Development Sandbox

A dynamic playground atop [Hedera Testnet](https://docs.hedera.com/hedera/networks/testnet), exhibiting seamless interaction via [hedera-sdk-go@v2](https://github.com/hashgraph/hedera-sdk-go).

# Getting Started

## Requirement

- [go](https://go.dev)
- [git](https://git-scm.com/)
- [Hedera Testnet Access](https://docs.hedera.com/hedera/getting-started/introduction)

## Quick start

```
git clone https://github.com/quiet-node/hedera-dev-sandbox.git
cd hedera-dev-sandbox
```

## Running the project

#### 1. Set environment variables

- create a `.env` file using the `.example.env` as the template and fill out the variables.

```
HED_ACCOUNT_ID=DER Encoded Private Key
HED_PRIVATE_KEY=Account ID
```

#### 2. Run project (dev mode)

```
make dev
```

## Resources:

- https://docs.hedera.com/hedera/getting-started
