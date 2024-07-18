#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-holesky.yaml \
    --bls-password buM47eslEETBTZVhp5DW \
    --ecdsa-password tGDaHRFp0LKNGcobHUYE \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=buM47eslEETBTZVhp5DW OPERATOR_ECDSA_KEY_PASSWORD=tGDaHRFp0LKNGcobHUYE operator --config tests/nodes/operators/2/config-holesky.yaml