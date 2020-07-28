const { GSNProvider } = require("@openzeppelin/gsn-provider");
const { signMessage } = require('./sigUtils')
const { network, params, addresses } = require('./config.json')
const { privateKeys } = require('../secrets.json')
const approveFunction = async ({
  from, to, encodedFunctionCall, txFee, gasPrice, gas, nonce, relayerAddress, relayHubAddress,
}) => {
  let response;
  response = signMessage ({
      from, to, encodedFunctionCall, txFee, gasPrice, gas, nonce, relayerAddress, relayHubAddress,
  });
  return response;
};
const Web3  = require('web3')
const gsnProvider = new GSNProvider(network.mumbai, { approveFunction })
const web3 = new Web3 (gsnProvider)
const wallet = web3.eth.accounts.wallet
const accounts = web3.eth.accounts
wallet.add (privateKeys.owner)

const tokenArtifacts = require('../build/contracts/SubredditPoints_v0.json')
const distributionsArtifacts = require('../build/contracts/Distributions_v0.json')
const subscriptionsArtifacts = require('../build/contracts/Subscriptions_v0.json')

const tokenABI = tokenArtifacts.abi
const distributionsABI = distributionsArtifacts.abi 
const subscriptionsABI = subscriptionsArtifacts.abi 

const tokenAddress = addresses.SubredditPoints
const distributionsAddress = addresses.Distributions
const subscriptionsAddress = addresses.Subscriptions

const token = new web3.eth.Contract(tokenABI, tokenAddress)
const subscriptions = new web3.eth.Contract(subscriptionsABI, subscriptionsAddress)
const distributions = new web3.eth.Contract(distributionsABI, distributionsAddress)

module.exports = {
  web3: web3,
  owner: wallet[0],
  gsnApprover: wallet[0], // keeping gsnapprover same as owner for testing
  token: token,
  wallet,
  subscriptions: subscriptions,
  distributions: distributions,
  params: params,
  approveFunction: approveFunction
}
