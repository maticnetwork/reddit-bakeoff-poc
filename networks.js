const { mnemonic } = require("./secrets.json");
const HDWalletProvider = require("@truffle/hdwallet-provider");

module.exports = {
  networks: {
    development: {
      protocol: "http",
      host: "localhost",
      port: 8545,
      gas: 5000000,
      gasPrice: 5e9,
      networkId: "*",
    },
    mumbai: {
      provider: () =>
        new HDWalletProvider(mnemonic, `https://rpc-mumbai.matic.today/`),
      networkId: "*",
      gasPrice: 10e9,
    },
    goerli: {
      provider: () =>
        new HDWalletProvider(
          mnemonic,
          `https://goerli.infura.io/v3/5687b932e64441e5a297a0bfba8895cd`
        ),
      networkId: 5,
      gas: 7000000,
      gasPrice: 10e9, // 10 gwei
    },
  },
};
