#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-holesky.yaml \
    --bls-password 2Vef5SPHZIUoVDyVJztv \
    --ecdsa-password ddKID9eipR6LSoWSjBsM \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=2Vef5SPHZIUoVDyVJztv OPERATOR_ECDSA_KEY_PASSWORD=ddKID9eipR6LSoWSjBsM operator --config tests/nodes/operators/2/config-holesky.yaml