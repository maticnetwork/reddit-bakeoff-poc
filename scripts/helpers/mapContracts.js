// map predicate on registry 
// manual map subreddit points on root and child 

// set network and root chain owner account's private key in environment variable
const mode = process.env["NETWORK"]
const ownerPvtKey = process.env["PARENT_OWNER"]


const { web3, wallet, token } = require('./utils')
const config = require('config')
const owner = '0x907f2e1F4A477319A700fC9a28374BA47527050e'

const rootRPC = config.get(mode + '.parent.rpc')
const childRPC = config.get(mode + '.network.rpc')

// Contract ABIs
const governanceAbi = require('../../build/contracts/Governance.json').abi
const registryAbi = require('../../build/contracts/Registry.json').abi
const childchainAbi = require('../../build/contracts/ChildChain.json').abi

// Contract Addresses
const governanceAddress = require('./mumbai.json').Main.Contracts.GovernanceProxy
const registryAddress = require('./mumbai.json').Main.Contracts.Registry
const childchainAddress = require('./mumbai.json').Matic.Contracts.ChildChain

// token addresses
const root = require('./mumbai.json').Main.Contracts.Tokens.TestSubredditToken
const child = token.options.address

// predicate address (root chain)
const predicate = require('./mumbai.json').Main.Contracts.Predicates.MintableERC20Predicate

// add owner (governance and child chain) to wallet
wallet.add(ownerPvtKey)

// Contracts 
const governance = new web3.eth.Contract (
  governanceAbi, governanceAddress
)
const registry = new web3.eth.Contract (
  registryAbi, registryAddress
)
const childchain = new web3.eth.Contract (
  childchainAbi, childchainAddress
)

async function mapPredicate (predicate) {
  await web3.setProvider(rootRPC)
  let r = await registry.methods.addPredicate (
    predicate, 3 /**custom predicate enum type */
  ).encodeABI()
  let tx = await governance.methods.update (
    registry.options.address, r
  ).send({
    from: wallet[owner].address,
    gas: 500000
  })
  console.log ('Added Custom Predicate.')
}
async function checkPredicate(predicate) {
  await web3.setProvider(rootRPC)
  let r = await registry.methods.predicates(
    predicate
  ).call()
  console.log ('\nPredicate: ', predicate, '\nType: ', r)
}
async function mapTokens (root, child) {
  let isNFT = false // only erc20s

  web3.setProvider(rootRPC)
  let a = await registry.methods.mapToken(
    root, child, isNFT
  ).encodeABI()
  console.log (a)
  let tx = await governance.methods.update (
    registry.options.address, a
  ).send({
    from: wallet[owner].address,
    gas: 900000
  })
  console.log ('Mapped on Root (registry.sol)', tx.transactionHash)

  web3.setProvider(childRPC)
  let b = await childchain.methods.mapToken(
    root, child, isNFT
  ).send ({
    from: wallet[owner].address,
    gas: 900000
  })
  console.log ('Mapped on child (childchain.sol)', b.transactionHash)
  console.log ('root: ',root, '\nchild:', child)
}
async function checkTokenMapping (root, child) { 
  console.log ('\nchecking mapping on Registry.sol')

  web3.setProvider(rootRPC)
  let a = await registry.methods.rootToChildToken(
    root
  ).call()
  console.log('root:', root, '\nchild:', a)
  console.log ('\nchecking mapping on Childchain.sol')

  web3.setProvider(childRPC)
  let b = await childchain.methods.tokens(
    root
  ).call()
  console.log('root:', root, '\nchild:', b)
}

async function main () {
  // await mapPredicate(predicate)
  await checkPredicate (predicate)
  // await mapTokens (root, child)
  await checkTokenMapping (root, child)
}

main ()
