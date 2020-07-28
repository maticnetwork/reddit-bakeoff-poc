## Usage

### Networks
|Network|Explorer|RPC|
|---|---|---|
|Matic's Mumbai Test network|https://mumbai-explorer.matic.today/|`https://rpc-mumbai.matic.today`|
|Görli Test network|https://goerli.etherscan.io/|`https://goerli.infura.io/v3/<api-key>`|

### Setup

for Mumbai test network
```bash
$ export NETWORK=mumbai && node scripts/<test-file>
```

## Contract Addresses
(all addresses can be found in `/scripts/config.json`)
### Network: Mumbai 
|Contract|Address (Mumbai)|
|---|---|
|SubredditPoints|[`0x7F4DeA17Cc1a3B65f10E72897b95948F1F1f5A29`](https://mumbai-explorer.matic.today/address/0x7F4DeA17Cc1a3B65f10E72897b95948F1F1f5A29/transactions)|
|Distributions|[`0x7f77947be6Cc1F955C8704fC90f009D65b91D874`](https://mumbai-explorer.matic.today/address/0x7f77947be6Cc1F955C8704fC90f009D65b91D874/transactions)|
|Subscriptions|[`0x46C78e978e6C9415237684dd5bA29d9d82959Fe1`](https://mumbai-explorer.matic.today/address/0x46C78e978e6C9415237684dd5bA29d9d82959Fe1/transactions)|

## Scripts

Following scripts can be tested for a *single user* flow

The contracts require signatures from certain accounts throughout the flow (requesting claims to mint new tokens, swapping exit nft for erc20 on Ethereum chain (while withdrawing, read: [What is ExitNFT?](docs/withdraw.md)))
|user|address|
|---|---|
|gsnApprover|`0xFd71Dc9721d9ddCF0480A582927c3dCd42f3064C`|
|contracts owner|`0xFd71Dc9721d9ddCF0480A582927c3dCd42f3064C`|

1. `init.js`: Test mint, transfer, subscribe
   
   Run:
   ```
   $ node scripts/test1.js
   ```
   To simulate a single user claim, transfer and subscribe on Matic's mumbai test network
2. `gsn-test.js`: Test plasma withdraw
   
   Run:
   ```
   $ node scripts/test2.js
   ```
   To test a single user withdraw from Matic chain (Matic -> Ethereum): Withdraw is a three step process
   