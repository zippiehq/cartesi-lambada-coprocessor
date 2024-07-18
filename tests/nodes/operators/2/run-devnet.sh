#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-devnet.yaml \
    --bls-password b04zX6VqYzFJJI6aEKIT \
    --ecdsa-password vyuNyepG5KzsSnZqBxoK \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=b04zX6VqYzFJJI6aEKIT OPERATOR_ECDSA_KEY_PASSWORD=vyuNyepG5KzsSnZqBxoK operator --config tests/nodes/operators/2/config-devnet.yaml