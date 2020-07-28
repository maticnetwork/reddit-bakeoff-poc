
const { 
  web3,  // accounts
  token, subscriptions, distributions, // initialised contracts
} = require('./helpers/utils')

let currentRound = 0
let user1Karma = 10000

// 3 random accounts - 0 balance
let owner 
let user1 
let user2

async function setAccounts () {
  owner = await web3.eth.personal.newAccount(web3.utils.randomHex(32));
  user1 = await web3.eth.personal.newAccount(web3.utils.randomHex(32));
  user2 = await web3.eth.personal.newAccount(web3.utils.randomHex(32));
  console.log (' 🔀 Random accounts generated')
  console.log (
    ' owner:',
    owner,
    '\n user1',
    user1,
    '\n user2',
    user2 
  )
  console.log (' ⛽️ ETH balances')
  let owner_ = await web3.eth.getBalance(owner)
  let user1_ = await web3.eth.getBalance(owner)
  let user2_ = await web3.eth.getBalance(owner)
  console.log (
    ' owner:', owner_,
    '\n user1: ', user1_,
    '\n user2: ', user2_
  )
}

async function claim () {
  // send claim from owner
  console.log ('\n🚀 sending transaction from owner to mint tokens for user1')
  await distributions.methods.claim (
    currentRound, user1, user1Karma
  ).send({ from: owner }).then((r) => {
    console.log (' claim sent from owner (', owner, ') for ', user1,'\n', ' txhash:', r.transactionHash ,'\n');
  })
}

async function tokenBalances() {
  console.log ('\n updated token balances\n')
  let user1balance = await token.methods.balanceOf(user1).call()

  let user2balance = await token.methods.balanceOf(user2).call()

  console.log (
    ' user1:', user1balance,
    '\n',
    'user2:', user2balance
  )
}

async function send () {
  // transfer some amount from user1 to user 2
  console.log ('\n🚀 transferring from user1 to user2\n')

  let userData = web3.utils.randomHex(32)
  let operatorData  = web3.utils.randomHex(32)

  await token.methods.operatorSend(
    user1, user2, 1, userData, operatorData
  ).send({
    from: user1,
    gas: 500000
  }).then((r) => {
    console.log (' transferred', r.transactionHash)
  })
}

async function subscribe () {
  // perform subscription from user 1
  console.log ('\n🚀 subscribing from user1\n')
  
  let recipient = web3.utils.randomHex(20)
  let renewable = false 

  await subscriptions.methods.subscribe(
    recipient, renewable
  ).send ({
    from: user1,
    gas: 500000
  }).then ((r) => {
    console.log ('subscribed', r.transactionHash)
  })
}

async function main () {
  await setAccounts () 
  await claim () 
  await tokenBalances ()
  await send ()
  await tokenBalances ()
  await subscribe () 
  await tokenBalances ()
}

main () 
