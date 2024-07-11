#!/bin/bash

cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d \
    --value 20ether 0x1161E84dDF39Dd4705833b62fe0Ac196fE9072B6


cast send \
    --rpc-url http://anvil:8545 \
    --private-key 0x43b289191f65b39584cf0042b2844e0a76f630d126694dffff8b59b99404cda9 \
    --value 10ether \
    0x94373a4919B3240D86eA41593D5eBa789FEF3848 'deposit()'


cli setup-operator \
    --config tests/nodes/operators/2/config.yaml \
    --bls-password DBCF8FbW2BODIs0LGRfX \
    --ecdsa-password 9N6rfWLwpXk3UNOOUVg5 \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=DBCF8FbW2BODIs0LGRfX OPERATOR_ECDSA_KEY_PASSWORD=9N6rfWLwpXk3UNOOUVg5 operator --config tests/nodes/operators/2/config.yaml