#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-devnet.yaml \
    --bls-password a8Q5Q1G1gVta5tSMbTBZ \
    --ecdsa-password Z2snWeHXk6PYTmZBcoas \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=a8Q5Q1G1gVta5tSMbTBZ OPERATOR_ECDSA_KEY_PASSWORD=Z2snWeHXk6PYTmZBcoas operator --config tests/nodes/operators/1/config-devnet.yaml