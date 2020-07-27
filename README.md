# Benchmark Testing

### Requirements

- Go version 1.13 or above


### Configuration for Benchmark
1. Add Benchmark config in `config/mumbai/reddit/reddit-benchmark.json` file,

```
{
    "NumClaims": 100000,
    "NumTransfers": 100000,
    "NumSubscribe": 25000
}
```

2. Add reddit Contract addresses in `config/mumbai/reddit/redditContracts.json` file,
```
{
      "SubredditPoints_v0": "0xC7B0235b89E097fE12652a8AD673f7390A90d43f",
      "SubredditPointsParent": "0xe356ECDa1d8F46733D555d6Dc0A6f3499CA7279g",
      "Subscriptions_v0": "0xD1A3290d701e2eabA0826c2475CD6693238b189x",
      "Distributions_v0": "0x3C86Dd45faC31990731049520c1A4C915648669d"
}
```

3. Add Infura key in `config/mumbai/network-config.json` file
```
    "eth_rpc_url" : "https://goerli.infura.io/v3/<INFURA_KEY>"
```


4. Add reddit contract owner details in `config/mumbai/reddit/reddit-accounts.json` file
```
    [{"address": "", "priv_key": "","pub_key": "0x"}]
```

5. Add faucet account details in `config/mumbai/matic-faucet.json` file
```
    {"address": "", "priv_key": "","pub_key": "0x"}
```



### Run Benchmark
```

   - Run on mumbai test network of matic
    Initialize(If not initialized already) - go test  -timeout 20m -network mumbai -bench=Benchmark/Reddit/Initialize --benchtime 1ns
    Claim -     go test -timeout 40m -network mumbai -bench=Benchmark/Reddit/Claim --benchtime 1ns > claims.txt
    Subcribe -  go test -timeout 40m -network mumbai -bench=Benchmark/Reddit/Subscribe --benchtime 1ns > subscribe.txt
    Transfer -  go test -timeout 40m -network mumbai -bench=Benchmark/Reddit/Transfer --benchtime 1ns > transfer.txt


Note... Change timeout flag incase of more time required.

Logs... You can find the logs in respective txt files.(claims.txt, transfer.txt,subscribe.txt )
```
