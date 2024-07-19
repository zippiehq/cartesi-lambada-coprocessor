#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-holesky.yaml \
    --bls-password XuCFvAtotd6f9BokY5C4 \
    --ecdsa-password tnMtegPwqKSiRZkrXB2Y \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=XuCFvAtotd6f9BokY5C4 OPERATOR_ECDSA_KEY_PASSWORD=tnMtegPwqKSiRZkrXB2Y operator --config tests/nodes/operators/1/config-holesky.yaml