#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/1/config-holesky.yaml \
    --bls-password mbAhhbncUefnXpYV1X5M \
    --ecdsa-password 2w1RU1MXk2LPdc3DJAZF \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=mbAhhbncUefnXpYV1X5M OPERATOR_ECDSA_KEY_PASSWORD=2w1RU1MXk2LPdc3DJAZF operator --config tests/nodes/operators/1/config-holesky.yaml