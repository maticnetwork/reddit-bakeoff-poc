const mode = process.env["NETWORK"] // 'dev' or 'mumbai'
const config = require('config');
const network = config.get(mode + '.network')


const Web3 = require('web3')
const web3 = new Web3(network.rpc)

const networkId = async () => await web3.eth.net.getId();

networkId().then(console.log)
