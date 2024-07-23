#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-holesky.yaml \
    --bls-password I0k5Um80exGzslUaSE0v \
    --ecdsa-password ui9Sa12QBtcteTNKokDV \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=I0k5Um80exGzslUaSE0v OPERATOR_ECDSA_KEY_PASSWORD=ui9Sa12QBtcteTNKokDV operator --config tests/nodes/operators/2/config-holesky.yaml