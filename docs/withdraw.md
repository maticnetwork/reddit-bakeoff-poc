# Withdrawal process from Matic

Withdrawal from Matic chain is a three-step process. The user is required to burn their asset on Matic, submit the proof of this burn on Main chain, which is verified against the checkpoint (submitted by the validators) including the block data consisting of the transaction. Once the verification is sucessful, the user receives an ExitNFT, which is an NFT representative of the asset withdrawn from Matic. The user can either choose to wait for the Challenge Exit Period (7 days on Mainnet; ~5 minutes on Testnet) or swap their ExitNFT for their token (paying a certain fee) from a liquidity maintainer contract. 

## What is Challenge Period?

We want to make sure that the output being referenced is actually unspent, so we introduce a challenge period. Basically, a challenge period is a period of time in which people can challenge the validity of the exit by proving that the UTXO is actually spent. Users can prove a UTXO is spent by revealing another transaction that spends the UTXO signed by the user who started the exit.

## What is an ExitNFT?

ExitNFT is an NFT that is minted for the user on confirmation of their withdraw on Ethereum chain. This NFT is representative of the amount and asset they're trying to withdraw. Using this ExitNFT the user can either choose to wait for the Challenge Exit Period and get their claimed assets on main chain or get the exitNFT swapped from a liquidity provider on the main chain. 

## The two types of Exits

The choice between the two types of Exits for the user comes in only after the user has confirmed their withdrawal process on Main chain (the second step of withdraw). Once the user is the owner of their ExitNFT they can either:
- swap it for ERC20 instantly from the liquidity provider contract
- wait for the Challenge Exit Period and then calling process exits to receive their tokens back (and burning the ExitNFT).

You can read more about Faster Exits [here](https://medium.com/matic-network/enabling-faster-plasma-exits-3ca5a936f215) and [here](https://blog.matic.network/faster-plasma-exits-with-nuo/).
