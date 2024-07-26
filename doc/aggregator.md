# Instruction for AVS deployment and running aggregator node

Holesky tesnet Anvil fork is used as an example (Anvil dev accounts are being used).

Create deployment configuration file by specifiying addresses of core Eigenlayer contracts an addresses, eligible to perform particular actions within AVS (task generation, signature aggregation, etc.) See the following [example](../contracts/script/input/deployment_parameters_holesky.json).

Run deployment script
```
cd contracts/

forge script \
    script/LambadaCoprocessorDeployerHoleksy.s.sol:LambadaCoprocessorDeployerHolesky \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --broadcast -v

cd ../
```

Create aggregator configuration file. See the following [example](../tests/nodes/aggregator/config-devnet.yaml).

Ensure aggregator has enough ETH to send transactions
```
cast send --private-key $FUNDER_PK --value 20ether 0xa0Ee7A142d267C1f36714E4a8F75612F20a79720`
```

Run aggregator
 ```
 go run aggregator/cmd/main.go \
    --private-key 0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6
    --config $AGGREGATOR_CONFIG
 ```