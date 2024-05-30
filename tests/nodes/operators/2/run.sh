#!/bin/bash

sleep 1

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0xAfE04e56545a90b15DcbDB682F55C151D7af2Eca


cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xb24997c3fa2ed5129e803e36e8f1224010f8aefb74e64b3de7017d889c5fd28e \
    0x7a2088a1bFc9d81c55368AE168C2C02570cB814F "mint(address,uint256)" 0xAfE04e56545a90b15DcbDB682F55C151D7af2Eca 10    


cli setup-operator \
    --config tests/nodes/operators/2/config.yaml \
    --bls-password MzrzurCOp4yMVn03wd4b \
    --ecdsa-password 9htFzi2Qm3710fPh65zB \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=MzrzurCOp4yMVn03wd4b OPERATOR_ECDSA_KEY_PASSWORD=9htFzi2Qm3710fPh65zB operator --config tests/nodes/operators/2/config.yaml