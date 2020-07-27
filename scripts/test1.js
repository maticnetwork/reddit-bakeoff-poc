/**
 * 
 * Network: Matic 
 * Testing: 
 * user1 sends claim
 * user1 transfers funds to user2
 * user1 subscribes (burning tokens)
 * 
 */


const { web3, owner, token, subscriptions, distributions, wait, wallet, newAccount, claimSig, karmaSource, params } = require('./helpers/utils')

let currentRound = 0
let user1Karma = 20

let user1, user2
async function setAndFundUsers () {
  let user1_ = await newAccount ()
  let user2_ = await newAccount ()

  let amount = web3.utils.toWei('1', 'ether')

  wallet.add (user1_)
  wallet.add (user2_)

  user1 = user1_.address 
  user2 = user2_.address

  console.log (
    'accounts:',
    '\n owner=', owner.address,
    '\n (randomly generated):',
    '\n user1=', user1,
    '\n user2=', user2
  )

  console.log ('\nfunding user1 and user2 (for gas)')
  
  await web3.eth.sendTransaction ({
    from: owner,
    to: user1, 
    gas: 500000,
    value: amount
  }).then(console.log ('user1 funded with', web3.utils.fromWei(amount, 'ether'), 'matic token'))
  await web3.eth.sendTransaction ({
    from: owner,
    to: user2, 
    gas: 500000,
    value: amount
  }).then(console.log ('user2 funded with', web3.utils.fromWei(amount, 'ether'), 'matic token'))

  user1bal = await web3.eth.getBalance(user1)
  user2bal = await web3.eth.getBalance(user2)

}

async function info () {
  console.log ('\ndisplaying contract info\n')

  const supply = await distributions.methods.initialSupply().call();
  const subreddit = await distributions.methods.subreddit().call()
  console.log (
  ' Round:', currentRound,
  '\n',
  'initial supply:', supply, 'tokens',
  '\n',
  'initial karma:', initialKarma,
  '\n',
  'subreddit:', subreddit
  )
  console.log (
    '\n Subscriptions',
    '\n',
    'duration:', params.duration, 'seconds', 
    '\n',
    'price:', params.subscriptionPrice, 'tokens',
    '\n',
    'renew before:', params.renewBefore, 'seconds'
  )
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

// sendClaim ()

async function balances () {
  console.log ('\ndisplaying token balances\n')
  let user1balance = await token.methods.balanceOf(user1).call()

  const user2balance = await token.methods.balanceOf(user2).call()

  console.log (
    ' user1:', user1balance,
    '\n',
    'user2:', user2balance
  )
}

async function send () {
  console.log ('\ntransferring from user1 to user2\n')

  let userData = web3.utils.randomHex (32)
  let operatorData  = web3.utils.randomHex (32)
  let amount = 1

  await token.methods.operatorSend (
    user1, user2, amount, userData, operatorData
  ).send({
    from: user1,
    gas: 500000
  }).then((r) => {
    console.log (' transferred', r.transactionHash)
  })
}

async function subscribe () {
  console.log ('\nsubscribing from user1\n')
  
  let recipient = web3.utils.randomHex(20)
  let renewable = false 

  let price = await subscriptions.methods.price().call()

  await subscriptions.methods.subscribe(
    recipient, renewable
  ).send ({
    from: user1,
    gas: 500000
  }).then ((r) => {
    console.log (
      'subscribed, at price:',
      price,
      'txhash:',
      r.transactionHash
    )
  })
}

async function main () {
  console.log ('Test: Create new users, request claim from user1, transfer user1 -> user2, subscribe (burn) from user1')
  
  await setAndFundUsers ()
  await info ()
  await balances ()
  await sendClaim ()
  await wait (3000)
  await balances () 
  await send () 
  await wait (2000)
  await balances ()
  await subscribe ()
  await wait (2000)
  await balances () 
}

main ()
