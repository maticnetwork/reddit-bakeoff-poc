/**
 * Network: Mainchain
 * Plasma exit with the token
 * performs confirm withdraw
 */
const info = require('../info.json')
const fs = require('fs')
const userPvtKey = info.userPvtKey
const user = info.user
const hash = info.burnTx

const { web3, owner, wait } = require('./helpers/utils')


const config = require('config')
const mode = process.env["NETWORK"]
const addresses = require('./helpers/mumbai.json')
const rootRPC = config.get(mode + '.parent.rpc')
const childRPC = config.get(mode + '.network.rpc')
const Matic = require('matic').default
const matic = new Matic ({
  maticProvider: childRPC,
  parentProvider: rootRPC,
  rootChain: addresses.Main.Contracts.RootChainProxy,
  withdrawManager: addresses.Main.Contracts.WithdrawManagerProxy,
  depositManager: addresses.Main.Contracts.DepositManagerProxy,
  registry: addresses.Main.Contracts.Registry,
})
let predicate = addresses.Main.Contracts.Predicates.MintableERC20Predicate

let root = addresses.Main.Contracts.Tokens.TestSubredditToken
let abi = require('../build/contracts/SubredditPointsParent.json').abi 

const exitnftAbi = require('../build/contracts/ExitNFT.json').abi
const exitnftAddress = addresses.Main.Contracts.ExitNFT

web3.setProvider(rootRPC)
const exitnft = new web3.eth.Contract (exitnftAbi, exitnftAddress)
const rootToken = new web3.eth.Contract(abi, root)

async function balances () {
  // display ownership of exit nft
  let a = await exitnft.methods.balanceOf(user).call()
  console.log ('balance of exitnft:',a)
  // display balance of token on root 
  let b = await rootToken.methods.balanceOf(user).call()
  console.log ('balance of subreddit token on root:', b)
  // let c = await web3.eth.getBalance(user)
  // console.log ('user eth balance:', web3.utils.fromWei(c, "ether"), 'ether')
}

async function fundUser () {
  web3.setProvider(rootRPC)
  console.log('funding user (to cover gas)')
  let a = await web3.eth.sendTransaction({
    from: owner,
    to: user,
    value: web3.utils.toWei("0.01", "ether"),
    gas: 500000
  }).then((r) => {
    console.log ('funded user.')
  })
}


async function main () {
  await matic.initialize()
  await matic.setWallet(userPvtKey)
  await fundUser ()
  console.log('---\nbalances\n')
  await balances ()
  console.log('\n---')
  console.log ('\nstarting withdraw\n---')
  a = await matic.withdrawMintableERC20Token(hash, predicate, { from: user, gasPrice: 10000000000 })
  fs.readFile('info.json', function (err, data) {
    let obj = JSON.parse(data)
    obj["confirmTx"] = a.transactionHash.toString()
    fs.writeFile('info.json', JSON.stringify(obj), () => {})
  })
  console.log ('Withdraw Confirmed on root chain.\n', 'txhash:', a.transactionHash)
  await balances()
  wait (3000)
  console.log('---\nbalances\n')
  await balances()
  console.log('\n---')
  console.log ('\nrun `test3-fast-exit.js` for testing faster exits. run `test3-plasma-exit.js` for plasma exits.')
}

main ()
