#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-holesky.yaml \
    --bls-password b04zX6VqYzFJJI6aEKIT \
    --ecdsa-password vyuNyepG5KzsSnZqBxoK \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=b04zX6VqYzFJJI6aEKIT OPERATOR_ECDSA_KEY_PASSWORD=vyuNyepG5KzsSnZqBxoK operator --config tests/nodes/operators/2/config-holesky.yaml