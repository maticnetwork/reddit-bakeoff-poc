/**
 * Network: Main network
 * Fast exit
 * this test is dependent on test2.js and test2-exit.js
 * picking data from ../info.json
 */

const { web3, wallet, owner, wait, fastexit, fastExitSig } = require('./helpers/utils')
const { user, userPvtKey, confirmTx } = require('../info.json')
const config = require('config')
const mode = process.env["NETWORK"]
const addresses = require('./helpers/mumbai.json')
const rootRPC = config.get(mode + '.parent.rpc')

const root = addresses.Main.Contracts.Tokens.TestSubredditToken
let abi = require('../build/contracts/SubredditPointsParent.json').abi 
const rootToken = new web3.eth.Contract(abi, root)

web3.setProvider(rootRPC)
wallet.add(userPvtKey)
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

  if (a > 0) {
    // get token Id if balance of exit nft > 0
    let receipt = await web3.eth.getTransactionReceipt(confirmTx)
    // console.log (receipt)
    let events = await exitnft.getPastEvents ("Transfer", {
      filter: { to: user },
      fromBlock: receipt.blockNumber,
      toBlock: receipt.blockNumber
    })
    // console.log (events)
    for (let i = 0; i < events.length; i ++) {
      if (events[i].transactionHash == confirmTx) {
        tokenId = events[i].returnValues.tokenId
      }
    }
  }
  console.log (tokenId)
  let d = await exitnft.methods.ownerOf(tokenId).call()
  console.log ('owner of', tokenId, ':',d)

  console.log ('exitNFT tokenId:', tokenId)
  console.log ('---\n')
}

async function info () {
  let interestRatePerCent = await fastexit.methods.interestRatePerCent().call()
  let liquidityProvider = await fastexit.methods.liquidityProvider().call()
  let liquidity = await fastexit.methods.liquidity(rootToken.options.address).call()
  console.log ('\n---info---\ninterest rate percent:', interestRatePerCent,'\nliquidity provider:', liquidityProvider,'\nliquidity:',liquidity,'\n---\n')
}

async function swap () {
  console.log ('\n---initiating swap---\n')
  let approve = await exitnft.methods.approve(
    fastexit.options.address, tokenId
  ).send({
    from: wallet[user].address,
    gas: 500000
  })
  console.log ('approved fastexit contract to swap', approve.transactionHash)
  // requesting sig from liquidity provider
  let sig = fastExitSig(tokenId, user)
  console.log ('sig:', sig)

  let swap = await fastexit.methods.swap(
    tokenId, sig
  ).send({
    from: wallet[user].address,
    gas: 500000,
    gasPrice: 10000000000
  })

  console.log ('swapped exitNft for erc20. txhash:', swap.transactionHash)
  console.log ('---')
}

async function fastExit () {
  // would need to fund user with enough eth to perform the swap before the test
  await balances()
  await info ()
  await swap()
  await wait (3000)
  await balances()
}

fastExit ()
