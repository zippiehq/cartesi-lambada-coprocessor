#!/bin/bash

sleep 1

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x3Ec76Db4e8ad4993E8d343862911ef598d679556



cli setup-operator \
    --config tests/nodes/operators/2/config.yaml \
    --deployment-parameters ./contracts/script/input/deployment_parameters_devnet.json \
    --bls-password JBSLtY9AGCWo4vsib2g0 \
    --ecdsa-password i3Qu2vB3NY2fepKidlqO \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=JBSLtY9AGCWo4vsib2g0 OPERATOR_ECDSA_KEY_PASSWORD=i3Qu2vB3NY2fepKidlqO operator --config tests/nodes/operators/2/config.yaml