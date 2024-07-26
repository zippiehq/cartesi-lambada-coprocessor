# Instructions for running operator node

Current AVS deployment on Holesky testnet is used as example.

Generate operator directory by providing addresses of deployed AVS contracts
```
go ./cli generate-operator --operator-dir $OPERATOR_DIR_PATH
```
Above command will print randomly generated address, its private key and passwords for BLS and ECDSA keystores. Store this information in safe place.

Ensure operator address has enough ETH to send transactions
```
cast send --private-key $PRIVATE_KEY --value 0.2ether $ADDRESS
```

Ensure operator has enough tokens to deposit into strategy (WETH in case of Holesky deployment)
```
cast send \
    --private-key $PRIVATE-KEY \
    --value 0.1ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'
```

Register operator
```
go run ./cli  setup-operator \
    --bls-password $BLS_PASSWORD \
    --ecdsa-password $ECDSA_PASSWORD \
    --config $CONFIG_PATH \
    --strategy-address $STRATEGY_ADDRESS_TO_DEPOSIT \
    --strategy-deposit-amount $STRATEGY_DEPOSIT_AMOUNT
```

Run operator machine
```
docker run \
    -v $OPERATOR_MACHINE_DATA_PATH:/data \ 
    -v ./machine/data-preload:/data/preload \
    -p 5001:5001 \
    -p 3033:3033 \
    -e COMPUTE_ONLY='1' \
    ghcr.io/zippiehq/cartesi-lambada:latest
```

Update configuration with correct path to AVS deployment output file, network addresses of aggregator, ETH rpc, IPFS, etc.

Run operator node by providing keystore passwords and path to configuration file
```
go run operator/cmd/main.go --bls-password=$BLS_PASSWORD --ecdsa-password=$ECDSA_PASSWORD --config $CONFIG_PATH
```
