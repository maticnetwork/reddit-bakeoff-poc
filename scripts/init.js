
const { 
  owner, gsnApprover, // accounts
  token, subscriptions, distributions, // contracts
  params,  // contract params
} = require('./utils')

const tokenAddress = token.options.address
const distributionsAddress = distributions.options.address
const subscriptionsAddress = subscriptions.options.address 

async function initialiseContracts () {
  console.log ('\n initialising contracts ... ✨ \n')

  await token.methods.initialize(
    owner.address, distributionsAddress, params.subredditName, params.subredditTokenName, params.subredditTokenSymbol, [], gsnApprover.address
  ).send({
    from: owner.address,
    gas: 5000000
  }).then((r) => {
    console.log(' SubredditPoints: initialised, ', r.transactionHash)
  })
  
  await distributions.methods.initialize(
    owner.address, tokenAddress, params.initialSupply, params.initialKarma, gsnApprover.address
  ).send({
    from: owner.address,
    gas: 500000
  }).then((r) => {
    console.log(' Distributions: initialised, ', r.transactionHash)
  })

  await subscriptions.methods.initialize(
    owner.address, tokenAddress, params.subscriptionPrice, params.duration, params.renewBefore, gsnApprover.address
  ).send({
    from: owner.address,
    gas: 500000
  }).then((r) => {
    console.log(' Subscriptions: initialised, ', r.transactionHash)
  })

  

  await token.methods.addDefaultOperator(
    subscriptionsAddress
  ).send({
    from: owner.address,
    gas: 5000000
  }).then((r) => {
    console.log(' 📌 SubredditPoints: added subscription contract as default operator ', r.transactionHash)
  })

}


initialiseContracts()
