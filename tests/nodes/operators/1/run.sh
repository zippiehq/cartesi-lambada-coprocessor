#!/bin/bash

sleep 0

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xD0d5E337d14D7ddeD60B381B6d4485F6454Df6Ea



cli setup-operator \
    --config tests/nodes/operators/1/config.yaml \
    --deployment-parameters ./contracts/script/input/deployment_parameters_devnet.json \
    --bls-password tVVint0uNwpu3AijyLwq \
    --ecdsa-password x5mQ0KHWJMtMO2pK4hm8 \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=tVVint0uNwpu3AijyLwq OPERATOR_ECDSA_KEY_PASSWORD=x5mQ0KHWJMtMO2pK4hm8 operator --config tests/nodes/operators/1/config.yaml