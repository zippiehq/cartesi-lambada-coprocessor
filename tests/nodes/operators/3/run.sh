#!/bin/bash

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a \
    --value 20ether 0xb96bACBd87968B1E89c905fd1c66ab2101F9E209


cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0x36a9f8849f5bec5178459701f89f47a1b6c7452413044708f33937d13fbf307c \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'


cli setup-operator \
    --config tests/nodes/operators/3/config.yaml \
    --bls-password p2OMZQ8RsCHJz8kfjEtc \
    --ecdsa-password azp2xLZsgBGPGvo7qNNT \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=p2OMZQ8RsCHJz8kfjEtc OPERATOR_ECDSA_KEY_PASSWORD=azp2xLZsgBGPGvo7qNNT operator --config tests/nodes/operators/3/config.yaml