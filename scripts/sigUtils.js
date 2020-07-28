const { network } = require('./config.json')
const Web3 = require('web3')
const localWeb3 = new Web3 ( network.mumbai )
const { privateKeys } = require('../secrets.json')
const { utils } = require("@openzeppelin/gsn-provider");

localWeb3.eth.accounts.wallet.add(privateKeys.gsnApprover)
const signMessage = async data => {
  // keeping gsnapprover same as owner for testing
  const hash = localWeb3.utils.soliditySha3(
    data.from,
    data.to,
    data.encodedFunctionCall,
    data.txFee,
    data.gasPrice,
    data.gas,
    data.nonce,
    data.relayerAddress,
    data.relayHubAddress
  );
  const signature = await localWeb3.eth.sign(hash, localWeb3.eth.accounts.wallet[0].address);
  return utils.fixSignature(signature);
};


// testing 
// signMessage({
//   from: localWeb3.utils.randomHex(20), 
//   to: localWeb3.utils.randomHex(20), 
//   encodedFunctionCall: localWeb3.utils.randomHex(4), 
//   txFee: 4, 
//   gasPrice: 1e6, 
//   gas: 1e6, 
//   nonce: 21, 
//   relayerAddress: localWeb3.utils.randomHex(20), 
//   relayHubAddress: localWeb3.utils.randomHex(20),
// }).then(console.log)
// console.log (signMessage)

module.exports = { signMessage }
