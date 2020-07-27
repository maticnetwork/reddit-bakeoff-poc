const { web3, token, subscriptions, distributions, params, owner, karmaSource } = require ('./helpers/utils')

// initialise contracts

const tokenAddress = token.options.address 
const subscriptionsAddress = subscriptions.options.address 
const distributionsAddress = distributions.options.address 

async function initializeContracts () {

  await token.methods.initializecontract (
    owner.address, distributionsAddress, params.subredditName, params.subredditTokenName, params.subredditTokenSymbol, []
  ).send ({
    from: owner.address, 
    gas: 500000
  }).then ((r) => {
    console.log('token initialised', r.transactionHash)
  })

  await token.methods.addDefaultOperator(
    subscriptionsAddress
  ).send ({
    from: owner.address,
    gas: 500000
  }).then ((r) => {
    console.log ('default operator added', r.transactionHash)
  })

  await distributions.methods.initializecontract (
    owner.address, tokenAddress, params.initialSupply, params.initialKarma, karmaSource.address
  ).send ({
    from: owner.address,
    gas: 500000
  }).then ( (r) => {
    console.log ('distributions initialised', r.transactionHash)
  })

  await subscriptions.methods.initializecontract (
    owner.address, tokenAddress, params.subscriptionPrice, params.duration, params.renewBefore
  ).send ({
    from: owner.address, 
    gas: 500000
  }).then ((r) => {
    console.log ('subscriptions initialised', r.transactionHash)
  })

  

}

initializeContracts()
