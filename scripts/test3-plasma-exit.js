/**
 * Network: Main network
 * Perform process exit here 
 * Note: process exits can be called by any account with enough funds
 * Plasma safety exit 
 * this test is dependent on test2.js and test2-exit.js
 * picking data from ../info.json
 */

const { web3, owner, wait } = require('./helpers/utils')
const info = require('../info.json')
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
const root = addresses.Main.Contracts.Tokens.TestSubredditToken
let abi = require('../build/contracts/SubredditPointsParent.json').abi 
const rootToken = new web3.eth.Contract(abi, root)
const user = info.user

web3.setProvider(rootRPC)

const exitnftAbi = require('../build/contracts/ExitNFT.json').abi
const exitnftAddress = addresses.Main.Contracts.ExitNFT
const exitnft = new web3.eth.Contract (exitnftAbi, exitnftAddress)


async function balances () {
  // display ownership of exit nft
  console.log ('\n---balances---')
  let a = await exitnft.methods.balanceOf(user).call()
  console.log ('balance of exitnft:',a)
  // display balance of token on root 
  let b = await rootToken.methods.balanceOf(user).call()
  console.log ('balance of subreddit token on root:', b)
  let c = await web3.eth.getBalance(user)
  console.log ('user eth balance:', web3.utils.fromWei(c, "ether"), 'ether')
  console.log ('---\n')
}

async function processExit () {
  await matic.initialize()
  await matic.setWallet(owner.privateKey)
  await balances()
  await matic.processExits(root, {
      from: owner.address,
  }).then((r) => {
    console.log ('completed process exits.\n', 'txhash:', r.transactionHash)
  })
  await wait (3000)
  await balances()
}

processExit ()
