const {
  deployRelayHub,
  fundRecipient,
  registerRelay
} = require('@openzeppelin/gsn-helpers')
const HDWalletProvider = require("truffle-hdwallet-provider")
const { network, addresses } = require('../config.json')
const { mnemonic } = require('../../secrets.json')
const Web3 = require('web3')

const provider = new HDWalletProvider(mnemonic, network.mumbai)
const web3 = new Web3(provider)

async function init () {
  // let r = await deployRelayHub(web3);
  // console.log (r)
  await registerRelay(web3, { 
    relayUrl: 'http://localhost:8090'
  });
  console.log ('completed registration')
  for (contract in addresses) {
    let r = await fundRecipient(web3, {
      recipient: addresses[contract],
      amount: web3.utils.toWei('1', 'ether')
    })
    console.log ('funded contract at',addresses[contract], r)
  }
}

init ()
