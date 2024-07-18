#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-mainnet.yaml \
    --bls-password mbAhhbncUefnXpYV1X5M \
    --ecdsa-password 2w1RU1MXk2LPdc3DJAZF \
    --strategy-address 0x1BeE69b7dFFfA4E2d53C2a2Df135C388AD25dCD2 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=mbAhhbncUefnXpYV1X5M OPERATOR_ECDSA_KEY_PASSWORD=2w1RU1MXk2LPdc3DJAZF operator --config tests/nodes/operators/1/config-mainnet.yaml