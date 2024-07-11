#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/3/config-devnet.yaml \
    --bls-password kHQKXbk4k5y2wYlONkCm \
    --ecdsa-password 6xByJMB6elTDeSPT8ppp \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=kHQKXbk4k5y2wYlONkCm OPERATOR_ECDSA_KEY_PASSWORD=6xByJMB6elTDeSPT8ppp operator --config tests/nodes/operators/3/config-devnet.yaml