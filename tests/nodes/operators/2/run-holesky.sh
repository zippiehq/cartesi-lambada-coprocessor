#!/bin/bash

cli setup-operator \
    --config tests/nodes/operators/2/config-holesky.yaml \
    --bls-password kySHxby2AAGcMLPLpKRw \
    --ecdsa-password U2kjFhz6DEeWYSXBBltt \
    --strategy-address 0x80528D6e9A2BAbFc766965E0E26d5aB08D9CFaF9 \
    --strategy-deposit-amount 10

OPERATOR_BLS_KEY_PASSWORD=kySHxby2AAGcMLPLpKRw OPERATOR_ECDSA_KEY_PASSWORD=U2kjFhz6DEeWYSXBBltt operator --config tests/nodes/operators/2/config-holesky.yaml