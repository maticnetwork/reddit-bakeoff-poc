// process.env["NODE_CONFIG_DIR"] = "../config/";
const mode = process.env["NETWORK"]
const config = require('config')
let ethUtils = require("ethereumjs-util");

const network = config.get(mode + '.network')
const params = { 
  initialSupply, 
  initialKarma,
  subscriptionPrice,
  duration, 
  renewBefore,
  subredditName,
  subredditTokenSymbol,
  subredditTokenName
} = config.get(mode + '.params')

const privateKeys = config.get(mode + '.privateKeys')

const Web3 = require('web3')
const web3 = new Web3 (network.rpc)
const wallet = web3.eth.accounts.wallet

const tokenArtifacts = require('../../build/contracts/SubredditPoints_v0.json')
const distributionsArtifacts = require('../../build/contracts/Distributions_v0.json')
const subscriptionsArtifacts = require('../../build/contracts/Subscriptions_v0.json')
const fastExitAbi = require('./fastExit.json')
let tokenAddress, distributionsAddress, subscriptionsAddress, fastExitAddress

// set deployed to true to fetch deployed and verified contract addresses on the two networks 
// if testing - set deployed to false.

const deployed = config.get(mode + '.deployed')
if (!deployed) {
  tokenAddress = tokenArtifacts.networks[config.get(mode + '.network.netId')].address
  distributionsAddress = distributionsArtifacts.networks[config.get(mode + '.network.netId')].address
  subscriptionsAddress = subscriptionsArtifacts.networks[config.get(mode + '.network.netId')].address
} else {
  tokenAddress = require('./mumbai.json').Matic.Contracts.Tokens.TestSubredditToken
  distributionsAddress = require('./mumbai.json').Matic.Contracts.Tokens.TestDistributions
  subscriptionsAddress = require('./mumbai.json').Matic.Contracts.Tokens.TestSubscriptions
}

fastExitAddress = require('./mumbai.json').Main.Contracts.FastExit
const token = new web3.eth.Contract(tokenArtifacts.abi, tokenAddress)
const subscriptions = new web3.eth.Contract(subscriptionsArtifacts.abi, subscriptionsAddress)
const distributions = new web3.eth.Contract(distributionsArtifacts.abi, distributionsAddress)
const fastexit = new web3.eth.Contract(fastExitAbi, fastExitAddress)

wallet.add(privateKeys.owner)

async function getNewAccount () {
  return await web3.eth.accounts.create();
}

function wait (delayms) {
  return new Promise(function(resolve, reject) {
    setTimeout(resolve, delayms)
  })
}


function claimSig (round, account, karma) {
  let name = params.subredditName.toLowerCase()

  let data = web3.utils.keccak256(web3.eth.abi.encodeParameters(['string', 'uint256', 'uint256', 'uint256'], [name, round, account, karma]))

  let obj = web3.eth.accounts.sign(data, karmaSource.privateKey)

  return obj
}


function fastExitSig (tokenId, user) {
  let dataHash = web3.utils.soliditySha3(
    tokenId,
    web3.utils.toChecksumAddress(user)
  );
  
  let privateKeyString = wallet[0].privateKey.substring(2)
  let vrs = ethUtils.ecsign(
    ethUtils.toBuffer(dataHash),
    // privateKeyString
    Uint8Array.from(Buffer.from(privateKeyString, 'hex'))
  );
  // console.log (vrs)

  return ethUtils.toRpcSig(vrs.v, vrs.r, vrs.s);
}


// set owner as karmaSource
const karmaSource = wallet[0]
// owner is the liquidity provider for the currently deployed contract on goerli
const liquidityProvider = wallet[0]

module.exports = {
  web3, 
  owner: wallet[0],
  token, 
  subscriptions,
  distributions,
  params,
  wait,
  newAccount: getNewAccount,
  wallet, 
  claimSig,
  fastexit,
  karmaSource,
  liquidityProvider,
  fastExitSig
}
