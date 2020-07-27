/**
 * Network: Matic -> Main 
 * Testing: 
 * new user accounts (fund from owner)
 * user1 claims
 * user1 burns tokens on Matic
 * 
 */

// Setup

const { web3, owner, token, subscriptions, distributions, wait, wallet, newAccount, claimSig, karmaSource, params } = require('./helpers/utils')
var fs = require('fs');

const mode = process.env["NETWORK"]
const addresses = require('./helpers/mumbai.json')
const Matic = require('matic').default
const config = require('config')
const rootRPC = config.get(mode + '.parent.rpc')
const childRPC = config.get(mode + '.network.rpc')
const matic = new Matic ({
  maticProvider: childRPC,
  parentProvider: rootRPC,
  rootChain: addresses.Main.Contracts.RootChainProxy,
  withdrawManager: addresses.Main.Contracts.WithdrawManagerProxy,
  depositManager: addresses.Main.Contracts.DepositManagerProxy,
  registry: addresses.Main.Contracts.Registry,
})
let user1
let user1PvtKey
let currentRound = 0
let user1Karma = 100
let tokenAddress = token.options.address
let amount = 50


async function setAndFundUsers () {
  let user1_ = await newAccount ()
  let amount = web3.utils.toWei('1', 'ether')
  wallet.add (user1_)
  user1PvtKey = user1_.privateKey
  await matic.initialize()
  await matic.setWallet(user1_.privateKey)
  user1_details = JSON.stringify(user1_)
  
  user1 = user1_.address
  console.log (
    'accounts:',
    '\n owner=', owner.address,
    '\n (randomly generated):',
    '\n user1=', user1
  )
  console.log ('\nfunding user1 (for gas)')
  await web3.eth.sendTransaction ({
    from: owner,
    to: user1, 
    gas: 500000,
    value: amount
  }).then(console.log ('user1 funded with', web3.utils.fromWei(amount, 'ether'), 'matic token'))
}
async function sendClaim () {

  console.log ('\ngenerating sig for claim (by karmaSource). Karma Source is\n', karmaSource.address)

  let sig = claimSig (currentRound, user1, user1Karma)
  
  let a = await distributions.methods.karmaSource().call()
  console.log ('\n karmasource on contract:', a)

  console.log ('\nsending claim from user1\n')
  await distributions.methods.claim (
    currentRound, user1, user1Karma, sig.signature
  ).send({
    from: user1,
    gas: 500000
  }).then((r) => {
    console.log (
      'claim sent from user1 (', 
      user1, 
      ') with sig ', 
      sig.signature, 
      ' with Karma', user1Karma,
      '\ntxhash:', 
      r.transactionHash ,
      '\n'
    );
  })
}
async function balances () {
  console.log ('\ndisplaying token balances')
  let user1balance = await token.methods.balanceOf(user1).call()
  console.log (
    ' user1:', user1balance
  )
}

async function startWithdraw() {
  await matic.startWithdrawMintableERC20(tokenAddress, amount, {
    from: user1
  })
  .then(async (r) => {
      hash = r.transactionHash
      console.log('initiated withdraw')
      console.log ('transaction hash:', r.transactionHash, '\nblock number:', r.blockNumber)

      let _info = {}
      _info["user"] = user1
      _info["userPvtKey"] = user1PvtKey
      _info["burnTx"] = r.transactionHash
      _info["blockNumber"] = r.blockNumber.toString()

      info = JSON.stringify(_info)
      fs.writeFile('info.json', info, 'utf8', (err, data) => {
        if (err) console.log (err)
      })
  })
  return hash
}


async function main () {
  await setAndFundUsers ()
  await balances ()
  console.log ('===')
  await sendClaim ()
  console.log ('===')
  wait (3000)
  await balances ()
  console.log ('===')
  // withdraw from matic network
  hash = await startWithdraw()
  console.log ('===')
  console.log ('it takes about ~10 minutes for the validators to submit checkpoint. Once the checkpoint with the blocknumber has been included in a header (to check, run: `helpers/check-inclusion.js <block-number>`), run `test2-exit.js <transaction-hash>`')
  console.log ('===')
}

main ()
// balances ()
