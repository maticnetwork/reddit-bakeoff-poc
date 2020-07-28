## Usage

### Networks
|Network|Explorer|RPC|
|---|---|---|
|Matic's Mumbai Test network|https://mumbai-explorer.matic.today/|`https://rpc-mumbai.matic.today`|

### Setup

for local testing

spin up Ganache
```bash
# new terminal instance
$ ganache-cli
```

run relayer
```bash
# new terminal instance
$ npx oz-gsn run-relayer
```

deploy contracts
```bash
# new terminal instance
$ oz compile

# run deployment for all three contracts
$ oz deploy
```

fund contracts
```bash
# new terminal instance

# run the following command for all three contract addresses
$ npx oz-gsn fund-recipient --recipient <contract-address>
```

## Scripts

Following scripts can be tested for a *single user* flow

The contracts require private key for an owner account (with enough funds) to initialise the contracts

1. `init.js`: Test mint, transfer, subscribe
   
   Run:
   ```
   $ node scripts/init.js
   ```
   To initialise the contracts on the chain 
2. `gsn-test.js`: Test plasma withdraw
   
   Run:
   ```
   $ node scripts/gsn-test.js
   ```
   To test a single user flow with GSN. (submit claim -> receive tokens, transfer some tokens, subscribe -> burn tokens)