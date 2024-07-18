#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-devnet.yaml \
    --bls-password mbAhhbncUefnXpYV1X5M \
    --ecdsa-password 2w1RU1MXk2LPdc3DJAZF \
    --strategy-address 0x09635F643e140090A9A8Dcd712eD6285858ceBef \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=mbAhhbncUefnXpYV1X5M OPERATOR_ECDSA_KEY_PASSWORD=2w1RU1MXk2LPdc3DJAZF operator --config tests/nodes/operators/1/config-devnet.yaml