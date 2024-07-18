#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-mainnet.yaml \
    --bls-password b04zX6VqYzFJJI6aEKIT \
    --ecdsa-password vyuNyepG5KzsSnZqBxoK \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=b04zX6VqYzFJJI6aEKIT OPERATOR_ECDSA_KEY_PASSWORD=vyuNyepG5KzsSnZqBxoK operator --config tests/nodes/operators/2/config-mainnet.yaml