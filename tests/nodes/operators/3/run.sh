#!/bin/bash

sleep 2

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
    --value 20ether 0x787424325C68C34637d82A853cBcf79e7bea66CC



cli setup-operator \
    --config tests/nodes/operators/3/config.yaml \
    --deployment-parameters ./contracts/script/input/deployment_parameters_devnet.json \
    --bls-password eUFlGEEtuvK19SitVtLk \
    --ecdsa-password j2HoUUhhUuDiqPUbJIqn \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=eUFlGEEtuvK19SitVtLk OPERATOR_ECDSA_KEY_PASSWORD=j2HoUUhhUuDiqPUbJIqn operator --config tests/nodes/operators/3/config.yaml