#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-mainnet.yaml \
    --bls-password a8Q5Q1G1gVta5tSMbTBZ \
    --ecdsa-password Z2snWeHXk6PYTmZBcoas \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=a8Q5Q1G1gVta5tSMbTBZ OPERATOR_ECDSA_KEY_PASSWORD=Z2snWeHXk6PYTmZBcoas operator --config tests/nodes/operators/1/config-mainnet.yaml