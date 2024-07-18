#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-holesky.yaml \
    --bls-password a8Q5Q1G1gVta5tSMbTBZ \
    --ecdsa-password Z2snWeHXk6PYTmZBcoas \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=a8Q5Q1G1gVta5tSMbTBZ OPERATOR_ECDSA_KEY_PASSWORD=Z2snWeHXk6PYTmZBcoas operator --config tests/nodes/operators/1/config-holesky.yaml